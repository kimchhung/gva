package dto

import "github.com/kimchhung/gva/backend-echo/internal/ent"

type PermissionResponse struct {
	*ent.Permission
}
