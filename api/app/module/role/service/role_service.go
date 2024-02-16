package service

import (
	"github.com/kimchhung/gva/app/module/role/dto"
	"github.com/kimchhung/gva/app/module/role/repository"

	"context"

	"github.com/kimchhung/gva/internal/ent"
	"github.com/kimchhung/gva/internal/ent/role"
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repository *repository.RoleRepository) *RoleService {
	return &RoleService{
		repo: repository,
	}
}

func (s *RoleService) GetRoles(ctx context.Context) ([]*ent.Role, error) {
	return s.repo.Client().Query().Order(ent.Asc(role.FieldID)).All(ctx)
}

func (s *RoleService) GetRoleByID(ctx context.Context, id int) (*ent.Role, error) {
	return s.repo.Client().Query().Where(role.IDEQ(id)).First(ctx)
}

func (s *RoleService) CreateRole(ctx context.Context, request dto.RoleRequest) (*ent.Role, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *RoleService) UpdateRole(ctx context.Context, id int, request dto.RoleRequest) (*ent.Role, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *RoleService) DeleteRole(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}
