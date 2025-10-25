package model

import (
	"backend/app/share/constant/table"
	"time"
)

type Admin struct {
	BaseModel
	Status       int      `json:"status"`
	PasswordHash string   `json:"-"`
	Name         string   `json:"name"`
	Username     string   `json:"username" validate:"required"`
	IpWhiteList  []string `json:"ipWhiteList" gorm:"serializer:json"`

	GoogleOTP       string `json:"googleOtp"`
	GoogleSecretKey string `json:"googleSecretKey"`

	CurrentLoginIP string `json:"currentLoginIP"`
	LastLoginIP    string `json:"lastLoginIP"`

	CurrentRegion string `json:"currentRegion"`
	LastRegion    string `json:"lastRegion"`

	CurrentLoginAt *time.Time `json:"currentLoginAt"`
	LastLoginAt    *time.Time `json:"lastLoginAt"`

	Roles []*AdminRole `json:"roles" gorm:"many2many:admin_admin_roles;"`

	// not bind with table
	RoleNameIds     []string `json:"roleNameId" gorm:"-"`
	RoleIds         []uint   `json:"roleIds" gorm:"-"`
	PermissionScope []string `json:"permissionScope" gorm:"-"`
	IsSuperAdmin    bool     `json:"isSuperAdmin" gorm:"-"`
}

func (t *Admin) TableName() string {
	return table.Admin
}

func (t *Admin) NewAdminModel() *Admin {
	return &Admin{
		BaseModel: NewBaseModel(),
	}
}
