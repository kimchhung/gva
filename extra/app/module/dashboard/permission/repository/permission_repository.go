package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
)

type PermissionRepository struct {
	DB *database.Database
}

func NewPermissionRepository(database *database.Database) *PermissionRepository {
	return &PermissionRepository{
		database,
	}
}

func (r *PermissionRepository) C() *ent.PermissionClient {
	return r.DB.Permission
}
