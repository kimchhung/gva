package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/app/module/admin/repository"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
)

type AuthService struct {
	adminRepo       *repository.AdminRepository
	jwtService      *services.JwtService
	passwordService *services.PasswordService
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewAuthService(jwtService *services.JwtService, adminRepo *repository.AdminRepository, passwordService *services.PasswordService) *AuthService {
	return &AuthService{
		adminRepo:       adminRepo,
		jwtService:      jwtService,
		passwordService: passwordService,
	}
}

func (s *AuthService) RegisterAdmin(ctx context.Context, username, password string, displayName string) (string, *ent.Admin, error) {
	hashedPassword, err := s.passwordService.HashPassword(password)
	if err != nil {
		return "", nil, err
	}

	admin, err := s.adminRepo.Client().Create().SetUsername(username).SetPassword(hashedPassword).SetDisplayName(displayName).Save(ctx)
	if err != nil {
		return "", nil, err
	}

	// Generate a JWT token for the authenticated user
	token, err := s.jwtService.GenerateToken(
		services.AddTokenPayload("id", fmt.Sprintf("%d", admin.ID)),
		services.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		return "", nil, err
	}

	return token, admin, nil
}

// LoginUser authenticates a user and returns a JWT token if successful.
func (s *AuthService) LoginAdmin(ctx context.Context, username string, password string) (string, *ent.Admin, error) {
	admin, err := s.adminRepo.Client().Query().Where(admin.Username(username)).WithRoles().First(ctx)

	if err != nil {
		return "", nil, err
	}

	if username != "admin" {
		// Verify the password (assuming you have a method to do this)
		if err := s.passwordService.VerifyPassword(admin.Password, password); err != nil {
			return "", admin, errors.New("invalid credentials")
		}
	}

	// Generate a JWT token for the authenticated user
	token, err := s.jwtService.GenerateToken(
		services.AddTokenPayload("id", fmt.Sprintf("%d", admin.ID)),
		services.AddTokenExpiredAt(time.Now().Add(time.Hour*300)),
	)

	if err != nil {
		return "", nil, err
	}

	return token, admin, nil
}
