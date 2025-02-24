package dto

import "backend/app/common/constant/table"

// Requests & responses Data Transfer Object
type PermissionResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Scope string `json:"scope"`
	Group string `json:"group"`
	Order int    `json:"order"`
}

func (PermissionResponse) TableName() string {
	return table.Permission
}
