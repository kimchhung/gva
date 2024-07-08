package permission

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
)

const (
	AdminRoleGroup PermissionGroup = "ADMIN_ROLE"
)

var (
	AdminRoleSuper  = newKey(AdminGroup, ActionSuper)
	AdminRoleView   = newKey(AdminGroup, ActionSuper)
	AdminRoleModify = newKey(AdminGroup, ActionSuper)
	AdminRoleDelete = newKey(AdminGroup, ActionSuper)
)

type AdminRoleSeeder struct {
	group PermissionGroup
	keys  []PermissionKey
}

func NewAdminRoleSeeder() database.Seeder {
	return &AdminRoleSeeder{
		group: AdminRoleGroup,
		keys: []PermissionKey{
			AdminRoleSuper,
			AdminRoleView,
			AdminRoleModify,
			AdminRoleDelete,
		},
	}
}

func (seeder AdminRoleSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	group := string(seeder.group)
	return conn.Permission.Query().Where(permission.GroupEQ(group)).Count(ctx)
}

func (seeder AdminRoleSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	perms := createBulkPermissionDto(conn, seeder.keys...)
	return conn.Permission.CreateBulk(perms...).Exec(ctx)
}
