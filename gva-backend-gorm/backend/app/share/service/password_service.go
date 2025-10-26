package service

import (
	"backend/core/env"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	cfg *env.Config
}

func NewPasswordService(cfg *env.Config) *PasswordService {
	return &PasswordService{
		cfg: cfg,
	}
}

// HashPassword hashes a password using bcrypt.
func (s *PasswordService) HashPassword(password string) (string, error) {
	cost := s.cfg.Admin.Auth.PasswordHashCost
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if a given password matches the hashed password.
func (s *PasswordService) VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
