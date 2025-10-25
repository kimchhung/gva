package repository

import (
	"backend/app/share/model"
	"backend/core/database"

	"go.uber.org/fx"
)

type AdminRoleRepo struct {
	IBaseRepository[model.AdminRole]
	*database.Database
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewAdminRoleRepo))
}

func NewAdminRoleRepo(db *database.Database) *AdminRoleRepo {
	return &AdminRoleRepo{
		NewBaseRepository[model.AdminRole](db.DB),
		db,
	}
}
