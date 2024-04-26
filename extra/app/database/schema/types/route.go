package types

type RouteMeta struct {
	Hidden     *bool    `json:"hidden,omitempty"`
	AlwaysShow *bool    `json:"alwaysShow,omitempty"`
	Title      *string  `json:"title,omitempty"`
	Icon       *string  `json:"icon,omitempty"`
	NoCache    *bool    `json:"noCache,omitempty"`
	Breadcrumb *bool    `json:"breadcrumb,omitempty"`
	Affix      *bool    `json:"affix,omitempty"`
	ActiveMenu *string  `json:"activeMenu,omitempty"`
	NoTagsView *bool    `json:"noTagsView,omitempty"`
	CanTo      *bool    `json:"canTo,omitempty"`
	Permission []string `json:"permission,omitempty"`
}
