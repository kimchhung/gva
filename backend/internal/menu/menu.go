package menu

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/route"
	"github.com/gva/utils/json"
	"github.com/gva/utils/routeutil"
)

func LoadRouteFromFile(filePath string) (routes []*ent.Route) {
	bytes, err := json.ReadJsonFile(filePath)
	if err != nil {
		log.Panicf("can't raed seed data %v", err)
	}

	if err := bytes.Out(&routes); err != nil {
		log.Panicf("can't parse seed data %v", err)
	}

	return routes
}

func PullRoutes(ctx context.Context, conn *ent.Client, filePath string) {
	flats, err := conn.Route.Query().All(ctx)
	if err != nil {
		log.Panicf("failed querying routes: %v", err)
	}

	nested := routeutil.GroupRouteToNested(flats)

	if err := json.WriteJsonToFile(json.MustJSON(nested), filePath); err != nil {
		log.Panicf("failed write routes: %v", err)
	}
}

func PushRouters(ctx context.Context, conn *ent.Client, filePath string) {
	routers := LoadRouteFromFile(filePath)
	flats := routeutil.FlattenNestedRoutes(routers)

	if len(routers) == 0 {
		return
	}

	childToParent := map[string]string{}
	database.WithTx(ctx, conn, func(tx *ent.Tx) error {
		tx.Route.Delete().ExecX(softdelete.SkipSoftDelete(ctx))

		for _, r := range flats {
			if r.ParentID != nil {
				childToParent[r.ID] = *r.ParentID
			}

			if strings.Contains(r.Component, "#") {
				r.Type = route.TypeCataLog
			} else {
				r.Type = route.TypeMenu
			}

			_, err := tx.Route.Create().SetID(r.ID).
				SetComponent(r.Component).
				SetPath(r.Path).
				SetIsEnable(true).
				SetMeta(r.Meta).
				SetName(r.Name).
				SetNillableRedirect(r.Redirect).
				SetType(r.Type).Save(context.Background())

			if err != nil {
				return fmt.Errorf("create  routers: %w", err)
			}

		}

		for cid, pid := range childToParent {
			_, err := tx.Route.UpdateOneID(cid).SetParentID(pid).Save(context.Background())
			if err != nil {
				return fmt.Errorf("save routers: %w", err)
			}
		}

		return nil

	})

}
