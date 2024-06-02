package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
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

// For mutation
func (r *AdminRepository) C() *ent.AdminClient {
	return r.db.Admin
}

// For query
func (r *AdminRepository) Q(opts ...pagi.InterceptorOption) *ent.AdminQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return r.C().Query()
}
