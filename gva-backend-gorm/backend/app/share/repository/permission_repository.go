package repository

import (
	"backend/app/share/model"
	"backend/core/database"

	"go.uber.org/fx"
)

type PermissionRepo struct {
	IBaseRepository[model.Permission]
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewPermissionRepo))
}

func NewPermissionRepo(db *database.Database) *PermissionRepo {
	return &PermissionRepo{
		NewBaseRepository[model.Permission](db.DB),
	}
}
