package seeds

import (
	"context"
	"fmt"

	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
)

var _ interface {
	database.Seeder
} = (*RouterSeeder)(nil)

type SuperAdminSeeder struct {
	cfg       *config.Config
	password_ *services.PasswordService
}

func (s SuperAdminSeeder) Count(conn *ent.Client) (int, error) {
	return conn.Admin.Query().
		Where(admin.Username(s.cfg.Seed.SuperAdmin.Username)).
		Count(context.Background())
}

func (s SuperAdminSeeder) Seed(conn *ent.Client) error {
	ctx := context.Background()
	tx, _ := conn.Tx(ctx)

	role := tx.Role.Create().
		SetDescription("Super Admin can control everything").
		SetOrder(0).
		SetName("SUPER_ADMIN").
		SetIsEnable(true).
		SetIsChangeable(false).
		SaveX(ctx)

	pw, _ := s.password_.HashPassword(s.cfg.Seed.SuperAdmin.Password)
	tx.Admin.Create().
		SetUsername(s.cfg.Seed.SuperAdmin.Username).
		SetPassword(pw).SetWhitelistIps([]string{"0.0.0.0"}).
		SetDisplayName("super admin").AddRoles(role).
		SaveX(ctx)

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("SuperAdminSeeder err %v", err)
	}

	return nil
}

func NewSuperAdminSeeder(cfg *config.Config) SuperAdminSeeder {
	return SuperAdminSeeder{
		cfg:       cfg,
		password_: services.NewPasswordService(cfg),
	}
}
