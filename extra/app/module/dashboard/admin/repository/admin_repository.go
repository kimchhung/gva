package repository

import (
	"context"

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

func (r *AdminRepository) Q(ctx context.Context, p *rql.Params) ([]*ent.Admin, error) {
	return r.C().Query().Where(
		func(s *sql.Selector) {
			s.Where(sql.ExprP(p.FilterExp.String(), p.FilterArgs...))
		}).
		Limit(p.Limit).
		Offset(p.Offset).
		All(ctx)
}

func (r *AdminRepository) DB() *database.Database {
	return r.db
}
