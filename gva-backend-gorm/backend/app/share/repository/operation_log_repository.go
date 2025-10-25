package repository

import (
	"backend/app/share/model"
	"backend/core/database"

	"go.uber.org/fx"
)

type OperationLogRepo struct {
	IBaseRepository[model.OperationLog]
	*database.Database
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewOperationLogRepo))
}

func NewOperationLogRepo(db *database.Database) *OperationLogRepo {
	return &OperationLogRepo{
		NewBaseRepository[model.OperationLog](db.DB),
		db,
	}
}
