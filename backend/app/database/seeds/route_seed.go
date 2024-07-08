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
	return db.Menu.Query().Count(ctx)
}

func getRouteData() (routes []*ent.Menu) {
	bytes, err := json.ReadJsonFile("./app/database/data/menu_data.json")
	if err != nil {
		log.Panicf("can't raed seed data %v", err)
	}

	if err := bytes.Out(&routes); err != nil {
		log.Panicf("can't parse seed data %v", err)
	}

	return routes
}

func (s RouterSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	routers := getRouteData()

	return database.WithTx(ctx, conn, func(tx *ent.Tx) error {

		_, err := s.seedRouteRecursively(ctx, tx, routers...)
		if err != nil {
			return err
		}

		return nil
	})
}

// seedRouteRecursively seeds a single route and its children recursively
func (s RouterSeeder) seedRouteRecursively(ctx context.Context, tx *ent.Tx, routes ...*ent.Menu) (createdRoutes []*ent.Menu, err error) {
	for i, r := range routes {
		createRoute := tx.Menu.Create().
			SetComponent(r.Component).
			SetPath(r.Path).
			SetIsEnable(true).
			SetMeta(r.Meta).
			SetName(r.Name).
			SetOrder(i).
			SetNillableRedirect(r.Redirect).
			SetType(r.Type).
			SetOrder(len(r.Edges.Children)) // Assuming order is determined by the number of children

		var createdChildren []*ent.Menu
		if len(r.Edges.Children) > 0 {
			createdChildren, err = s.seedRouteRecursively(ctx, tx, r.Edges.Children...)
			if err != nil {
				return nil, fmt.Errorf("failed to seed child route: %v", err)
			}
		}

		createdRoute := createRoute.AddChildren(createdChildren...).SaveX(ctx)
		createdRoutes = append(createdRoutes, createdRoute)
	}

	return createdRoutes, nil
}
