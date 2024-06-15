package module_template

var Repository = `package repository

import (
	"github.com/kimchhung/gva/backend-echo/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend-echo/internal/ent"
)

type {{.EntityPascal}}Repository struct {
	db *database.Database
}

func New{{.EntityPascal}}Repository(database *database.Database) *{{.EntityPascal}}Repository {
	return &{{.EntityPascal}}Repository{
		database,
	}
}

func (r *{{.EntityPascal}}Repository) C() *ent.{{.EntityPascal}}Client {
	return r.db.{{.EntityPascal}}
}

func (r *{{.EntityPascal}}Repository) DB() *database.Database {
	return r.db
}

`
