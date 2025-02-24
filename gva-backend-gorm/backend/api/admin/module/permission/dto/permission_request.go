package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type CreatePermissionRequest struct {
	Title string `json:"title" validate:"required"`
}

type GetPermissionRequest struct {
	ID uint `json:"id" validate:"required"`
}

type UpdatePermissionRequest struct {
	Title string `json:"title"`
}


// not nil will update
type UpdatePatchPermissionRequest struct {
	Title *string `json:"title"`
}

type GetManyQuery struct {
	pagi.QueryDto
}
