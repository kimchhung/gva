package code_gen

var service_template = `package service

import (
	"gva/app/module/{{.EntitySnake}}/repo"
	"gva/app/module/{{.EntitySnake}}/request"

	"gva/internal/ent"
	"gva/internal/ent/{{.EntityAllLower}}"
	"context"
)

type {{.Entity}}Service struct {
	repo *repo.{{.Entity}}Repository
}

func New{{.Entity}}Service(repo *repo.{{.Entity}}Repository) *{{.Entity}}Service {
	return &{{.Entity}}Service{
		repo: repo,
	}
}

func (s *{{.Entity}}Service) Get{{.Entity}}s(ctx context.Context) ([]*ent.{{.Entity}}, error) {
	return s.repo.Client().Query().Order(ent.Asc({{.EntityAllLower}}.FieldID)).All(ctx)
}

func (s *{{.Entity}}Service) Get{{.Entity}}ByID(ctx context.Context, id int) (*ent.{{.Entity}}, error) {
	return s.repo.Client().Query().Where({{.EntityAllLower}}.IDEQ(id)).First(ctx)
}

func (s *{{.Entity}}Service) Create{{.Entity}}(ctx context.Context, request request.{{.Entity}}Request) (*ent.{{.Entity}}, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *{{.Entity}}Service) Update{{.Entity}}(ctx context.Context, id int, request request.{{.Entity}}Request) (*ent.{{.Entity}}, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *{{.Entity}}Service) Delete{{.Entity}}(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}

`
