package dto

import "github.com/gva/internal/ent"

func ToMenuResponse(value ...*ent.Menu) []*MenuResponse {
	list := make([]*MenuResponse, len(value))
	for i, v := range value {
		list[i] = &MenuResponse{Menu: v}
	}

	return list
}
