package permissions

import (
	"context"

	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/permission"
)

const (
	AdminRoleGroup = "ADMIN_ROLE"

	AdminRoleSuper  PermissionKey = "ADMIN_ROLE.SUPER"
	AdminRoleView   PermissionKey = "ADMIN_ROLE.VIEW"
	AdminRoleModify PermissionKey = "ADMIN_ROLE.MODIFY"
	AdminRoleDelete PermissionKey = "ADMIN_ROLE.DELETE"
)

var _ database.Seeder = (*AdminRolePermissionSeeder)(nil)

type AdminRolePermissionSeeder struct {
}

func (AdminRolePermissionSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	group := AdminRoleGroup

	return conn.Permission.Query().Where(permission.GroupEQ(group)).Count(ctx)
}

func (AdminRolePermissionSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	perms := createBulkPermissionDto(conn, AdminRoleGroup,
		AdminRoleSuper,
		AdminRoleView,
		AdminRoleModify,
		AdminRoleDelete,
	)

	return conn.Permission.CreateBulk(perms...).Exec(ctx)
}
