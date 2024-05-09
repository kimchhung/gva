package service

import (
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/dto"
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/repository"
	"github.com/kimchhung/gva/extra/utils/pagi"
	"github.com/kimchhung/gva/extra/utils/routeutil"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/route"
	"github.com/kimchhung/gva/extra/internal/rql"
)

type RouteService struct {
	repo *repository.RouteRepository
}

func NewRouteService(repository *repository.RouteRepository) *RouteService {
	return &RouteService{
		repo: repository,
	}
}

func (s *RouteService) Paginate(ctx context.Context, p *rql.Params) ([]*ent.Route, *pagi.Meta, error) {
	list := s.repo.RQL(p).WithRoles().AllX(ctx)
	list = routeutil.GroupRouteToNested(list)

	meta := &pagi.Meta{
		Total:  s.repo.RQLCount(ctx, p),
		Limit:  p.Limit,
		Offset: p.Offset,
	}

	return list, meta, nil
}

func (s *RouteService) GetRouteByID(ctx context.Context, id int) (*ent.Route, error) {
	return s.repo.C().Query().Where(route.IDEQ(id)).First(ctx)
}

func (s *RouteService) CreateRoute(ctx context.Context, r *dto.RouteRequest) (*dto.RouteResponse, error) {
	route, err := s.repo.C().Create().
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

	return &dto.RouteResponse{Route: route}, nil
}

func (s *RouteService) UpdateRoute(ctx context.Context, id int, r *dto.RouteRequest) (*dto.RouteResponse, error) {
	route, err := s.repo.C().UpdateOneID(id).
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

	return &dto.RouteResponse{Route: route}, nil
}

func (s *RouteService) DeleteRoute(ctx context.Context, id int) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
