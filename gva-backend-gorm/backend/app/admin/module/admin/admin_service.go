package admin

import (
	"context"
	"errors"

	admincontext "backend/app/admin/context"
	adminerror "backend/app/admin/error"
	"backend/app/admin/module/admin/dto"
	"backend/app/share/constant"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	"backend/app/share/service"
	coreerror "backend/core/error"
	"backend/core/utils"
	"backend/internal/gormq"
	"backend/internal/pagi"

	"gorm.io/gorm"
)

type AdminService struct {
	repo       *repository.AdminRepo
	password_s *service.PasswordService
	totp_s     *service.TOTPService
}

// NewAdminService initializes a new AdminService with a JwtService and a UserStore.
func NewAdminService(
	repo *repository.AdminRepo,
	password_s *service.PasswordService,
	totp_s *service.TOTPService,
) *AdminService {
	return &AdminService{
		repo:       repo,
		password_s: password_s,
		totp_s:     totp_s,
	}
}

// CreateAdmin creates a new Admin.
func (s *AdminService) CreateAdmin(ctx context.Context, p *dto.CreateAdminRequest) (*dto.AdminResponse, error) {
	body := utils.MustCopy(new(model.Admin), p)
	body.BaseModel = model.NewBaseModel()

	s.verifySuperAdminRoleUsage(ctx, body.Roles)

	passwordHash, err := s.password_s.HashPassword(p.Password)
	if err != nil {
		return nil, err
	}

	body.PasswordHash = passwordHash
	body.Status = 1

	created, err := s.repo.Create(ctx, body)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, adminerror.ErrUsernameExists
		}
		return nil, err
	}

	return utils.MustCopy(new(dto.AdminResponse), created), nil
}

// GetAdmin gets a Admin by ID.
func (s *AdminService) GetAdmin(ctx context.Context, id uint) (*dto.AdminResponse, error) {
	admin, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.AdminResponse), admin), nil
}

func (s *AdminService) lockTargetForUpdate(ctx context.Context, id uint, out *model.Admin, selects ...gormq.Option) gormq.Tx {
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

// UpdateAdmin updates a Admin.
func (s *AdminService) UpdateAdmin(ctx context.Context, id uint, dtoReq *dto.UpdateAdminRequest) (resp *dto.AdminResponse, err error) {
	var (
		target model.Admin
	)

	err = s.repo.MultiTransaction(
		s.lockTargetForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(&target, dtoReq)

			s.verifySuperAdminRoleUsage(ctx, body.Roles)

			if dtoReq.Roles != nil {
				err := tx.Model(body).Association("Roles").Replace(body.Roles)
				if err != nil {
					return err
				}
			}

			updated, err := s.repo.Tx(tx).UpdateById(ctx, id, body, gormq.WithSelect("name", "username"))
			if err != nil {
				return err
			}

			resp = utils.MustCopy(new(dto.AdminResponse), updated)
			return nil
		})
	return
}

// UpdateAdmin updates a Admin.
func (s *AdminService) UpdatePatchAdmin(ctx context.Context, id uint, p *dto.UpdatePatchAdminRequest) (resp map[string]any, err error) {
	var (
		target model.Admin
	)

	err = s.repo.MultiTransaction(
		s.lockTargetForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
				"name":        gormq.Ignore(),
				"username":    gormq.Ignore(),
				"password":    gormq.ToColumn("password_hash"),
				"ipWhiteList": gormq.ToSnake(),
				"status":      gormq.Ignore(),
			})

			dbCols, res := utils.StructToMap(p, columnMap)
			if p.Password != nil {
				passwordHash, err := s.password_s.HashPassword(*p.Password)
				if err != nil {
					return err
				}
				dbCols["password_hash"] = passwordHash
			}

			err = tx.Model(target).Scopes(gormq.Equal("id", target.ID)).Updates(dbCols).Error
			if err != nil {
				return err
			}
			resp = res
			return nil
		},
	)

	return
}

