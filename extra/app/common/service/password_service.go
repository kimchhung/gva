package service

import (
	"github.com/kimchhung/gva/extra/config"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	cfg *config.Config
}

func NewPasswordService(cfg *config.Config) *PasswordService {
	return &PasswordService{
		cfg: cfg,
	}
}

// HashPassword hashes a password using bcrypt.
func (s *PasswordService) HashPassword(password string) (string, error) {
	cost := s.cfg.Password.HashCost
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
