package dto

import (
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rql"
)

// Requests & responses Data Transfer Object
type AdminRequest struct {
	ent.Admin
}

type AdminPaginateRequest struct {
	rql.Params
	IsCount bool `query:"isCount"`
}
