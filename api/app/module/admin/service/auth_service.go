package service

import (
	"github.com/kimchhung/gva/app/module/admin/repository"
)

type AuthService struct {
	repo *repository.AdminRepository
}

func NewAuthService(repository *repository.AdminRepository) *AuthService {
	return &AuthService{
		repo: repository,
	}
}

func (s *AuthService) Login() {

}
