package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
)

type RoleRepository struct {
	db *database.Database
}

func NewRoleRepository(database *database.Database) *RoleRepository {
	return &RoleRepository{
		database,
	}
}

func (r *RoleRepository) C() *ent.RoleClient {
	return r.db.Role
}

func (r *RoleRepository) DB() *database.Database {
	return r.db
}
