package service

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/extra/app/common/contexts"
	"github.com/kimchhung/gva/extra/app/module/dashboard/admin/dto"
	"github.com/kimchhung/gva/extra/app/module/dashboard/admin/repository"

	"github.com/kimchhung/gva/extra/utils/routeutil"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
	"github.com/kimchhung/gva/extra/internal/ent/role"
	"github.com/kimchhung/gva/extra/internal/ent/route"
	"github.com/kimchhung/gva/extra/internal/rql"
)

type AdminService struct {
	repo *repository.AdminRepository
}

func NewAdminService(repository *repository.AdminRepository) *AdminService {
	return &AdminService{
		repo: repository,
	}
}

func (s *AdminService) GetAdmins(ctx context.Context, p *rql.Params) ([]*ent.Admin, error) {
	return s.repo.C().Query().Where(func(s *sql.Selector) {
		s.Where(sql.ExprP(p.FilterExp.String(), p.FilterArgs...))
	}).
		Limit(p.Limit).
		Offset(p.Offset).
		All(ctx)
}

func (s *AdminService) GetAdminByID(ctx context.Context, id int) (*ent.Admin, error) {
	return s.repo.C().Query().Where(admin.IDEQ(id)).First(ctx)
}

func (s *AdminService) CreateAdmin(ctx context.Context, request dto.AdminRequest) (*ent.Admin, error) {
	return s.repo.C().Create().
		SetPassword(request.Password).
		SetUsername(request.Username).
		SetDisplayName(request.DisplayName).
		Save(ctx)
}

func (s *AdminService) UpdateAdmin(ctx context.Context, id int, request dto.AdminRequest) (*ent.Admin, error) {
	return s.repo.C().UpdateOneID(id).
		SetDisplayName(request.DisplayName).
		SetUsername(request.Username).
		Save(ctx)
}

func (s *AdminService) DeleteAdmin(ctx context.Context, id int) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}

func (s *AdminService) GetAdminNestedRouteById(ctx context.Context, adminId int) ([]*ent.Route, error) {
	if contexts.MustAdminContext(ctx).IsSuperAdmin() {
		routes, err := s.repo.DB().Route.Query().Where(route.IsEnable(true)).All(ctx)
		if err != nil {
			return nil, err
		}

		return routeutil.GroupRouteToNested(routes), nil
	}

	routes, err := s.repo.DB().Role.Query().
		Where(role.HasAdminsWith(admin.ID(adminId))).
		QueryRoutes().Where(route.IsEnable(true)).All(ctx)

	if err != nil {
		return nil, err
	}

	return routeutil.GroupRouteToNested(routes), nil
}

func (s *AdminService) GetAdminPermissionById(ctx context.Context, adminId int) ([]*ent.Permission, error) {
	routes, err := s.repo.DB().Role.Query().
		Where(
			role.HasAdminsWith(admin.ID(adminId)),
		).
		QueryPermissions().All(ctx)

	if err != nil {
		return nil, err
	}

	return routes, nil
}
