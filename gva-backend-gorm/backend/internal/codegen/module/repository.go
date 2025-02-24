package module_template

var Repository = `package repository

import (
	"backend/app/common/model"
	"backend/internal/bootstrap/database"

	"go.uber.org/fx"
)

type {{.EntityPascal}}Repo struct {
	IBaseRepository[model.{{.EntityPascal}}]
	*database.Database
}

func init() {
	dependencies = append(dependencies, fx.Provide(New{{.EntityPascal}}Repo))
}

func New{{.EntityPascal}}Repo(db *database.Database) *{{.EntityPascal}}Repo {
	return &{{.EntityPascal}}Repo{
		NewBaseRepository[model.{{.EntityPascal}}](db.DB),
		db,
	}
}
`
