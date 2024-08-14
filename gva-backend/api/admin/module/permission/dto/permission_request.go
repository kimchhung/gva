package dto

import "github.com/gva/internal/rql"

// Requests Data Transfer Object
type PermissionRequest struct {
	Group string `json:"group" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Scope string `json:"scope" validate:"required"`
	Order int    `json:"order" validate:"required"`
}

type PermissionPagedRequest struct {
	rql.Params
	Selects string `query:"selects"`
}
