package service

import (
	"backend/core/env"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TOTPService struct {
	cfg *env.Config
}

func NewTOTPService(cfg *env.Config) *TOTPService {
	return &TOTPService{
		cfg: cfg,
	}
}

func (s *TOTPService) GenerateSecretKey(accountName string) *otp.Key {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.cfg.App.Name,
		AccountName: accountName,
	})

	if err != nil {
		panic(err)
	}
	return key
}

// VerifyPassword checks if a given password matches the hashed password.
func (s *TOTPService) VerifyTOTP(secretKey, code string) bool {
	testCode := s.cfg.Admin.Auth.TotpTestCode
	if testCode != "" && testCode == code {
		return true
	}

	if secretKey == "" {
		return false
	}

	return totp.Validate(code, secretKey)
}
