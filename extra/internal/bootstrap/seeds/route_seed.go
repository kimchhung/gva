package seeds

import (
	"context"
	"fmt"
	"strings"

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
		routeType := 0
		if r.ParentID != nil {
			childToParent[r.ID] = *r.ParentID
		}

		if !strings.Contains(r.Component, "#") {
			routeType = 1
		}

		_, err := tx.Route.Create().SetID(r.ID).
			SetTitle(r.Title).
			SetComponent(r.Component).
			SetPath(r.Path).
			SetIsEnable(true).
			SetMeta(r.Meta).
			SetName(r.Name).
			SetRedirect(r.Redirect).
			SetType(routeType).Save(context.Background())

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

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("commit routers: %w", err)
	}

	return nil
}
