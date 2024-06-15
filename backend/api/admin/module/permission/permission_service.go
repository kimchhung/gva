package permission

import (
	"context"

	"github.com/kimchhung/gva/backend-echo/app/common/repository"
	"github.com/kimchhung/gva/backend-echo/internal/ent"
	"github.com/kimchhung/gva/backend/api/admin/module/permission/dto"
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
