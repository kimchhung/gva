package repository

import (
	"backend/app/share/model"
	"backend/core/database"

	"go.uber.org/fx"
)

type TodoRepo struct {
	IBaseRepository[model.Todo]
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewTodoRepo))
}

func NewTodoRepo(db *database.Database) *TodoRepo {
	return &TodoRepo{
		NewBaseRepository[model.Todo](db.DB),
	}
}
