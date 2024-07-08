package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/pagi"
)

type RouteRepository struct {
	db *database.Database
}

func NewRouteRepository(database *database.Database) *RouteRepository {
	return &RouteRepository{
		db: database,
	}
}

func (r *RouteRepository) C() *ent.MenuClient {
	return r.db.Menu
}

// For query
func (r *RouteRepository) Q(opts ...pagi.InterceptorOption) *ent.MenuQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}
