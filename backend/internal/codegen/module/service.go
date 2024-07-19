package module_template

var Service = `package {{.EntityAllLower}}

import (
	"github.com/gva/app/common/repository"
	"github.com/gva/api/admin/module/{{.EntityAllLower}}/dto"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/{{.EntityAllLower}}"
	"context"
)

type {{.EntityPascal}}Service struct {
	repo *repository.{{.EntityPascal}}Repository
}

func New{{.EntityPascal}}Service(repository *repository.{{.EntityPascal}}Repository) *{{.EntityPascal}}Service {
	return &{{.EntityPascal}}Service{
		repo: repository,
	}
}

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}s(ctx context.Context) ([]*ent.{{.EntityPascal}}, error) {
	return s.repo.C().Query().Order(ent.Asc({{.EntityAllLower}}.FieldID)).All(ctx)
}

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}ByID(ctx context.Context, id xid.ID) (*ent.{{.EntityPascal}}, error) {
	return s.repo.C().Query().Where({{.EntityAllLower}}.IDEQ(id)).First(ctx)
}

func (s *{{.EntityPascal}}Service) Create{{.EntityPascal}}(ctx context.Context, payload *dto.{{.EntityPascal}}Request) (*ent.{{.EntityPascal}}, error) {
	return s.repo.C().Create().
		Save(ctx)
}

func (s *{{.EntityPascal}}Service) Update{{.EntityPascal}}(ctx context.Context, id xid.ID, payload *dto.{{.EntityPascal}}Request) (*ent.{{.EntityPascal}}, error) {
	return s.repo.C().UpdateOneID(id).
		Save(ctx)
}

func (s *{{.EntityPascal}}Service) Delete{{.EntityPascal}}(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}

`
