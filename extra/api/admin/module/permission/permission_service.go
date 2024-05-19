package permission

import (
	"context"

	repository "github.com/kimchhung/gva/extra/app/common/repository"
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
