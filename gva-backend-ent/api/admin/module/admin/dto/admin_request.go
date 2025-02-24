package dto

import (
	"github.com/gva/internal/ent"
	"github.com/gva/internal/rql"
)

// Requests & responses Data Transfer Object
type AdminRequest struct {
	ent.Admin
}

type AdminPaginateRequest struct {
	rql.Params
	IsCount bool `query:"isCount"`
}
