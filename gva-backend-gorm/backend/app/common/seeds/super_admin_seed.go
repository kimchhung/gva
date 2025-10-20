package seeds

import (
	"backend/app/common/model"
	"backend/app/common/service"
	"backend/env"
	"backend/internal/bootstrap/database"
	"backend/internal/ctxutil"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type SuperAdminSeeder struct {
}

func NewSuperAdminSeeder() database.Seeder {
	return &SuperAdminSeeder{}
}

func (s SuperAdminSeeder) Name() string {
	return "SuperAdminSeeder"
}

func (s SuperAdminSeeder) Count(ctx context.Context, db *gorm.DB) (int, error) {
	cfg := ctxutil.MustValue[*env.Config](ctx)

	var total int64
	if err := db.Model(model.Admin{}).Where("username = ?", cfg.DB.Seed.SuperAdmin.Username).Count(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}

func (s SuperAdminSeeder) Seed(ctx context.Context, db *gorm.DB) error {
	cfg := ctxutil.MustValue[*env.Config](ctx)
	password_ := ctxutil.MustValue[*service.PasswordService](ctx)

	db.Transaction(func(tx *gorm.DB) error {
		adminRole := &model.AdminRole{
			BaseModel:   model.NewBaseModel(),
			Name:        "Super Admin",
			NameID:      "SUPER_ADMIN",
			Description: "Super Admin can control everything",
			Status:      1,
			Type:        model.AdminRoleTypeNotChangeable,
			Order:       0,
		}
		adminRole.ID = 1
		pw, err := password_.HashPassword(cfg.DB.Seed.SuperAdmin.Password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}

		admin := &model.Admin{
			BaseModel:    model.NewBaseModel(),
			Name:         "Super Admin",
			Username:     cfg.DB.Seed.SuperAdmin.Username,
			Roles:        []*model.AdminRole{adminRole},
			PasswordHash: pw,
			Status:       1,
		}

		if err := tx.Create(admin).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}
