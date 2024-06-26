package permission

import (
	"context"

	"github.com/gva/api/admin/module/permission/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/internal/ent"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repository *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		repo: repository,
	}
}

func (s *PermissionService) toDto(value ...*ent.Permission) []*dto.PermissionResponse {
	list := make([]*dto.PermissionResponse, len(value))
	for i, v := range value {
		list[i] = &dto.PermissionResponse{
			Permission: v,
		}
	}

	return list
}

func (s *PermissionService) AllPermissions(ctx context.Context) ([]*dto.PermissionResponse, error) {
	data, err := s.repo.Q().All(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(data...), nil
}
