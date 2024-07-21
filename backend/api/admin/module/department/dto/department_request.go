package dto

import (
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/rql"
)

// Requests Data Transfer Object
type DepartmentRequest struct {
	Name     string  `json:"name" validate:"required"`
	NameId   string  `json:"nameId" validate:"required"`
	ParentId *xid.ID `json:"parentId"`
	IsEnable *bool   `json:"isEnable"`
}

type DepartmentPagedRequest struct {
	rql.Params
	Selects string `query:"selects"`
}
