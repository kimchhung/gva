package seeds

import (
	"context"
	"fmt"

	"github.com/gva/app/common/service"
	"github.com/gva/env"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/admin"
)

type SuperAdminSeeder struct {
}

func NewSuperAdminSeeder() database.Seeder {
	return &SuperAdminSeeder{}
}

func (s SuperAdminSeeder) Count(ctx context.Context, db *ent.Client) (int, error) {
	cfg := ctx.Value(env.Config{}).(*env.Config)

	return db.Admin.Query().
		Where(admin.Username(cfg.Seed.SuperAdmin.Username)).
		Count(ctx)
}

func (s SuperAdminSeeder) Seed(ctx context.Context, conn *ent.Client) error {
	cfg := ctx.Value(env.Config{}).(*env.Config)
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
			return fmt.Errorf("hash password: %w", err)
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
