package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kimchhung/gva/extra/app/common/contexts"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/app/middleware/token"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
)

type JwtService struct {
	cfg *config.Config
	db  *database.Database
}

func NewJwtService(cfg *config.Config, db *database.Database) *JwtService {
	return &JwtService{
		cfg: cfg,
		db:  db,
	}
}

func (s *JwtService) ProtectAdmin() fiber.Handler {
	return token.New(
		token.VerifyFunc(func(c *fiber.Ctx, headerValue string) error {
			token := strings.Replace(headerValue, "Bearer ", "", 1)
			if token == "" {
				return app_err.ErrUnauthorized
			}

			admin := new(ent.Admin)
			if _, err := s.ValidateToken(token, s.AdminValidator(admin)); err != nil {
				return app_err.ErrUnauthorized
			}

			adminCtx := contexts.NewAdminContext(
				c.UserContext(),
				contexts.WithAdmin(admin),
			)

			c.SetUserContext(adminCtx)
			return nil
		}),
	)
}

func (s *JwtService) AdminValidator(out *ent.Admin) ClaimValidator {
	return func(claims jwt.MapClaims) error {
		idStr, ok := claims["id"].(string)
		if !ok {
			return app_err.ErrUnauthorized
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return app_err.ErrUnauthorized
		}

		admin, err := s.db.Admin.Query().Where(admin.IDEQ(id)).
			WithRoles(func(rq *ent.RoleQuery) {
				rq.WithPermissions()
			}).First(context.Background())

		if err != nil {
			// Handle error, for example, if the admin is not found
			return app_err.ErrUnauthorized // or return a more specific error
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
	tokenString, err := token.SignedString([]byte(s.cfg.Jwt.Secret))
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
			return []byte(s.cfg.Jwt.Secret), nil
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

func AddTokenPayload(key string, value string) ClaimOption {
	return func(claims jwt.MapClaims) {
		claims[key] = value
	}
}

func AddTokenExpiredAt(deadline time.Time) ClaimOption {
	return func(claims jwt.MapClaims) {
		claims["exp"] = deadline.Unix()
	}
}
