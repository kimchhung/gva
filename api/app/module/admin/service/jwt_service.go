package service

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	app_err "github.com/kimchhung/gva/app/error"
	"github.com/kimchhung/gva/app/middleware/token"
	"github.com/kimchhung/gva/config"
)

type JWTService struct {
	secret string
}

func NewJWTService(cfg *config.Config) *JWTService {
	return &JWTService{
		secret: cfg.Jwt.Secret,
	}
}

func (s *JWTService) Protected() fiber.Handler {
	config := token.NewConfig(&token.Config{
		HeaderName: "authorization",
		Secret:     s.secret,
		VerifyFunc: func(headerValue string) error {
			token := strings.Replace(headerValue, "Bearer ", "", 1)
			return s.VerifyToken(token)
		},
	})

	return token.New(config)
}

func (s *JWTService) GenerateToken(username string) (string, error) {
	// Define the token claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString(s.secret)
	return tokenString, err
}

func (s *JWTService) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return app_err.ErrUnauthorized
	}

	return nil
}
