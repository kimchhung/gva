package permissions

import (
	"context"

	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/permission"
)

const (
	AdminGroup = "ADMIN"

	AdminSuper  PermissionKey = "ADMIN.SUPER"
	AdminView   PermissionKey = "ADMIN.VIEW"
	AdminModify PermissionKey = "ADMIN.MODIFY"
	AdminDelete PermissionKey = "ADMIN.DELETE"
)

var _ database.Seeder = (*AdminPermissionSeeder)(nil)

type AdminPermissionSeeder struct {
}

func (AdminPermissionSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	group := AdminGroup
	return conn.Permission.Query().Where(permission.GroupEQ(group)).Count(ctx)
}

func (AdminPermissionSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	group := AdminGroup

	perms := createBulkPermissionDto(conn, group,
		AdminSuper,
		AdminView,
		AdminModify,
		AdminDelete,
	)

	return conn.Permission.CreateBulk(perms...).Exec(ctx)
}
