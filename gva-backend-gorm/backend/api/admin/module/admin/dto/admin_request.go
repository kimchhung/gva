package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type CreateAdminRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Roles    []struct {
		ID uint `json:"id" validate:"required"`
	} `json:"roles" validate:"required,dive"`
}

type GetAdminRequest struct {
	ID uint `param:"id" validate:"required"`
}

type UpdateAdminRequest struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Roles    *[]struct {
		ID uint `json:"id"`
	} `json:"roles"`
}

type UpdatePatchAdminRequest struct {
	Password    *string  `json:"password"`
	IpWhiteList []string `json:"ipWhiteList"`
	Status      *int     `json:"status"`
}

type GetManyQuery struct {
	pagi.QueryDto
}

type SetTOTPAdminRequest struct {
	TOTP string `json:"totp" validate:"required"`
}
