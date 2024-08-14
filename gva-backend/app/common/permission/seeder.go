package permission

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
)

type PermissionSeeder struct {
	group  PermissionGroup
	scopes []permissionScope
}

func NewSeeder(group PermissionGroup, scopes ...permissionScope) database.Seeder {
	return &PermissionSeeder{
		group:  group,
		scopes: scopes,
	}
}

func (seeder PermissionSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	return conn.Permission.Query().Where(permission.GroupEQ(string(seeder.group))).Count(ctx)
}

func (seeder PermissionSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	return conn.Permission.CreateBulk(createBulkPermissionDto(conn, seeder.scopes...)...).Exec(ctx)
}
