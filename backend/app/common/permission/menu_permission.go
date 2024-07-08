package permission

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
)

const (
	MenuGroup PermissionGroup = "Menu"
)

var (
	MenuSuper  = newKey(MenuGroup, ActionSuper)
	MenuView   = newKey(MenuGroup, ActionView)
	MenuAdd    = newKey(MenuGroup, ActionAdd)
	MenuEdit   = newKey(MenuGroup, ActionEdit)
	MenuDelete = newKey(MenuGroup, ActionDelete)
)

type MenuSeeder struct {
	group PermissionGroup
	keys  []PermissionKey
}

func init() {
	allSeeders = append(allSeeders, NewMenuSeeder())
}

func NewMenuSeeder() database.Seeder {
	return &MenuSeeder{
		group: MenuGroup,
		keys: []PermissionKey{
			MenuSuper,
			MenuView,
			MenuAdd,
			MenuDelete,
		},
	}
}

func (seeder MenuSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	return conn.Permission.Query().Where(permission.GroupEQ(string(seeder.group))).Count(ctx)
}

func (seeder MenuSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	return conn.Permission.CreateBulk(createBulkPermissionDto(conn, seeder.keys...)...).Exec(ctx)
}
