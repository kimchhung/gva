package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
)

type RouteRepository struct {
	DB *database.Database
}

func NewRouteRepository(database *database.Database) *RouteRepository {
	return &RouteRepository{
		database,
	}
}

func (r *RouteRepository) Client() *ent.RouteClient {
	return r.DB.Ent.Route
}

