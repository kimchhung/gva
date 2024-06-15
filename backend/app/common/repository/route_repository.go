package repository

import (
	"github.com/kimchhung/gva/backend/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend/internal/ent"
	"github.com/kimchhung/gva/backend/utils/pagi"
)

type RouteRepository struct {
	db *database.Database
}

func NewRouteRepository(database *database.Database) *RouteRepository {
	return &RouteRepository{
		db: database,
	}
}

func (r *RouteRepository) C() *ent.RouteClient {
	return r.db.Route
}

// For query
func (r *RouteRepository) Q(opts ...pagi.InterceptorOption) *ent.RouteQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}
