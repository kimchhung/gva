package dto

import (
	"github.com/gva/app/database/schema/types"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent/menu"
	"github.com/gva/internal/rql"
)

// Requests & responses Data Transfer Object
type MenuRequest struct {
	Name      string         `json:"name,omitempty" validate:"required"`
	Path      string         `json:"path,omitempty" validate:"required"`
	Component string         `json:"component,omitempty" validate:"required"`
	IsEnable  bool           `json:"isEnable"  validate:"required"`
	Type      menu.Type      `json:"type,omitempty" validate:"required"`
	Meta      types.MenuMeta `json:"meta,omitempty" validate:"required"`
	Order     int            `json:"order,omitempty" validate:"number"`

	// optionals
	Redirect *string `json:"redirect,omitempty"`
	Pid      *xid.ID `json:"pid,omitempty"`
}

type MenuPagedRequest struct {
	rql.Params
	Selects string `query:"selects"`
}
