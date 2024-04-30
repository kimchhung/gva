package repository

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rql"
	"github.com/kimchhung/gva/extra/utils/pagi"
)

type RouteRepository struct {
	db *database.Database
}

func NewRouteRepository(database *database.Database) *RouteRepository {
	return &RouteRepository{
		database,
	}
}

func (r *RouteRepository) C() *ent.RouteClient {
	return r.db.Route
}

func (r *RouteRepository) DB() *database.Database {
	return r.db
}

func (r *RouteRepository) RQL(p *rql.Params, opts ...func(*sql.Selector)) *ent.RouteQuery {
	return pagi.RQL(r.C().Query(), p, opts...)
}
