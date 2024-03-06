package service

import (
	"github.com/kimchhung/gva/extra/app/module/admin/dto"
	"github.com/kimchhung/gva/extra/app/module/admin/repository"
	"github.com/kimchhung/gva/extra/utils/routeutil"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
	"github.com/kimchhung/gva/extra/internal/ent/role"
)

type AdminService struct {
	repo *repository.AdminRepository
}

func NewAdminService(repository *repository.AdminRepository) *AdminService {
	return &AdminService{
		repo: repository,
	}
}

func (s *AdminService) GetAdmins(ctx context.Context) ([]*ent.Admin, error) {
	return s.repo.Client().Query().Order(ent.Asc(admin.FieldID)).All(ctx)
}

func (s *AdminService) GetAdminByID(ctx context.Context, id int) (*ent.Admin, error) {
	return s.repo.Client().Query().Where(admin.IDEQ(id)).First(ctx)
}

func (s *AdminService) CreateAdmin(ctx context.Context, request dto.AdminRequest) (*ent.Admin, error) {
	return s.repo.Client().Create().
		SetPassword(request.Password).
		SetUsername(request.Username).
		SetDisplayName(request.DisplayName).
		Save(ctx)
}

func (s *AdminService) UpdateAdmin(ctx context.Context, id int, request dto.AdminRequest) (*ent.Admin, error) {
	return s.repo.Client().UpdateOneID(id).
		SetDisplayName(request.DisplayName).
		SetUsername(request.Username).
		Save(ctx)
}

func (s *AdminService) DeleteAdmin(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}

func (s *AdminService) GetAdminNestedRouteById(ctx context.Context, adminId int) ([]*ent.Route, error) {
	routes, err := s.repo.DB.Ent.Role.Query().
		Where(role.HasAdminsWith(admin.ID(adminId))).
		QueryRoutes().All(ctx)

	if err != nil {
		return nil, err
	}

	return routeutil.GroupRouteToNested(routes), nil
}

func (s *AdminService) GetAdminPermissionById(ctx context.Context, adminId int) ([]*ent.Permission, error) {
	routes, err := s.repo.DB.Ent.Role.Query().
		Where(
			role.HasAdminsWith(admin.ID(adminId)),
		).
		QueryPermissions().All(ctx)

	if err != nil {
		return nil, err
	}

	return routes, nil
}
