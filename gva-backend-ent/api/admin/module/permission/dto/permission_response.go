package dto

import (
	"github.com/gva/app/database/schema/pxid"
)

// Responses Data Transfer Object
type PermissionResponse struct {
	ID    pxid.ID `json:"id"`
	Group string  `json:"group"`
	Name  string  `json:"name"`
	Scope string  `json:"scope"`
	Order int     `json:"order"`
}
