package service

import (
	"github.com/kimchhung/gva/extra/app/module/route/repository"
	"github.com/kimchhung/gva/extra/app/module/route/dto"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/route"
	"context"
)

type RouteService struct {
	repo *repository.RouteRepository
}

func NewRouteService(repository *repository.RouteRepository) *RouteService {
	return &RouteService{
		repo: repository,
	}
}

func (s *RouteService) GetRoutes(ctx context.Context) ([]*ent.Route, error) {
	return s.repo.Client().Query().Order(ent.Asc(route.FieldID)).All(ctx)
}

func (s *RouteService) GetRouteByID(ctx context.Context, id int) (*ent.Route, error) {
	return s.repo.Client().Query().Where(route.IDEQ(id)).First(ctx)
}

func (s *RouteService) CreateRoute(ctx context.Context, payload *dto.RouteRequest) (*ent.Route, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *RouteService) UpdateRoute(ctx context.Context, id int, payload *dto.RouteRequest) (*ent.Route, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *RouteService) DeleteRoute(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}

