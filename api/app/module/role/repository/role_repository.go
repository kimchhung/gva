package repository

import (
	"gva/internal/bootstrap/database"
	"gva/internal/ent"
)

type RoleRepository struct {
	DB *database.Database
}

func NewRoleRepository(database *database.Database) *RoleRepository {
	return &RoleRepository{
		database,
	}
}

func (r *RoleRepository) Client() *ent.RoleClient {
	return r.DB.Ent.Role
}

