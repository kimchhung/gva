package auth

import (
	"context"
	"time"

	"github.com/gva/api/admin/module/auth/dto"
	apperror "github.com/gva/app/common/error"
	repository "github.com/gva/app/common/repository"
	"github.com/gva/app/common/service"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/lang"
)

type AuthService struct {
	admin_r    *repository.AdminRepository
	jwt_s      *service.JwtService
	password_s *service.PasswordService
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewAuthService(jwtService *service.JwtService, admin_r *repository.AdminRepository, password_s *service.PasswordService) *AuthService {
	return &AuthService{
		admin_r:    admin_r,
		jwt_s:      jwtService,
		password_s: password_s,
	}
}

func (s *AuthService) RegisterAdmin(ctx context.Context, p *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := s.password_s.HashPassword(p.Password)
	if err != nil {
		return nil, err
	}

	admin, err := s.admin_r.C().Create().
		SetUsername(p.Username).
		SetPassword(hashedPassword).
		SetDisplayName(p.DisplayName).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			panic(apperror.ErrUsernameExists)
		}

		return nil, err
	}

	// Generate a JWT token for the authenticated user
	token, err := s.jwt_s.GenerateToken(
		service.AddClaimPayload("id", admin.ID),
		service.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		return nil, err
	}

	resp := &dto.RegisterResponse{
		Token: token,
		Admin: admin,
	}

	return resp, nil
}

// LoginUser authenticates a user and returns a JWT token if successful.
func (s *AuthService) LoginAdmin(ctx context.Context, p *dto.LoginRequest) (*dto.LoginResponse, error) {
	admin, err := s.admin_r.C().Query().Where(admin.Username(p.Username)).WithRoles().First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, apperror.NewError(
				apperror.ErrNotFound,
				apperror.Prefix(lang.ForContext(ctx), "Admin"),
			)
		}

		return nil, err
	}

	// Verify the password (assuming you have a method to do this)
	if err := s.password_s.VerifyPassword(admin.Password, p.Password); err != nil {
		panic(apperror.ErrPasswordValidationError)
	}

	// Generate a JWT token for the authenticated user
	token, err := s.jwt_s.GenerateToken(
		service.AddClaimPayload("id", admin.ID),
		service.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		panic(apperror.ErrPasswordValidationError)
	}

	resp := &dto.LoginResponse{
		Token: token,
		Admin: admin,
	}

	return resp, nil
}
