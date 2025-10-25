package dto

import (
	"backend/app/share/constant/table"
	"backend/app/share/model"
)

// Requests & responses Data Transfer Object
type AdminRoleResponse struct {
	model.BaseModel
	Name        string              `json:"name"`
	NameID      string              `json:"nameId"`
	Description string              `json:"description"`
	Status      int                 `json:"status"`
	Permissions []*model.Permission `json:"permissions,omitempty" gorm:"many2many:admin_role_permissions;"`
}

func (AdminRoleResponse) TableName() string {
	return table.AdminRole
}
