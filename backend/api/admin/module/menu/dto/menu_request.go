package dto

import (
	"github.com/gva/app/database/schema/types"
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

	// optionals
	ParentID *int `json:"parentId,omitempty"  validate:"min=0,omitempty"`
}

type MenuPaginateRequest struct {
	rql.Params
	IsCount       bool `query:"isCount"`
	IsGroupNested bool `query:"isGroupNested"`
}
