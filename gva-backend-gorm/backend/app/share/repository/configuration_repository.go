package repository

import (
	"backend/app/share/model"
	"backend/core/database"

	"go.uber.org/fx"
)

type ConfigurationRepo struct {
	IBaseRepository[model.Configuration]
	*database.Database
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewConfigurationRepo))
}

func NewConfigurationRepo(db *database.Database) *ConfigurationRepo {
	return &ConfigurationRepo{
		NewBaseRepository[model.Configuration](db.DB),
		db,
	}
}
