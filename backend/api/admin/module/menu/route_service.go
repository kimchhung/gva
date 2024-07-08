package menu

import (
	"github.com/gva/api/admin/module/menu/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/utils/pagi"
	"github.com/gva/utils/routeutil"

	"context"

	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/menu"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{
		repo: repo,
	}
}

func (s *MenuService) toDto(value ...*ent.Menu) []*dto.MenuResponse {
	list := make([]*dto.MenuResponse, len(value))
	for i, v := range value {
		list[i] = &dto.MenuResponse{Menu: v}
	}

	return list
}

func (s *MenuService) Paginate(ctx context.Context, p *dto.MenuPaginateRequest) ([]*dto.MenuResponse, *pagi.Meta, error) {
	query := s.repo.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	)

	meta := &pagi.Meta{
		Limit:  p.Limit,
		Offset: p.Offset,
	}

	if p.IsCount {
		total := query.CountX(ctx)
		meta.Total = &total
	}

	list := query.Modify(pagi.WithLimitOffset(p.Limit, p.Offset)).AllX(ctx)
	if p.IsGroupNested {
		list = routeutil.GroupRouteToNested(list)
	}

	return s.toDto(list...), meta, nil
}

func (s *MenuService) GetRouteByID(ctx context.Context, id xid.ID) (*dto.MenuResponse, error) {
	data, err := s.repo.Q().Where(menu.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *MenuService) CreateRoute(ctx context.Context, r *dto.MenuRequest) (*dto.MenuResponse, error) {
	data, err := s.repo.C().Create().
		SetComponent(r.Component).
		SetPath(r.Path).
		SetIsEnable(r.IsEnable).
		SetMeta(r.Meta).
		SetName(r.Name).
		SetType(r.Type).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *MenuService) UpdateRoute(ctx context.Context, id xid.ID, r *dto.MenuRequest) (*dto.MenuResponse, error) {
	data, err := s.repo.C().UpdateOneID(id).
		SetComponent(r.Component).
		SetPath(r.Path).
		SetIsEnable(r.IsEnable).
		SetMeta(r.Meta).
		SetName(r.Name).
		SetType(r.Type).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *MenuService) DeleteRoute(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
