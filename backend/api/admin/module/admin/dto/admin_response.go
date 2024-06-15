package dto

import "github.com/kimchhung/gva/backend/internal/ent"

// Requests & responses Data Transfer Object
type AdminResponse struct {
	*ent.Admin
}
