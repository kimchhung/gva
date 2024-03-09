package repository

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rql"
)

type AdminRepository struct {
	db *database.Database
}

func NewAdminRepository(database *database.Database) *AdminRepository {
	return &AdminRepository{
		database,
	}
}

func (r *AdminRepository) C() *ent.AdminClient {
	return r.db.Admin
}

func (r *AdminRepository) RQL(p *rql.Params) *ent.AdminQuery {
	return r.C().Query().
		Where(func(s *sql.Selector) { s.Where(sql.ExprP(p.FilterExp.String(), p.FilterArgs...)) }).
		Order(func(s *sql.Selector) { s.OrderBy(p.Sort...) }).
		Limit(p.Limit).
		Offset(p.Offset)
}

func (r *AdminRepository) DB() *database.Database {
	return r.db
}
