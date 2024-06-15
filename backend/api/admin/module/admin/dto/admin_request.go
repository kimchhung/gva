package dto

import (
	"github.com/kimchhung/gva/backend/internal/ent"
	"github.com/kimchhung/gva/backend/internal/rql"
)

// Requests & responses Data Transfer Object
type AdminRequest struct {
	ent.Admin
}

type AdminPaginateRequest struct {
	rql.Params
	IsCount bool `query:"isCount"`
}
