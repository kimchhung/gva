package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/pagi"
)

type MenuRepository struct {
	db *database.Database
}

func NewMenuRepository(database *database.Database) *MenuRepository {
	return &MenuRepository{
		db: database,
	}
}

func (r *MenuRepository) C() *ent.MenuClient {
	return r.db.Menu
}

// For query
func (r *MenuRepository) Q(opts ...pagi.InterceptorOption) *ent.MenuQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}
