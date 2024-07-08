package permission

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
)

const (
	AdminGroup PermissionGroup = "ADMIN"
)

var (
	AdminSuper  = newKey(AdminGroup, ActionSuper)
	AdminView   = newKey(AdminGroup, ActionView)
	AdminAdd    = newKey(AdminGroup, ActionAdd)
	AdminEdit   = newKey(AdminGroup, ActionEdit)
	AdminDelete = newKey(AdminGroup, ActionDelete)
)

type AdminSeeder struct {
	group PermissionGroup
	keys  []PermissionKey
}

func NewAdminSeeder() database.Seeder {
	return &AdminSeeder{
		group: AdminGroup,
		keys: []PermissionKey{
			AdminSuper,
			AdminView,
			AdminAdd,
			AdminDelete,
		},
	}
}

func (seeder AdminSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	return conn.Permission.Query().Where(permission.GroupEQ(string(seeder.group))).Count(ctx)
}

func (seeder AdminSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	return conn.Permission.CreateBulk(createBulkPermissionDto(conn, seeder.keys...)...).Exec(ctx)
}
