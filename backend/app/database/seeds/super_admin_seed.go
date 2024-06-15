package seeds

import (
	"context"
	"fmt"

	"github.com/kimchhung/gva/backend-echo/app/common/service"
	"github.com/kimchhung/gva/backend-echo/config"
	"github.com/kimchhung/gva/backend-echo/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend-echo/internal/ent"
	"github.com/kimchhung/gva/backend-echo/internal/ent/admin"
)

var _ interface {
	database.Seeder
} = (*RouterSeeder)(nil)

type SuperAdminSeeder struct {
}

func (s SuperAdminSeeder) Count(ctx context.Context, conn *ent.Client) (int, error) {
	cfg := ctx.Value(config.Config{}).(*config.Config)

	return conn.Admin.Query().
		Where(admin.Username(cfg.Seed.SuperAdmin.Username)).
		Count(ctx)
}

func (s SuperAdminSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	cfg := ctx.Value(config.Config{}).(*config.Config)
	password_ := ctx.Value(service.PasswordService{}).(*service.PasswordService)

	database.WithTx(ctx, conn, func(tx *ent.Tx) error {
		role := tx.Role.Create().
			SetDescription("Super Admin can control everything").
			SetOrder(0).
			SetName("SUPER_ADMIN").
			SetIsEnable(true).
			SetIsChangeable(false).
			SaveX(ctx)

		pw, err := password_.HashPassword(cfg.Seed.SuperAdmin.Password)
		if err != nil {
			return rollback(tx, fmt.Errorf("hash password: %w", err))
		}

		tx.Admin.Create().
			SetUsername(cfg.Seed.SuperAdmin.Username).
			SetPassword(pw).SetWhitelistIps([]string{"0.0.0.0"}).
			SetDisplayName("super admin").AddRoles(role).
			SaveX(ctx)

		return nil
	})

	return nil
}
