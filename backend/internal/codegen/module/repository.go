package module_template

var Repository = `package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/pagi"
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

// For query
func (r *{{.EntityPascal}}Repository) Q(opts ...pagi.InterceptorOption) *ent.{{.EntityPascal}}Query {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}

`
