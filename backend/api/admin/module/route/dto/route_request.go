package dto

import (
	"github.com/kimchhung/gva/backend-echo/app/database/schema/types"
	"github.com/kimchhung/gva/backend-echo/internal/ent/route"
	"github.com/kimchhung/gva/backend-echo/internal/rql"
)

// Requests & responses Data Transfer Object
type RouteRequest struct {
	Name      string          `json:"name,omitempty" validate:"required"`
	Path      string          `json:"path,omitempty" validate:"required"`
	Component string          `json:"component,omitempty" validate:"required"`
	IsEnable  bool            `json:"isEnable"  validate:"required"`
	Type      route.Type      `json:"type,omitempty" validate:"required"`
	Meta      types.RouteMeta `json:"meta,omitempty" validate:"required"`

	// optionals
	ParentID *int `json:"parentId,omitempty"  validate:"min=0,omitempty"`
}

type RoutePaginateRequest struct {
	rql.Params
	IsCount       bool `query:"isCount"`
	IsGroupNested bool `query:"isGroupNested"`
}
