package service

import (
	"github.com/kimchhung/gva/app/module/permission/repository"
	"github.com/kimchhung/gva/app/module/permission/dto"

	"github.com/kimchhung/gva/internal/ent"
	"github.com/kimchhung/gva/internal/ent/permission"
	"context"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repository *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		repo: repository,
	}
}

func (s *PermissionService) GetPermissions(ctx context.Context) ([]*ent.Permission, error) {
	return s.repo.Client().Query().Order(ent.Asc(permission.FieldID)).All(ctx)
}

func (s *PermissionService) GetPermissionByID(ctx context.Context, id int) (*ent.Permission, error) {
	return s.repo.Client().Query().Where(permission.IDEQ(id)).First(ctx)
}

func (s *PermissionService) CreatePermission(ctx context.Context, request dto.PermissionRequest) (*ent.Permission, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *PermissionService) UpdatePermission(ctx context.Context, id int, request dto.PermissionRequest) (*ent.Permission, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *PermissionService) DeletePermission(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}

