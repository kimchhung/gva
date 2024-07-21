package dto

import (
	"github.com/gva/app/database/schema/xid"
)

type PermissionResponse struct {
	ID xid.ID `json:"id" rql:"filter,sort"`

	Group string `json:"group,omitempty"`

	Name string `json:"name,omitempty"`

	Key string `json:"key,omitempty"`

	Order int `json:"order,omitempty"`
}
