package repository

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rql"
	"github.com/kimchhung/gva/extra/utils/pagi"
)

type AdminRepository struct {
	db *database.Database
}

func NewAdminRepository(database *database.Database) *AdminRepository {
	return &AdminRepository{
		database,
	}
}

func (r *AdminRepository) DB() *ent.Client {
	return r.db.Client
}

func (r *AdminRepository) C() *ent.AdminClient {
	return r.DB().Admin
}

func (r *AdminRepository) RQL(p *rql.Params, opts ...func(*sql.Selector)) *ent.AdminQuery {
	return pagi.RQL(r.C().Query(), p, opts...)
}
