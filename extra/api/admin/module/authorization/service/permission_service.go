package service

import (
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/repository"

	"context"

	"github.com/kimchhung/gva/extra/internal/ent"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repository *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		repo: repository,
	}
}

func (s *PermissionService) AllPermissions(ctx context.Context) ([]*ent.Permission, error) {
	return s.repo.C().Query().All(ctx)
}
