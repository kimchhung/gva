package seeds

import (
	"context"
	"fmt"

	"github.com/kimchhung/gva/extra/app/common/mock"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/utils/routeutil"
)

var _ interface {
	database.Seeder
} = (*RouterSeeder)(nil)

type RouterSeeder struct {
}

func (RouterSeeder) Count(conn *ent.Client) (int, error) {
	return conn.Route.Query().Count(context.TODO())
}

func (RouterSeeder) Seed(conn *ent.Client) error {
	routers := mock.GetRoutes()
	flats := routeutil.FlattenNestedRoutes(routers)

	childToParent := map[int]int{}

	ctx := context.Background()
	tx, _ := conn.Tx(ctx)

	for _, r := range flats {
		saved, err := tx.Route.Create().SetID(r.ID).
			SetTitle(r.Title).
			SetComponent(r.Component).
			SetPath(r.Path).
			SetIsEnable(true).
			SetMeta(r.Meta).
			SetName(r.Name).
			SetRedirect(r.Redirect).
			SetType(r.Type).Save(context.Background())

		if err != nil {
			return fmt.Errorf("create  routers: %w", err)
		}

		if r.ParentID != nil {
			childToParent[saved.ID] = *r.ParentID
		}
	}

	for cid, pid := range childToParent {
		_, err := tx.Route.UpdateOneID(cid).SetParentID(pid).Save(context.Background())
		if err != nil {
			return fmt.Errorf("save routers: %w", err)
		}
	}

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("commit routers: %w", err)
	}

	return nil
}
