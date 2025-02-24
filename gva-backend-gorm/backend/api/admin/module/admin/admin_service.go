package admin

import (
	"context"
	"errors"

	"backend/api/admin/module/admin/dto"
	appctx "backend/app/common/context"
	apperror "backend/app/common/error"
	"backend/app/common/model"
	repository "backend/app/common/repository"
	"backend/app/common/service"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/utils"

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
			return nil, apperror.ErrUsernameExists
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

// UpdateAdmin updates a Admin.
func (s *AdminService) UpdateAdmin(ctx context.Context, id uint, p *dto.UpdateAdminRequest) (resp *dto.AdminResponse, err error) {
	err = s.repo.Transaction(
		func(tx *gorm.DB) error {
			body := utils.MustCopy(new(model.Admin), p)
			body.ID = id

			s.verifySuperAdminRoleUsage(ctx, body.Roles)

			if p.Roles != nil {
				err := tx.Model(body).Association("Roles").Replace(body.Roles)
				if err != nil {
					return err
				}
			}

			updated, err := s.repo.Tx(tx).UpdateById(ctx, id, body)
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
	err = s.repo.Transaction(
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

			err = tx.Model(&model.Admin{}).Scopes(gormq.Equal("id", id)).Updates(dbCols).Error
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
func (s *AdminService) SetAdminTOTP(ctx context.Context, id uint, p *dto.SetTOTPAdminRequest) (*dto.SetTOTPAdminResponse, error) {
	// validate requester admin otp first before proceed
	requester := appctx.MustAdminContext(ctx).Admin
	if !s.totp_s.VerifyTOTP(requester.GoogleSecretKey, p.TOTP) {
		return nil, apperror.ErrInvalidTOTP
	}

	// generate new secret key for the target targetAdmin
	targetAdmin, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	totpKey := s.totp_s.GenerateSecretKey(targetAdmin.Username)
	_, err = s.repo.UpdateById(ctx, id, &model.Admin{
		GoogleSecretKey: totpKey.Secret(),
	}, gormq.WithSelect("google_secret_key"))

	return &dto.SetTOTPAdminResponse{
		TOTPKey: totpKey.Secret(),
		TOTPURL: totpKey.URL(),
	}, err
}

// DeleteAdmin deletes a Admin by ID.
func (s *AdminService) DeleteAdmin(ctx context.Context, id uint) error {
	admin, err := s.repo.GetById(ctx, id, func(q *gorm.DB) *gorm.DB {
		return q.Select("id").Preload("Roles")
	})

	if err != nil {
		return err
	}

	s.verifySuperAdminRoleUsage(ctx, admin.Roles)

	return s.repo.DeleteById(ctx, id)
}

// validate only super admin can be modify by other super admin
func (s *AdminService) verifySuperAdminRoleUsage(ctx context.Context, roles []*model.AdminRole) error {
	adminCtx := appctx.MustAdminContext(ctx)

	if adminCtx.IsSuperAdmin() {
		return nil
	}

	for _, role := range roles {
		if role.ID == appctx.RoleIdSuperAdmin {
			panic(apperror.ErrUnauthorized)
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
			adminCtx := appctx.MustAdminContext(ctx)

			if adminCtx.IsSuperAdmin() {
				return q
			}

			return q.
				Joins("LEFT JOIN admin_admin_roles ON admins.id = admin_admin_roles.admin_id").
				Where("admins.id NOT IN (?)", s.repo.Table("admin_admin_roles").
					Select("admin_id").
					Where("admin_role_id = ?", appctx.RoleIdSuperAdmin)).
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
