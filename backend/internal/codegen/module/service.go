package module_template

var Service = `package {{.EntityAllLower}}

import (
	"context"

	"github.com/gva/api/admin/module/{{.EntityAllLower}}/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/{{.EntityAllLower}}"
)

type {{.EntityPascal}}Service struct {
	repo *repository.{{.EntityPascal}}Repository
}

func New{{.EntityPascal}}Service(repository *repository.{{.EntityPascal}}Repository) *{{.EntityPascal}}Service {
	return &{{.EntityPascal}}Service{
		repo: repository,
	}
}

func (s *{{.EntityPascal}}Service) toDto(value ...*ent.{{.EntityPascal}}) []*dto.{{.EntityPascal}}Response {
	list := make([]*dto.{{.EntityPascal}}Response, len(value))
	for i, _ := range value {
		// todo: map value to response value here
		list[i] = &dto.{{.EntityPascal}}Response{}
	}
	return list
}

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}s(ctx context.Context) ([]*dto.{{.EntityPascal}}Response, error) {
	list, err := s.repo.C().Query().Order(ent.Asc({{.EntityAllLower}}.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(list...), nil
}

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}ByID(ctx context.Context, id xid.ID) (*dto.{{.EntityPascal}}Response, error) {
	data, err := s.repo.C().Query().Where({{.EntityAllLower}}.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(data)[0], nil 
}

func (s *{{.EntityPascal}}Service) Create{{.EntityPascal}}(ctx context.Context, payload *dto.{{.EntityPascal}}Request) (*dto.{{.EntityPascal}}Response, error) {
	create := s.repo.C().Create()
	create.SetName(payload.Name)
	created, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(created)[0], nil
}

func (s *{{.EntityPascal}}Service) Update{{.EntityPascal}}(ctx context.Context, id xid.ID, payload *dto.{{.EntityPascal}}Request) (*dto.{{.EntityPascal}}Response, error) {
	update := s.repo.C().UpdateOneID(id)
	update.SetName(payload.Name)
	updated, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(updated)[0], nil
}

func (s *{{.EntityPascal}}Service) Delete{{.EntityPascal}}(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
`
