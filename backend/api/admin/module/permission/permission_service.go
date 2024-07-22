package permission

import (
	"context"
	"strings"

	"github.com/gva/api/admin/module/permission/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/utils"
	"github.com/gva/utils/pagi"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repository *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		repo: repository,
	}
}

func (s *PermissionService) mapDto(v *ent.Permission) *dto.PermissionResponse {
	return &dto.PermissionResponse{
		ID:    v.ID,
		Group: v.Group,
		Name:  v.Name,
		Scope: v.Scope,
		Order: v.Order,
	}
}

func (s *PermissionService) toDto(value ...*ent.Permission) []*dto.PermissionResponse {
	list := make([]*dto.PermissionResponse, len(value))
	for i, v := range value {
		list[i] = s.mapDto(v)
	}
	return list
}

func (s *PermissionService) GetPermissions(ctx context.Context, p *dto.PermissionPagedRequest) ([]*dto.PermissionResponse, *pagi.Meta, error) {
	if p.Selects == "" {
		p.Selects = "count,list"
	}
	query := s.repo.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	)

	countQuery := query.Clone()
	listQuery := query.Modify(pagi.WithLimitOffset(p.Limit, p.Offset)).Select(
		permission.FieldID,
		permission.FieldGroup,
		permission.FieldOrder,
		permission.FieldScope,
		permission.FieldName,
	)
	metaCh := utils.Async(func() *pagi.Meta {
		m := &pagi.Meta{Limit: p.Limit, Offset: p.Offset}
		if !strings.Contains(p.Selects, "count") {
			return m
		}
		m.Total = countQuery.CountX(ctx)
		return m
	})
	listCh := utils.Async(func() []*ent.Permission {
		if !strings.Contains(p.Selects, "list") {
			return nil
		}
		return listQuery.AllX(ctx)
	})
	list := <-listCh
	meta := <-metaCh
	return s.toDto(list...), meta, nil
}

func (s *PermissionService) GetPermissionByID(ctx context.Context, id xid.ID) (*dto.PermissionResponse, error) {
	data, err := s.repo.Q().Where(permission.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(data)[0], nil
}

func (s *PermissionService) CreatePermission(ctx context.Context, body *dto.PermissionRequest) (*dto.PermissionResponse, error) {
	create := s.repo.C().Create()
	create.SetName(body.Name)
	create.SetScope(body.Name)
	create.SetGroup(body.Group)
	create.SetOrder(body.Order)
	created, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(created)[0], nil
}

func (s *PermissionService) UpdatePermission(ctx context.Context, id xid.ID, body *dto.PermissionRequest) (*dto.PermissionResponse, error) {
	update := s.repo.C().UpdateOneID(id)
	update.SetName(body.Name)
	update.SetScope(body.Scope)
	update.SetGroup(body.Group)
	update.SetOrder(body.Order)
	updated, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(updated)[0], nil
}

func (s *PermissionService) DeletePermission(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
