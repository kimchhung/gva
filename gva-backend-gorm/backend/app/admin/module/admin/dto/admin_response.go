package dto

import (
	"backend/app/share/model"
	"time"
)

// Requests & responses Data Transfer Object
type AdminResponse struct {
	model.BaseModel

	Name           string             `json:"name"`
	Username       string             `json:"username"`
	IpWhiteList    []string           `json:"ipWhiteList"`
	CurrentLoginIp string             `json:"currentLoginIp"`
	CurrentRegion  string             `json:"currentRegion"`
	CurrentLoginAt *time.Time         `json:"currentLoginAt"`
	Roles          []*model.AdminRole `json:"roles" gorm:"many2many:admin_admin_roles;"`
	Status         int                `json:"status"`
	EnableTOTP     bool               `json:"enableTOTP"`
}

type SetTOTPAdminResponse struct {
	TOTPKey string `json:"totpKey"`
	TOTPURL string `json:"totpURL"`
}
