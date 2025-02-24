package permission

import (
	"context"
	"backend/app/common/model"
	"backend/internal/bootstrap/database"

	"gorm.io/gorm"
)

type PermissionSeeder struct {
	group  TPermissionGroup
	scopes []permissionScope
}

func NewSeeder(group TPermissionGroup, scopes ...permissionScope) database.Seeder {
	return &PermissionSeeder{
		group:  group,
		scopes: scopes,
	}
}

func (s PermissionSeeder) Name() string {
	return string(s.group)
}

func (seeder PermissionSeeder) Count(ctx context.Context, db *gorm.DB) (int, error) {
	var total int64
	if err := db.Model(model.Permission{}).
		Where(model.Permission{Group: string(seeder.group)}).
		Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (seeder PermissionSeeder) Seed(ctx context.Context, db *gorm.DB) error {
	return db.Model(model.Permission{}).Create(createBulkPermissionDto(seeder.scopes...)).Error
}
