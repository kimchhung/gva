package permission

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
)

const (
	RouteGroup PermissionGroup = "ROUTE"
)

var (
	RouteSuper  = newKey(RouteGroup, ActionSuper)
	RouteView   = newKey(RouteGroup, ActionView)
	RouteAdd    = newKey(RouteGroup, ActionAdd)
	RouteEdit   = newKey(RouteGroup, ActionEdit)
	RouteDelete = newKey(RouteGroup, ActionDelete)
)

type RouteSeeder struct {
	group PermissionGroup
	keys  []PermissionKey
}

func NewRouteSeeder() database.Seeder {
	return &RouteSeeder{
		group: RouteGroup,
		keys: []PermissionKey{
			RouteSuper,
			RouteView,
			RouteAdd,
			RouteDelete,
		},
	}
}

func (seeder RouteSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	return conn.Permission.Query().Where(permission.GroupEQ(string(seeder.group))).Count(ctx)
}

func (seeder RouteSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	return conn.Permission.CreateBulk(createBulkPermissionDto(conn, seeder.keys...)...).Exec(ctx)
}
