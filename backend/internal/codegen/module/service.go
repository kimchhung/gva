package module_template

var Service = `package {{.EntityAllLower}}

import (
	"context"
	"strings"

	"github.com/gva/api/admin/module/{{.EntityAllLower}}/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/{{.EntityAllLower}}"
	"github.com/gva/utils"
	"github.com/gva/utils/pagi"
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
	for i, v := range value {
		// todo: map value to response value here
		list[i] = &dto.{{.EntityPascal}}Response{v}
	}
	return list
}

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}s(ctx context.Context, p *dto.{{.EntityPascal}}PagedRequest) ([]*dto.{{.EntityPascal}}Response, *pagi.Meta, error) {
	if p.Selects == "" {
		p.Selects = "count,list"
	}
	query := s.repo.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	)
	countQuery := query.Clone()
	listQuery := query.Modify(pagi.WithLimitOffset(p.Limit, p.Offset))
	metaCh := utils.Async(func() *pagi.Meta {
		m := &pagi.Meta{Limit: p.Limit, Offset: p.Offset}
		if !strings.Contains(p.Selects, "count") {
			return m
		}
		m.Total = countQuery.CountX(ctx)
		return m
	})
	listCh := utils.Async(func() []*ent.{{.EntityPascal}} {
		if !strings.Contains(p.Selects, "list") {
			return nil
		}
		return listQuery.AllX(ctx)
	})
	list := <-listCh
	meta := <-metaCh
	return s.toDto(list...), meta, nil
}

func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}ByID(ctx context.Context, id xid.ID) (*dto.{{.EntityPascal}}Response, error) {
	data, err := s.repo.Q().Where({{.EntityAllLower}}.IDEQ(id)).First(ctx)
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
