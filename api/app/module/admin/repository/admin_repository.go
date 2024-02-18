package repository

import (
	"github.com/kimchhung/gva/internal/bootstrap/database"
	"github.com/kimchhung/gva/internal/ent"
)

type AdminRepository struct {
	DB *database.Database
}

func NewAdminRepository(database *database.Database) *AdminRepository {
	return &AdminRepository{
		database,
	}
}

func (r *AdminRepository) Client() *ent.AdminClient {
	return r.DB.Ent.Admin
}

