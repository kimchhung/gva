package types

import (
	"encoding/json"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

var _ interface {
	graphql.Marshaler
	graphql.Unmarshaler
} = (*RouteMeta)(nil)

type RouteMeta struct {
	Hidden     *bool    `json:"hidden,omitempty"`
	AlwaysShow *bool    `json:"alwaysShow,omitempty"`
	Title      *string  `json:"title,omitempty" rql:"filter,sort"`
	Icon       *string  `json:"icon,omitempty"`
	NoCache    *bool    `json:"noCache,omitempty"`
	Breadcrumb *bool    `json:"breadcrumb,omitempty"`
	Affix      *bool    `json:"affix,omitempty"`
	ActiveMenu *string  `json:"activeMenu,omitempty"`
	NoTagsView *bool    `json:"noTagsView,omitempty"`
	CanTo      *bool    `json:"canTo,omitempty"`
	Permission []string `json:"permission,omitempty"`
}

func (u *RouteMeta) UnmarshalGQL(v interface{}) error {
	return json.Unmarshal(v.([]byte), u)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u RouteMeta) MarshalGQL(w io.Writer) {
	jsonData, _ := json.Marshal(u)
	w.Write(jsonData)

}
