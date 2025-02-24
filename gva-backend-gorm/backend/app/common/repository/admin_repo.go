package repository

import (
	appctx "backend/app/common/context"
	"backend/app/common/model"
	"backend/internal/bootstrap/database"

	"go.uber.org/fx"
)

// UserQueries struct for queries from User model.
type AdminRepo struct {
	IBaseRepository[model.Admin]
	*database.Database
}

func init() {
	dependencies = append(dependencies, fx.Provide(NewAdminRepo))
}

func NewAdminRepo(db *database.Database) *AdminRepo {
	return &AdminRepo{
		NewBaseRepository[model.Admin](db.DB),
		db,
	}
}

func (r *AdminRepo) GetRolesByID(adminID uint, admin *model.Admin) error {
	err := r.DB.Model(&model.AdminRole{}).
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
		if role.NameID == appctx.RoleNameIDSuperAdmin {
			admin.IsSuperAdmin = true
		}
	}
	admin.RoleIds = adminRoleIds

	err = r.DB.Model(&model.Permission{}).
		Distinct("scope").
		Joins("inner join admin_role_permissions on permissions.id = admin_role_permissions.permission_id").
		Where("admin_role_permissions.admin_role_id IN (?)", adminRoleIds).
		Pluck("scope", &admin.PermissionScope).Error

	if err != nil {
		return err
	}

	return nil
}
