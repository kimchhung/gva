package menu

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/menu"
	"github.com/gva/internal/json"
	"github.com/gva/internal/utils/routeutil"
)

func LoadMenuFromFile(filePath string) (menus []*ent.Menu) {
	bytes, err := json.ReadJsonFile(filePath)
	if err != nil {
		log.Panicf("can't raed seed data %v", err)
	}

	if err := bytes.Out(&menus); err != nil {
		log.Panicf("can't parse seed data %v", err)
	}

	return menus
}

func PullMenuList(ctx context.Context, conn *ent.Client, filePath string) {
	flats, err := conn.Menu.Query().All(ctx)
	if err != nil {
		log.Panicf("failed querying menus: %v", err)
	}

	nested := routeutil.GroupRouteToNested(flats)

	if err := json.WriteJsonToFile(json.MustJSON(nested), filePath); err != nil {
		log.Panicf("failed write menus: %v", err)
	}
}

func PushMenuList(ctx context.Context, conn *ent.Client, filePath string) {
	routers := LoadMenuFromFile(filePath)
	flats := routeutil.FlattenNestedMenu(routers)

	if len(routers) == 0 {
		return
	}

	childToParent := map[pxid.ID]pxid.ID{}
	database.WithTx(ctx, conn, func(tx *ent.Tx) error {
		tx.Menu.Delete().ExecX(softdelete.SkipSoftDelete(ctx))
		for _, r := range flats {
			if r.Pid != nil {
				childToParent[r.ID] = *r.Pid
			}

			if strings.Contains(r.Component, "#") {
				r.Type = menu.TypeCataLog
			} else {
				r.Type = menu.TypeMenu
			}

			_, err := tx.Menu.Create().SetID(r.ID).
				SetComponent(r.Component).
				SetPath(r.Path).
				SetIsEnable(true).
				SetMeta(r.Meta).
				SetName(r.Name).
				SetNillableRedirect(r.Redirect).
				SetType(r.Type).
				Save(ctx)

			if err != nil {
				return fmt.Errorf("create  routers: %w", err)
			}
		}

		for cid, pid := range childToParent {
			_, err := tx.Menu.UpdateOneID(cid).SetPid(pid).Save(context.Background())
			if err != nil {
				return fmt.Errorf("save routers: %w", err)
			}
		}

		return nil

	})

}
