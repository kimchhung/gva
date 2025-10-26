package auth

import (
	"context"
	"errors"
	"time"

	adminerror "backend/app/admin/error"
	"backend/app/admin/module/auth/dto"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	"backend/app/share/service"
	coreerror "backend/core/error"
	"backend/internal/datetime"
	"backend/internal/gormq"

	"gorm.io/gorm"
)

type AuthService struct {
	admin_r    *repository.AdminRepo
	jwt_s      *service.JwtService
	password_s *service.PasswordService
	ip_s       *service.IPService
	totop_s    *service.TOTPService
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewAuthService(
	jwtService *service.JwtService,
	admin_r *repository.AdminRepo,
	password_s *service.PasswordService,
	ip_s *service.IPService,
	totp_s *service.TOTPService,
) *AuthService {
	return &AuthService{
		admin_r:    admin_r,
		jwt_s:      jwtService,
		password_s: password_s,
		ip_s:       ip_s,
		totop_s:    totp_s,
	}
}

// LoginUser authenticates a user and returns a JWT token if successful.
func (s *AuthService) LoginAdmin(ctx context.Context, p *dto.LoginRequest, currentIP string) (*dto.LoginResponse, error) {
	admin, err := s.admin_r.GetOne(ctx, gormq.Equal("username", p.Username))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, coreerror.ErrNotFound
		}

		return nil, err
	}

	if err := s.password_s.VerifyPassword(admin.PasswordHash, p.Password); err != nil {
		panic(coreerror.ErrInvalidCredentials)
	}

	if admin.IpWhiteList != nil {
		if err := s.ip_s.VerifyWhiteListIP(currentIP, admin.IpWhiteList); err != nil {
			panic(adminerror.ErrAdminWhitelistInvalid)
		}
	}

	if !s.totop_s.VerifyTOTP(admin.GoogleSecretKey, p.TOTP) {
		panic(coreerror.ErrInvalidTOTP)
	}

	token, err := s.jwt_s.GenerateToken(
		service.AddClaimPayload("id", admin.ID),
		service.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		panic(coreerror.ErrInvalidCredentials)
	}

	err = s.admin_r.GetRolesByID(admin.ID, admin)
	if err != nil {
		return nil, err
	}

	err = s.UpdateLoginInfo(ctx, admin, currentIP)

	if err != nil {
		return nil, err
	}

	resp := &dto.LoginResponse{
		Token: token,
		Admin: admin,
	}

	return resp, nil
}

func (s *AuthService) UpdateLoginInfo(ctx context.Context, admin *model.Admin, currentIP string) error {

	ipRecord, err := s.ip_s.GetIPRecord(currentIP)
	if err != nil {
		return err
	}

	now := datetime.Must().ToTime()

	admin.LastLoginIP = admin.CurrentLoginIP
	admin.LastRegion = admin.CurrentRegion
	admin.LastLoginAt = admin.CurrentLoginAt
	admin.CurrentLoginIP = currentIP
	admin.CurrentRegion = ipRecord.Country.Names["zh-CN"]
	admin.CurrentLoginAt = &now

	fields := []string{
		"last_login_ip",
		"last_region",
		"last_login_at",
		"current_login_ip",
		"current_region",
		"current_login_at",
	}

	_, err = s.admin_r.UpdateById(ctx, admin.ID, admin,
		gormq.WithSelect(fields...),
	)

	if err != nil {
		return err
	}

	return nil
}
