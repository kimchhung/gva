package permission

import (
	"backend/app/share/model"
	coretype "backend/core/type"
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	allSeeders []coretype.Seeder
)

func AllSeeders() []coretype.Seeder {
	return append([]coretype.Seeder{}, allSeeders...)
}

type PermissionSeeder struct {
	group  TPermissionGroup
	scopes []permissionScope
}

func NewSeeder(group TPermissionGroup, scopes ...permissionScope) coretype.Seeder {
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

func createBulkPermissionDto(scopes ...permissionScope) []*model.Permission {
	bulks := make([]*model.Permission, len(scopes))

	for i, scope := range scopes {
		group, _, err := scope.Value()
		if err != nil {
			panic(err)
		}

		bulks[i] = &model.Permission{
			CreatedAt: time.Now(),
			Group:     string(group),
			Scope:     string(scope),
			Order:     i,
			Name:      scope.Name(),
		}
	}

	return bulks
}
