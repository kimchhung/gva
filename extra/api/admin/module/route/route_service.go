package route

import (
	"github.com/kimchhung/gva/extra/api/admin/module/route/dto"
	dbr "github.com/kimchhung/gva/extra/app/common/repository"
	"github.com/kimchhung/gva/extra/utils/pagi"
	"github.com/kimchhung/gva/extra/utils/routeutil"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/route"
	"github.com/kimchhung/gva/extra/internal/rql"
)

type RouteService struct {
	route_r *dbr.RouteRepository
}

func NewRouteService(route_r *dbr.RouteRepository) *RouteService {
	return &RouteService{
		route_r: route_r,
	}
}

func (s *RouteService) RQL(ctx context.Context, p *rql.Params) (any, error) {
	data := []any{}

	s.route_r.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	).Modify().Where(
		route.HasChildrenWith(route.TypeIn("menu")),
	).WithChildren().AllX(ctx)

	return data, nil
}

func (s *RouteService) Paginate(ctx context.Context, p *dto.RoutePaginateRequest) ([]*ent.Route, *pagi.Meta, error) {
	q := s.route_r.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	).Modify()

	total := q.CountX(ctx)
	list := q.Modify(pagi.WithLimitOffset(p.Limit, p.Offset)).AllX(ctx)

	meta := &pagi.Meta{
		Limit:  p.Limit,
		Offset: p.Offset,
		Total:  total,
	}

	if p.IsGroupNested {
		list = routeutil.GroupRouteToNested(list)
	}

	return list, meta, nil
}

func (s *RouteService) GetRouteByID(ctx context.Context, id int) (*ent.Route, error) {
	return s.route_r.C().Query().Where(route.IDEQ(id)).First(ctx)
}

func (s *RouteService) CreateRoute(ctx context.Context, r *dto.RouteRequest) (*dto.RouteResponse, error) {
	route, err := s.route_r.C().Create().
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
	route, err := s.route_r.C().UpdateOneID(id).
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
	return s.route_r.C().DeleteOneID(id).Exec(ctx)
}
