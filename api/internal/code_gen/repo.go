package code_gen

var repo_template = `package repo

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

func (repo *{{.Entity}}Repository) Client() *ent.{{.Entity}}Client {
	return repo.DB.Ent.{{.Entity}}
}

`
