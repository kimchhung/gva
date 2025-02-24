package relay

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type CursorDto struct {
	First  *int    `json:"first,omitempty" query:"first"`
	Last   *int    `json:"last,omitempty" query:"last"`
	After  *string `json:"after,omitempty" query:"after"`
	Before *string `json:"before,omitempty" query:"before"`

	// sort="createdAt,-amount,+title",
	// eq order by created_at asc, ammount asc, title asc
	Sorts []string `json:"sorts,omitempty"`

	// whether return count record or not, default false
	IsCount bool `json:"isCount,omitempty" query:"isCount"`

	// ?[id][eq]=1&[id][createdAt][eq]=2024-05-17T09:49:15.466Z
	Filters map[string]map[string]any
}

type CursorOption func(*CursorDto)

func ParseFilter(c echo.Context, opts ...CursorOption) (*CursorDto, error) {
	payload := &CursorDto{}

	if err := c.Bind(payload); err != nil {
		return payload, err
	}

	for _, opt := range opts {
		opt(payload)
	}

	return payload, nil
}

func (q *CursorDto) GetOrders(mapper map[string]string) map[string]string {
	if len(q.Sorts) == 0 && len(mapper) == 0 {
		return nil
	}

	maps := make(map[string]string, len(q.Sorts))

	for i, sort := range q.Sorts {
		sort = strings.TrimSpace(sort)
		if sort == "" {
			continue
		}

		var (
			column    = sort
			direction = "asc"
		)

		switch []rune(sort)[0] {
		case '-':
			column = strings.TrimPrefix(sort, "-")
			direction = "desc"

		case '+':
			column = strings.TrimPrefix(sort, "+")
			direction = "asc"
		}

		if v, hasMap := mapper[column]; hasMap {
			// allow only whitelist
			maps[fmt.Sprintf("%d-%s", i, v)] = direction
		}
	}

	return maps
}

func (q *CursorDto) GetFilters(mapper map[string]string) any {
	validFilter := make(map[string]any)

	for key, v := range q.Filters {
		validKey, ok := mapper[key]
		if ok {
			validFilter[validKey] = v
		}
	}

	return validFilter
}
