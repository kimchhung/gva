package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/pagi"
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

// For query
func (r *PermissionRepository) Q(opts ...pagi.InterceptorOption) *ent.PermissionQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}

