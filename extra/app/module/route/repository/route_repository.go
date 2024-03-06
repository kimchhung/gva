package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
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
