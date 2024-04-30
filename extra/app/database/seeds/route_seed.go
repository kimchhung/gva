package seeds

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/route"
	"github.com/kimchhung/gva/extra/utils/json"
	"github.com/kimchhung/gva/extra/utils/routeutil"
)

var _ interface {
	database.Seeder
} = (*RouterSeeder)(nil)

type RouterSeeder struct {
}

func (RouterSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	return conn.Route.Query().Count(context.TODO())
}

func getRouteData() (routes []*ent.Route) {
	bytes, err := json.ReadJsonFile("./app/database/data/routes_data.json")
	if err != nil {
		log.Panicf("can't raed seed data %v", err)
	}

	if err := bytes.Out(&routes); err != nil {
		log.Panicf("can't parse seed data %v", err)
	}

	return routes
}

func (RouterSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	routers := getRouteData()
	flats := routeutil.FlattenNestedRoutes(routers)

	childToParent := map[int]int{}

	database.WithTx(ctx, conn, func(tx *ent.Tx) error {

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

	return nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
