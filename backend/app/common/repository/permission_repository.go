package repository

import (
	"github.com/kimchhung/gva/backend/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend/internal/ent"
	"github.com/kimchhung/gva/backend/utils/pagi"
)

type PermissionRepository struct {
	db *database.Database
}

func NewPermissionRepository(database *database.Database) *PermissionRepository {
	return &PermissionRepository{
		database,
	}
}

// For mutation
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
