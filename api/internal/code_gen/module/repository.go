package module_template

var Repository = `package repository

import (
	"gva/internal/bootstrap/database"
	"gva/internal/ent"
)

type {{.Entity}}Repository struct {
	DB *database.Database
}

func New{{.Entity}}Repository(database *database.Database) *{{.Entity}}Repository {
	return &{{.Entity}}Repository{
		database,
	}
}

func (r *{{.Entity}}Repository) Client() *ent.{{.Entity}}Client {
	return r.DB.Ent.{{.Entity}}
}

`
