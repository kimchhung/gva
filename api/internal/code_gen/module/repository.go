package module_template

var Repository = `package repository

import (
	"github.com/kimchhung/gva/internal/bootstrap/database"
	"github.com/kimchhung/gva/internal/ent"
)

type {{.EntityPascal}}Repository struct {
	DB *database.Database
}

func New{{.EntityPascal}}Repository(database *database.Database) *{{.EntityPascal}}Repository {
	return &{{.EntityPascal}}Repository{
		database,
	}
}

func (r *{{.EntityPascal}}Repository) Client() *ent.{{.EntityPascal}}Client {
	return r.DB.Ent.{{.EntityPascal}}
}

`
