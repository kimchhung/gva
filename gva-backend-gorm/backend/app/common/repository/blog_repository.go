package repository

import (
	"backend/app/common/model"
	"backend/internal/bootstrap/database"

	"go.uber.org/fx"
)

type BlogRepo struct {
	IBaseRepository[model.Blog]
	*database.Database
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewBlogRepo))
}

func NewBlogRepo(db *database.Database) *BlogRepo {
	return &BlogRepo{
		NewBaseRepository[model.Blog](db.DB),
		db,
	}
}
