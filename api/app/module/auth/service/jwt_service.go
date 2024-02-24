package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	app_err "github.com/kimchhung/gva/app/error"
	"github.com/kimchhung/gva/app/middleware/token"
	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/bootstrap/database"
	"github.com/kimchhung/gva/internal/ent"
	"github.com/kimchhung/gva/internal/ent/admin"
)

type JwtService struct {
	secret string
	db     *database.Database
}

func NewJwtService(cfg *config.Config, db *database.Database) (*JwtService, error) {
	secret := cfg.Jwt.Secret
	if secret == "" {
		return nil, fmt.Errorf("cfg.Jwt.Secret is required %v", cfg.Jwt)
	}

	return &JwtService{
		secret: secret,
		db:     db,
	}, nil
}

func (s *JwtService) Protect(opts ...ClaimValidator) fiber.Handler {
	return token.New(token.NewConfig(
		&token.Config{
			Next:       nil,
			HeaderName: "authorization",
			VerifyFunc: func(c *fiber.Ctx, headerValue string) error {
				token := strings.Replace(headerValue, "Bearer ", "", 1)

				if len(opts) == 0 {
					defautValidator := s.AdminValidator(c.UserContext(), new(ent.Admin))
					opts = append(opts, defautValidator)
				}

				_, err := s.ValidateToken(token, opts...)
				return err
			},
		},
	))
}

func (s *JwtService) AdminValidator(ctx context.Context, out *ent.Admin) ClaimValidator {
	return func(claims jwt.MapClaims) error {
		idStr, ok := claims["id"].(string)
		if !ok {
			return app_err.ErrUnauthorized
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return app_err.ErrUnauthorized
		}

		admin, err := s.db.Ent.Admin.Query().Where(admin.IDEQ(id)).First(ctx)
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

	// Create the token using the claims and the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secret))
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
			return []byte(s.secret), nil
		},
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for _, opt := range opts {
			opt(claims)
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func AddPayload(key string, value string) ClaimOption {
	return func(claims jwt.MapClaims) {
		claims[key] = value
	}
}

func AddExpiredAt(deadline time.Time) ClaimOption {
	return func(claims jwt.MapClaims) {
		claims["exp"] = deadline.Unix()
	}
}
