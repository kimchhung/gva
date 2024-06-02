package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/kimchhung/gva/extra/api/admin/module/auth/dto"
	apperror "github.com/kimchhung/gva/extra/app/common/error"
	repository "github.com/kimchhung/gva/extra/app/common/repository"
	"github.com/kimchhung/gva/extra/app/common/service"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
	"github.com/kimchhung/gva/extra/lang"
)

type AuthService struct {
	admin_r    *repository.AdminRepository
	jwt_s      *service.JwtService
	password_s *service.PasswordService
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewAuthService(jwtService *service.JwtService, admin_r *repository.AdminRepository, passwordService *service.PasswordService) *AuthService {
	return &AuthService{
		admin_r:    admin_r,
		jwt_s:      jwtService,
		password_s: passwordService,
	}
}

func (s *AuthService) RegisterAdmin(ctx context.Context, dto *dto.RegisterRequest) (string, *ent.Admin, error) {
	hashedPassword, err := s.password_s.HashPassword(dto.Password)
	if err != nil {
		panic(err)
	}

	admin, err := s.admin_r.C().Create().
		SetUsername(dto.Username).
		SetPassword(hashedPassword).
		SetDisplayName(dto.DisplayName).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			panic(apperror.ErrUsernameExists)
		}

		panic(err)
	}

	// Generate a JWT token for the authenticated user
	token, err := s.jwt_s.GenerateToken(
		service.AddTokenPayload("id", fmt.Sprintf("%d", admin.ID)),
		service.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		panic(err)
	}

	return token, admin, nil
}

// LoginUser authenticates a user and returns a JWT token if successful.
func (s *AuthService) LoginAdmin(ctx context.Context, dto *dto.LoginRequest) (string, *ent.Admin, error) {
	admin, err := s.admin_r.C().Query().Where(admin.Username(dto.Username)).WithRoles().First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			panic(
				apperror.NewError(
					apperror.ErrNotFound,
					apperror.Prefix(lang.ForContext(ctx), "Admin"),
				),
			)
		}

		panic(err)
	}

	// Verify the password (assuming you have a method to do this)
	if err := s.password_s.VerifyPassword(admin.Password, dto.Password); err != nil {
		panic(apperror.ErrPasswordValidationError)
	}

	// Generate a JWT token for the authenticated user
	token, err := s.jwt_s.GenerateToken(
		service.AddTokenPayload("id", fmt.Sprintf("%d", admin.ID)),
		service.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		panic(apperror.ErrPasswordValidationError)
	}

	return token, admin, nil
}
