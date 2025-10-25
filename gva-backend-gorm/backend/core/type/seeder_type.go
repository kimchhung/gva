package coretype

import (
	"context"

	"gorm.io/gorm"
)

type Seeder interface {
	Name() string
	Count(ctx context.Context, conn *gorm.DB) (int, error)
	Seed(ctx context.Context, conn *gorm.DB) error
}
