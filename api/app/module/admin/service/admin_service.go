package service

import (
	"gva/app/module/admin/repos"
	"gva/app/module/admin/request"

	"context"
	"gva/internal/ent"
	"gva/internal/ent/admin"
)

type AdminService struct {
	repo *repos.AdminRepository
}

func NewAdminService(repo *repos.AdminRepository) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (s *AdminService) GetAdmins(ctx context.Context) ([]*ent.Admin, error) {
	return s.repo.Client().Query().Order(ent.Asc(admin.FieldID)).All(ctx)
}

func (s *AdminService) GetAdminByID(ctx context.Context, id int) (*ent.Admin, error) {
	return s.repo.Client().Query().Where(admin.IDEQ(id)).First(ctx)
}

func (s *AdminService) CreateAdmin(ctx context.Context, request request.AdminRequest) (*ent.Admin, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *AdminService) UpdateAdmin(ctx context.Context, id int, request request.AdminRequest) (*ent.Admin, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *AdminService) DeleteAdmin(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}
