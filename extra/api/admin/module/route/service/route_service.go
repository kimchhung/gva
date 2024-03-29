package service

import (
	"github.com/kimchhung/gva/extra/api/admin/module/route/dto"
	"github.com/kimchhung/gva/extra/api/admin/module/route/repository"
	"github.com/kimchhung/gva/extra/utils/routeutil"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/route"
)

type RouteService struct {
	repo *repository.RouteRepository
}

func NewRouteService(repository *repository.RouteRepository) *RouteService {
	return &RouteService{
		repo: repository,
	}
}

func (s *RouteService) GetNestedRoutes(ctx context.Context) ([]*ent.Route, error) {
	flatRoutes, err := s.repo.C().Query().Order(ent.Asc(route.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}

	return routeutil.GroupRouteToNested(flatRoutes), nil
}

func (s *RouteService) GetRouteByID(ctx context.Context, id int) (*ent.Route, error) {
	return s.repo.C().Query().Where(route.IDEQ(id)).First(ctx)
}

func (s *RouteService) CreateRoute(ctx context.Context, payload *dto.RouteRequest) (*ent.Route, error) {
	return s.repo.C().Create().
		Save(ctx)
}

func (s *RouteService) UpdateRoute(ctx context.Context, id int, payload *dto.RouteRequest) (*ent.Route, error) {
	return s.repo.C().UpdateOneID(id).
		Save(ctx)
}

func (s *RouteService) DeleteRoute(ctx context.Context, id int) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
