package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/pagi"
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

	return pagi.WithInterceptor(r.C().Query(), opts...)
}
