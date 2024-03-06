package service

import (
	"github.com/kimchhung/gva/extra/app/module/permission/dto"
	"github.com/kimchhung/gva/extra/app/module/permission/repository"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/permission"
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
	return s.repo.C().Query().Order(ent.Asc(permission.FieldID)).All(ctx)
}

func (s *PermissionService) GetPermissionByID(ctx context.Context, id int) (*ent.Permission, error) {
	return s.repo.C().Query().Where(permission.IDEQ(id)).First(ctx)
}

func (s *PermissionService) CreatePermission(ctx context.Context, payload *dto.PermissionRequest) (*ent.Permission, error) {
	return s.repo.C().Create().
		Save(ctx)
}

func (s *PermissionService) UpdatePermission(ctx context.Context, id int, payload *dto.PermissionRequest) (*ent.Permission, error) {
	return s.repo.C().UpdateOneID(id).
		Save(ctx)
}

func (s *PermissionService) DeletePermission(ctx context.Context, id int) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
