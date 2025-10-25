package model

import "backend/app/share/constant/table"

type AdminRoleType int

const (
	AdminRoleTypeChangeable    AdminRoleType = 1
	AdminRoleTypeNotChangeable AdminRoleType = 0
)

type AdminRole struct {
	BaseModel
	Description string        `json:"description,omitempty"`
	Name        string        `json:"name,omitempty"`
	NameID      string        `json:"nameId,omitempty"`
	Order       int           `json:"order,omitempty"`
	Status      int           `json:"status,omitempty"`
	Type        AdminRoleType `json:"type,omitempty"`
	Permissions []Permission  `json:"permissions,omitempty" gorm:"many2many:admin_role_permissions;"`
}

func (t *AdminRole) TableName() string {
	return table.AdminRole
}
