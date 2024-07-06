package seeds

import (
	"context"
	"fmt"
	"log"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/json"
)

var _ interface{ database.Seeder } = (*RouterSeeder)(nil)

type RouterSeeder struct {
}

func (RouterSeeder) Count(ctx context.Context, db *ent.Client) (int, error) {
	return db.Route.Query().Count(ctx)
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

	return database.WithTx(ctx, conn, func(tx *ent.Tx) error {
		creatRoutes := make([]*ent.RouteCreate, len(routers))
		for i, r := range routers {

			createChildren := make([]*ent.RouteCreate, len(r.Edges.Children))
			for _i, c := range r.Edges.Children {
				createChildren[_i] = tx.Route.Create().
					SetIsEnable(true).
					SetPath(c.Path).
					SetComponent(c.Component).
					SetNillableRedirect(c.Redirect).
					SetName(c.Name).
					SetType(c.Type).
					SetMeta(c.Meta).
					SetOrder(_i)
			}

			createdChildren, err := tx.Route.CreateBulk(createChildren...).Save(ctx)
			if err != nil {
				return fmt.Errorf("createdChildren: %v", err)
			}

			creatRoutes[i] = tx.Route.Create().
				SetComponent(r.Component).
				SetPath(r.Path).
				SetIsEnable(true).
				SetMeta(r.Meta).
				SetName(r.Name).
				SetNillableRedirect(r.Redirect).
				SetType(r.Type).
				SetOrder(i).
				AddChildren(createdChildren...)
		}

		_, err := tx.Route.CreateBulk(creatRoutes...).Save(ctx)
		return err
	})

}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
