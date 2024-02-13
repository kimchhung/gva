package service

import (
	"gva/app/module/admin/dto"
	"gva/app/module/admin/repository"

	"context"
	"gva/internal/ent"
	"gva/internal/ent/admin"
)

type AdminService struct {
	repo *repository.AdminRepository
}

func NewAdminService(repository *repository.AdminRepository) *AdminService {
	return &AdminService{
		repo: repository,
	}
}

func (s *AdminService) GetAdmins(ctx context.Context) ([]*ent.Admin, error) {
	return s.repo.Client().Query().Order(ent.Asc(admin.FieldID)).All(ctx)
}

func (s *AdminService) GetAdminByID(ctx context.Context, id int) (*ent.Admin, error) {
	return s.repo.Client().Query().Where(admin.IDEQ(id)).First(ctx)
}

func (s *AdminService) CreateAdmin(ctx context.Context, request dto.AdminRequest) (*ent.Admin, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *AdminService) UpdateAdmin(ctx context.Context, id int, request dto.AdminRequest) (*ent.Admin, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *AdminService) DeleteAdmin(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}
