package dto

import (
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/rql"
)

// Requests Data Transfer Object
type DepartmentRequest struct {
	Name     string   `json:"name" validate:"required"`
	NameId   string   `json:"nameId" validate:"required"`
	Pid      *pxid.ID `json:"pid"`
	IsEnable *bool    `json:"isEnable"`
}

type DepartmentPagedRequest struct {
	rql.Params
	Selects string `query:"selects"`
}