// SetAdminTOTP sets a Admin's TOTP.
func (s *AdminService) SetAdminTOTP(ctx context.Context, id uint, p *dto.SetTOTPAdminRequest) (resp *dto.SetTOTPAdminResponse, err error) {
	var (
		target model.Admin
	)

	err = s.repo.MultiTransaction(s.lockTargetForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			// validate requester admin otp first before proceed
			requester := admincontext.MustAdminContext(ctx).Admin
			if !s.totp_s.VerifyTOTP(requester.GoogleSecretKey, p.TOTP) {
				return coreerror.ErrInvalidTOTP
			}

			totpKey := s.totp_s.GenerateSecretKey(target.Username)
			_, err = s.repo.UpdateById(ctx, id, &model.Admin{
				GoogleSecretKey: totpKey.Secret(),
			}, gormq.WithSelect("google_secret_key"))

			resp = &dto.SetTOTPAdminResponse{
				TOTPKey: totpKey.Secret(),
				TOTPURL: totpKey.URL(),
			}

			return nil
		},
	)

	return
}

// DeleteAdmin deletes a Admin by ID.
func (s *AdminService) DeleteAdmin(ctx context.Context, id uint) error {
	var (
		target model.Admin
	)

	return s.repo.MultiTransaction(
		s.lockTargetForUpdate(ctx, id, &target,
			func(q *gorm.DB) *gorm.DB {
				return q.Select("id").Preload("Roles")
			},
		),
		func(tx *gorm.DB) error {
			s.verifySuperAdminRoleUsage(ctx, target.Roles)
			return s.repo.Tx(tx).DeleteById(ctx, id)
		},
	)
}

// validate only super admin can be modify by other super admin
func (s *AdminService) verifySuperAdminRoleUsage(ctx context.Context, roles []*model.AdminRole) error {
	adminCtx := admincontext.MustAdminContext(ctx)

	if adminCtx.IsSuperAdmin() {
		return nil
	}

	for _, role := range roles {
		if role.ID == constant.RoleIdSuperAdmin {
			panic(coreerror.ErrUnauthorized)
		}
	}

	return nil
}

// GetAdmins gets all Admins.
func (s *AdminService) GetAdmins(ctx context.Context, query *dto.GetManyQuery) ([]dto.AdminResponse, *pagi.MetaDto, error) {

	adminTableName := "admins"
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.WithPrefix(adminTableName),
		"createdAt": gormq.WithPrefix(adminTableName, gormq.ToSnake()),
		"username":  gormq.WithPrefix(adminTableName),
		"name":      gormq.WithPrefix(adminTableName),
	})

	searchColumns := columnMap.Pick(
		"id",
		"username",
	).Values()

	_resp, respMeta := pagi.PrepareResponse[model.Admin](&query.QueryDto)

	err := s.repo.GetManyAndCount(ctx, &_resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.Where(gormq.WithFilters(query.Filters, columnMap)),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
		func(q *gorm.DB) *gorm.DB {
			adminCtx := admincontext.MustAdminContext(ctx)

			if adminCtx.IsSuperAdmin() {
				return q
			}

			return q.
				Joins("LEFT JOIN admin_admin_roles ON admins.id = admin_admin_roles.admin_id").
				Where("admins.id NOT IN (?)", s.repo.DB().Table("admin_admin_roles").
					Select("admin_id").
					Where("admin_role_id = ?", constant.RoleIdSuperAdmin)).
				Group("admins.id")

		},
		func(q *gorm.DB) *gorm.DB {
			return q.Preload("Roles")
		},
	)

	resp := make([]dto.AdminResponse, 0, len(_resp))
	for i := range _resp {
		resp = append(resp, *utils.MustCopy(new(dto.AdminResponse), &_resp[i]))
		resp[i].EnableTOTP = _resp[i].GoogleSecretKey != ""
	}

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}
