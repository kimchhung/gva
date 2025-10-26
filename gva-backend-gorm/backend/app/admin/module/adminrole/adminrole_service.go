package adminrole

import (
	admincontext "backend/app/admin/context"
	adminerror "backend/app/admin/error"
	"backend/app/admin/module/adminrole/dto"
	"backend/app/share/constant"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	coreerror "backend/core/error"
	"backend/core/utils"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type AdminRoleService struct {
	repo         *repository.AdminRoleRepo
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
		repo:         repo,
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

	created, err := s.repo.Create(ctx, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.AdminRoleResponse), created), nil
}

// GetAdminRole gets a AdminRole by ID.
func (s *AdminRoleService) GetAdminRole(ctx context.Context, id uint) (*dto.AdminRoleResponse, error) {
	adminrole, err := s.repo.GetById(ctx, id, func(q *gorm.DB) *gorm.DB {
		return q.Preload("Permissions")
	})
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.AdminRoleResponse), adminrole), nil
}

func (s *AdminRoleService) lockTargetForUpdate(ctx context.Context, id uint, out *model.AdminRole, selects ...gormq.Option) gormq.Tx {
	return func(tx *gorm.DB) error {
		found, err := s.repo.Tx(tx).GetById(ctx, id, gormq.Multi(selects...), gormq.WithLockUpdate())
		if err != nil {
			// Check if the error is a not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(coreerror.ErrNotFound)
			}

			return err
		}

		if out != nil && found != nil {
			*out = *found
		}

		return nil
	}
}

// UpdateAdminRole updates a AdminRole.
func (s *AdminRoleService) UpdateAdminRole(ctx context.Context, id uint, dtoReq *dto.UpdateAdminRoleRequest) (res *dto.AdminRoleResponse, err error) {
	// super admin cannot be updated
	if id == constant.RoleIdSuperAdmin {
		return nil, coreerror.ErrForbidden
	}

	var (
		target model.AdminRole
	)

	err = s.repo.MultiTransaction(
		s.lockTargetForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(&target, dtoReq)
			body.ID = target.ID
			body.Permissions = s.getPermissionsByScope(ctx, dtoReq.Permissions)

			err := tx.Model(body).Association("Permissions").Replace(body.Permissions)
			if err != nil {
				return err
			}

			updated, err := s.repo.Tx(tx).UpdateById(ctx, target.ID, body, gormq.WithSelect("name", "description"))
			if err != nil {
				return err
			}

			res = utils.MustCopy(new(dto.AdminRoleResponse), updated)
			return nil
		},
	)

	return
}

// UpdateAdminRole updates a AdminRole.
func (s *AdminRoleService) UpdatePatchAdminRole(ctx context.Context, id uint, p *dto.UpdatePatchAdminRoleRequest) (resp map[string]any, err error) {
	// super admin cannot be updated
	if id == constant.RoleIdSuperAdmin {
		return nil, coreerror.ErrForbidden
	}

	var (
		target model.AdminRole
	)

	err = s.repo.MultiTransaction(
		s.lockTargetForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
				"status": gormq.Ignore(),
			})

			dbCols, res := utils.StructToMap(p, columnMap)
			if err := s.repo.DB().Model(&target).
				Scopes(
					gormq.Where(gormq.Equal("id", id)),
				).
				Updates(dbCols).Error; err != nil {
				return err
			}

			resp = res
			return nil
		})

	return
}

// DeleteAdminRole deletes a AdminRole by ID.
func (s *AdminRoleService) DeleteAdminRole(ctx context.Context, id uint) error {
	if id == constant.RoleIdSuperAdmin {
		return coreerror.ErrForbidden
	}

	// check if role is in use
	var count int64 = 0
	err := s.repo.DB().Table("admin_admin_roles").
		Joins("inner join admins on admins.id = admin_admin_roles.admin_id").
		Where("admin_admin_roles.admin_role_id = ?", id).
		Where("admins.deleted_at = 0").
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return adminerror.ErrAdminRoleIsInUse
	}

	return s.repo.DeleteById(ctx, id)
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

	err := s.repo.GetManyAndCount(ctx, &resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.Where(gormq.WithFilters(query.Filters, columnMap)),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
		func(q *gorm.DB) *gorm.DB {
			adminCtx := admincontext.MustAdminContext(ctx)
			if adminCtx.IsSuperAdmin() {
				return q
			}
			return q.Where("id != ?", constant.RoleIdSuperAdmin)
		},
	)

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}
