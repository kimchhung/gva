package repository

import (
	"backend/app/share/constant"
	"backend/app/share/model"
	"backend/core/database"

	"go.uber.org/fx"
)

// UserQueries struct for queries from User model.
type AdminRepo struct {
	IBaseRepository[model.Admin]
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewAdminRepo))
}

func NewAdminRepo(db *database.Database) *AdminRepo {
	return &AdminRepo{
		NewBaseRepository[model.Admin](db.DB),
	}
}

func (r *AdminRepo) GetRolesByID(adminID uint, admin *model.Admin) error {
	err := r.DB().Model(&model.AdminRole{}).
		Select("admin_roles.id", "admin_roles.name_id").
		Joins("inner join admin_admin_roles on admin_roles.id = admin_admin_roles.admin_role_id").
		Where("admin_admin_roles.admin_id = ?", adminID).
		Where("admin_roles.status = 1").
		Find(&admin.Roles).Error

	if err != nil {
		return err
	}

	adminRoleIds := make([]uint, 0)
	for _, role := range admin.Roles {
		adminRoleIds = append(adminRoleIds, role.ID)
		admin.RoleNameIds = append(admin.RoleNameIds, role.NameID)
		if role.NameID == constant.RoleNameIDSuperAdmin {
			admin.IsSuperAdmin = true
		}
	}
	admin.RoleIds = adminRoleIds

	err = r.DB().Model(&model.Permission{}).
		Distinct("scope").
		Joins("inner join admin_role_permissions on permissions.id = admin_role_permissions.permission_id").
		Where("admin_role_permissions.admin_role_id IN (?)", adminRoleIds).
		Pluck("scope", &admin.PermissionScope).Error

	if err != nil {
		return err
	}

	return nil
}
