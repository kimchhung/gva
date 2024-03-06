package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
)

type AdminRepository struct {
	db *database.Database
}

func NewAdminRepository(database *database.Database) *AdminRepository {
	return &AdminRepository{
		database,
	}
}

func (r *AdminRepository) C() *ent.AdminClient {
	return r.db.Admin
}

func (r *AdminRepository) DB() *database.Database {
	return r.db
}
