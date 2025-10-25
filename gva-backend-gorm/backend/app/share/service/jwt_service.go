package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	admincontext "backend/app/admin/context"
	apperror "backend/app/share/error"
	"backend/app/share/model"
	"backend/app/share/repository"
	"backend/env"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type JwtService struct {
	cfg     *env.Config
	log     *zap.Logger
	admin_r *repository.AdminRepo
}

func NewJwtService(
	cfg *env.Config,
	admin_r *repository.AdminRepo,
	log *zap.Logger,
) *JwtService {
	return &JwtService{
		cfg:     cfg,
		admin_r: admin_r,
		log:     log,
	}
}

func (s *JwtService) RequiredAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.TrimSpace(strings.Replace(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ", "", 1))
			if token == "" {
				return apperror.ErrUnauthorized
			}

			ctx := c.Request().Context()
			admin := new(model.Admin)
			if _, err := s.ValidateToken(token, s.AdminValidator(ctx, admin)); err != nil {
				s.log.Debug("s.ValidateToken", zap.Error(err))
				return apperror.ErrUnauthorized
			}

			adminctx := admincontext.NewAdminContext(admincontext.WithAdmin(admin))
			admincontext.SetAdminContext(ctx, adminctx)
			return next(c)
		}
	}
}

func (s *JwtService) AdminValidator(ctx context.Context, out *model.Admin) ClaimValidator {
	return func(claims jwt.MapClaims) error {
		id, ok := claims["id"].(float64)
		if !ok || id == 0 {
			return apperror.ErrUnauthorized
		}

		admin, err := s.admin_r.GetById(ctx, uint(id))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.ErrNotFound
			}

			return err
		}

		if admin.Status == 0 {
			return apperror.ErrValidationError
		}

		err = s.admin_r.GetRolesByID(admin.ID, admin)
		if err != nil {
			return err
		}

		*out = *admin
		return nil
	}
}

type ClaimOption func(claims jwt.MapClaims)

// GenerateToken generates a new JWT token for the given user ID.
func (s *JwtService) GenerateToken(opt ClaimOption, opts ...ClaimOption) (string, error) {
	// Define the token claims
	claims := jwt.MapClaims{}
	for _, op := range append(opts, opt) {
		op(claims)
	}

	// Create the token using the claims and the.cfg.Jwt.Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.API.Admin.Auth.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type ClaimValidator func(claims jwt.MapClaims) error

// ValidateToken validates the provided JWT token and returns the user ID if valid.
func (s *JwtService) ValidateToken(tokenString string, opts ...ClaimValidator) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.cfg.API.Admin.Auth.JwtSecret), nil
		},
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for _, opt := range opts {
			if err := opt(claims); err != nil {
				return nil, err
			}
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func AddClaimPayload(key string, value any) ClaimOption {
	return func(claims jwt.MapClaims) {
		claims[key] = value
	}
}

func AddTokenExpiredAt(deadline time.Time) ClaimOption {
	return func(claims jwt.MapClaims) {
		claims["exp"] = deadline.Unix()
	}
}
