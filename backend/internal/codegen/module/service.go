package module_template

var Service = `package service

import (
	"github.com/kimchhung/gva/backend/api/admin/module/{{.EntitySnake}}/repository"
	"github.com/kimchhung/gva/backend/api/admin/module/{{.EntitySnake}}/dto"

	"github.com/kimchhung/gva/backend/internal/ent"
	"github.com/kimchhung/gva/backend/internal/ent/{{.EntityAllLower}}"
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

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}ByID(ctx context.Context, id int) (*ent.{{.EntityPascal}}, error) {
	return s.repo.C().Query().Where({{.EntityAllLower}}.IDEQ(id)).First(ctx)
}

func (s *{{.EntityPascal}}Service) Create{{.EntityPascal}}(ctx context.Context, payload *dto.{{.EntityPascal}}Request) (*ent.{{.EntityPascal}}, error) {
	return s.repo.C().Create().
		Save(ctx)
}

func (s *{{.EntityPascal}}Service) Update{{.EntityPascal}}(ctx context.Context, id int, payload *dto.{{.EntityPascal}}Request) (*ent.{{.EntityPascal}}, error) {
	return s.repo.C().UpdateOneID(id).
		Save(ctx)
}

func (s *{{.EntityPascal}}Service) Delete{{.EntityPascal}}(ctx context.Context, id int) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}

`
