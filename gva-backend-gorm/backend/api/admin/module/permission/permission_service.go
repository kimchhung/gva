package permission

import (
	"context"
	"backend/api/admin/module/permission/dto"
	"backend/app/common/model"
	repository "backend/app/common/repository"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/utils"
)

type PermissionService struct {
	repo *repository.PermissionRepo
}

// NewPermissionService initializes a new PermissionService with a JwtService and a UserStore.
func NewPermissionService(
	repo *repository.PermissionRepo,
) *PermissionService {
	return &PermissionService{
		repo: repo,
	}
}

// CreatePermission creates a new Permission.
func (s *PermissionService) CreatePermission(ctx context.Context, p *dto.CreatePermissionRequest) (*dto.PermissionResponse, error) {
	body := utils.MustCopy(new(model.Permission), p)
	created, err := s.repo.Create(ctx, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.PermissionResponse), created), nil
}

// GetPermission gets a Permission by ID.
func (s *PermissionService) GetPermission(ctx context.Context, id uint) (*dto.PermissionResponse, error) {
	permission, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.PermissionResponse), permission), nil
}

// UpdatePermission updates a Permission.
func (s *PermissionService) UpdatePermission(ctx context.Context, id uint, p *dto.UpdatePermissionRequest) (*dto.PermissionResponse, error) {
	body := utils.MustCopy(new(model.Permission), p)
	updated, err := s.repo.UpdateById(ctx, id, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.PermissionResponse), updated), nil
}

// UpdatePermission updates a Permission.
func (s *PermissionService) UpdatePatchPermission(ctx context.Context, id uint, p *dto.UpdatePatchPermissionRequest) (*dto.PermissionResponse, error) {
	body := utils.MustCopy(new(model.Permission), p)
	updated, err := s.repo.UpdateById(ctx, id, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.PermissionResponse), updated), nil
}

// DeletePermission deletes a Permission by ID.
func (s *PermissionService) DeletePermission(ctx context.Context, id uint) error {
	return s.repo.DeleteById(ctx, id)
}

// GetPermissions gets all Permissions.
func (s *PermissionService) GetPermissions(ctx context.Context, query *dto.GetManyQuery) ([]dto.PermissionResponse, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.Ignore(),
		"createdAt": gormq.ToSnake(),
	})

	searchColumns := columnMap.Pick(
		"id",
	).Values()

	resp, respMeta := pagi.PrepareResponse[dto.PermissionResponse](&query.QueryDto)

	err := s.repo.GetManyAndCount(ctx, &resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.Where(gormq.WithFilters(query.Filters, columnMap)),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
	)

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}
