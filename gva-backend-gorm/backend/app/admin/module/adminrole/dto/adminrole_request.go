package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type CreateAdminRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions" validate:"required"`
}

type GetAdminRoleRequest struct {
	ID uint `param:"id" validate:"required"`
}

type UpdateAdminRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions" validate:"required"`
}

// not nil will update
type UpdatePatchAdminRoleRequest struct {
	Status *int `json:"status"`
}

type GetManyQuery struct {
	pagi.QueryDto
}
