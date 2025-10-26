package service

import (
	"fmt"
	"time"

	"backend/app/share/repository"
	"backend/env"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type JwtService struct {
	cfg *env.Config
	log *zap.Logger
}

func NewJwtService(
	cfg *env.Config,
	admin_r *repository.AdminRepo,
	log *zap.Logger,
) *JwtService {
	return &JwtService{
		cfg: cfg,
		log: log,
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
	tokenString, err := token.SignedString([]byte(s.cfg.Admin.Auth.JwtSecret))
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
			return []byte(s.cfg.Admin.Auth.JwtSecret), nil
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
