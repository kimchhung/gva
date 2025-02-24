package adminrole

import (
	"backend/api/admin/module/adminrole/dto"
	appctx "backend/app/common/context"
	apperror "backend/app/common/error"
	"backend/app/common/model"
	repository "backend/app/common/repository"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/utils"
	"context"
	"strings"

	"gorm.io/gorm"
)

type AdminRoleService struct {
	adminRole_r  *repository.AdminRoleRepo
	admin_r      *repository.AdminRepo
	permission_r *repository.PermissionRepo
}

// NewAdminRoleService initializes a new AdminRoleService with a JwtService and a UserStore.
func NewAdminRoleService(
	repo *repository.AdminRoleRepo,
	adminRepo *repository.AdminRepo,
	permissionRepo *repository.PermissionRepo,
) *AdminRoleService {
	return &AdminRoleService{
		adminRole_r:  repo,
		admin_r:      adminRepo,
		permission_r: permissionRepo,
	}
}

func (s *AdminRoleService) getNameId(name string) string {
	return strings.Replace(strings.TrimSpace(strings.ToLower(name)), " ", "_", -1)
}

func (s *AdminRoleService) getPermissionsByScope(ctx context.Context, scopes []string) []model.Permission {
	permissions, err := s.permission_r.GetMany(ctx, gormq.In("scope", scopes), gormq.WithSelect("id"))
	if err != nil {
		panic(err)
	}
	return permissions
}

// CreateAdminRole creates a new AdminRole.
func (s *AdminRoleService) CreateAdminRole(ctx context.Context, p *dto.CreateAdminRoleRequest) (*dto.AdminRoleResponse, error) {
	body := utils.MustCopy(new(model.AdminRole), p)
	body.BaseModel = model.NewBaseModel()

	body.NameID = s.getNameId(body.Name)
	body.Status = 1

	body.Permissions = s.getPermissionsByScope(ctx, p.Permissions)

	created, err := s.adminRole_r.Create(ctx, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.AdminRoleResponse), created), nil
}

// GetAdminRole gets a AdminRole by ID.
func (s *AdminRoleService) GetAdminRole(ctx context.Context, id uint) (*dto.AdminRoleResponse, error) {
	adminrole, err := s.adminRole_r.GetById(ctx, id, func(q *gorm.DB) *gorm.DB {
		return q.Preload("Permissions")
	})
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.AdminRoleResponse), adminrole), nil
}

// UpdateAdminRole updates a AdminRole.
func (s *AdminRoleService) UpdateAdminRole(ctx context.Context, id uint, p *dto.UpdateAdminRoleRequest) (res *dto.AdminRoleResponse, err error) {

	// super admin cannot be updated
	if id == appctx.RoleIdSuperAdmin {
		return nil, apperror.ErrForbidden
	}

	err = s.adminRole_r.Transaction(func(tx *gorm.DB) error {
		body := utils.MustCopy(new(model.AdminRole), p)
		body.ID = id
		body.Permissions = s.getPermissionsByScope(ctx, p.Permissions)

		err := tx.Model(body).Association("Permissions").Replace(body.Permissions)
		if err != nil {
			return err
		}

		updated, err := s.adminRole_r.Tx(tx).UpdateById(ctx, id, body)
		if err != nil {
			return err
		}

		res = utils.MustCopy(new(dto.AdminRoleResponse), updated)
		return nil
	})

	return
}

// UpdateAdminRole updates a AdminRole.
func (s *AdminRoleService) UpdatePatchAdminRole(ctx context.Context, id uint, p *dto.UpdatePatchAdminRoleRequest) (map[string]any, error) {
	// super admin cannot be updated
	if id == appctx.RoleIdSuperAdmin {
		return nil, apperror.ErrForbidden
	}

	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"status": gormq.Ignore(),
	})

	dbCols, resp := utils.StructToMap(p, columnMap)
	err := s.adminRole_r.Model(&model.AdminRole{}).
		Scopes(
			gormq.Where(gormq.Equal("id", id)),
		).
		Updates(dbCols).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteAdminRole deletes a AdminRole by ID.
func (s *AdminRoleService) DeleteAdminRole(ctx context.Context, id uint) error {
	if id == appctx.RoleIdSuperAdmin {
		return apperror.ErrForbidden
	}

	// check if role is in use
	var count int64 = 0
	err := s.adminRole_r.Table("admin_admin_roles").
		Joins("inner join admins on admins.id = admin_admin_roles.admin_id").
		Where("admin_admin_roles.admin_role_id = ?", id).
		Where("admins.deleted_at = 0").
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return apperror.ErrAdminRoleIsInUse
	}

	return s.adminRole_r.DeleteById(ctx, id)
}

// GetAdminRoles gets all AdminRoles.
func (s *AdminRoleService) GetAdminRoles(ctx context.Context, query *dto.GetManyQuery) ([]dto.AdminRoleResponse, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.Ignore(),
		"createdAt": gormq.ToSnake(),
		"status":    gormq.Ignore(),
		"name":      gormq.Ignore(),
	})

	searchColumns := columnMap.Pick(
		"id",
		"name",
	).Values()

	resp, respMeta := pagi.PrepareResponse[dto.AdminRoleResponse](&query.QueryDto)

	err := s.adminRole_r.GetManyAndCount(ctx, &resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.Where(gormq.WithFilters(query.Filters, columnMap)),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
		func(q *gorm.DB) *gorm.DB {
			adminCtx := appctx.MustAdminContext(ctx)
			if adminCtx.IsSuperAdmin() {
				return q
			}
			return q.Where("id != ?", appctx.RoleIdSuperAdmin)
		},
	)

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}
