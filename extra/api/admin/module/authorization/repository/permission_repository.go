package repository

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rql"
	"github.com/kimchhung/gva/extra/utils/pagi"
)

type PermissionRepository struct {
	db *database.Database
}

func NewPermissionRepository(database *database.Database) *PermissionRepository {
	return &PermissionRepository{
		database,
	}
}

func (r *PermissionRepository) C() *ent.PermissionClient {
	return r.db.Permission
}

func (r *PermissionRepository) DB() *database.Database {
	return r.db
}

func (r *PermissionRepository) RQL(p *rql.Params, opts ...func(*sql.Selector)) *ent.PermissionQuery {
	return pagi.RQL(r.C().Query(), p, opts...)
}
