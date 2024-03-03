package dto

import "github.com/kimchhung/gva/extra/internal/ent"

// Requests & responses Data Transfer Object
type PermissionResponse struct {
	*ent.Permission
}
