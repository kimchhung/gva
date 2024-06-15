package permission

import (
	"context"

	"github.com/kimchhung/gva/backend/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend/internal/ent"
	"github.com/kimchhung/gva/backend/internal/ent/permission"
)

const (
	AdminGroup = "ADMIN"

	AdminSuper  PermissionKey = "ADMIN.SUPER"
	AdminView   PermissionKey = "ADMIN.VIEW"
	AdminModify PermissionKey = "ADMIN.MODIFY"
	AdminDelete PermissionKey = "ADMIN.DELETE"
)

func init() {
	groups = append(groups, AdminGroup)
	keys = append(keys, AdminSuper, AdminView, AdminModify, AdminDelete)
}

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
