package repos

import (
	"gva/internal/bootstrap/database"
	"gva/internal/ent"
)

type AdminRepository struct {
	DB *database.Database
}

func NewAdminRepository(database *database.Database) *AdminRepository {
	return &AdminRepository{
		database,
	}
}

func (repo *AdminRepository) Client() *ent.AdminClient {
	return repo.DB.Ent.Admin
}
