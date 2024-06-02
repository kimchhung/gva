package route

import (
	"github.com/kimchhung/gva/extra/api/admin/module/route/dto"
	"github.com/kimchhung/gva/extra/app/common/repository"
	"github.com/kimchhung/gva/extra/utils/pagi"
	"github.com/kimchhung/gva/extra/utils/routeutil"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/route"
)

type RouteService struct {
	repo *repository.RouteRepository
}

func NewRouteService(repo *repository.RouteRepository) *RouteService {
	return &RouteService{
		repo: repo,
	}
}

func (s *RouteService) toDto(value ...*ent.Route) []*dto.RouteResponse {
	list := make([]*dto.RouteResponse, len(value))
	for i, v := range value {
		list[i] = &dto.RouteResponse{
			Route: v,
		}
	}

	return list
}

func (s *RouteService) Paginate(ctx context.Context, p *dto.RoutePaginateRequest) ([]*dto.RouteResponse, *pagi.Meta, error) {
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

func (s *RouteService) GetRouteByID(ctx context.Context, id int) (*dto.RouteResponse, error) {
	data, err := s.repo.Q().Where(route.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *RouteService) CreateRoute(ctx context.Context, r *dto.RouteRequest) (*dto.RouteResponse, error) {
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

func (s *RouteService) UpdateRoute(ctx context.Context, id int, r *dto.RouteRequest) (*dto.RouteResponse, error) {
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

func (s *RouteService) DeleteRoute(ctx context.Context, id int) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
