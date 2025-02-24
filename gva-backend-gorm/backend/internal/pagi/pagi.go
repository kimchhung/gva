package pagi

import (
	"backend/internal/gormq"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

type FilterMap map[string]map[gormq.OpName]any

type QueryDto struct {
	Page    int       `json:"page" query:"page"`
	Limit   int       `json:"limit" query:"limit"`
	Filters FilterMap `json:"filters" query:"filters"`
	Sorts   []string  `json:"sorts" query:"sorts"`
	Search  string    `json:"search" query:"search"`

	// [totalCount, list]
	Selects []string `json:"selects" query:"selects"`

	// don't add,
	// extends if want to add field

}

type MetaDto struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalCount *int64 `json:"totalCount,omitempty"`
	HasNext    bool   `json:"hasNext"`
}

func PrepareResponse[T any](dto *QueryDto) (resp []T, meta *MetaDto) {
	respMeta := &MetaDto{
		Page:  dto.Page,
		Limit: dto.Limit,
	}

	if slices.Contains(dto.Selects, "totalCount") {
		respMeta.TotalCount = new(int64)
		return nil, respMeta
	}

	return resp, respMeta
}

func (m *MetaDto) UpdatePagination(dataLength int) *MetaDto {
	if m.TotalCount != nil {
		m.HasNext = *m.TotalCount > int64(m.Page*m.Limit)
	} else {
		m.HasNext = dataLength == m.Limit
	}
	return m
}

type QueryOption func(*QueryDto)

// createdAt desc
func DefaultSort(sort string) QueryOption {
	return func(q *QueryDto) {
		q.DefaultSort(sort)
	}
}

func (q *QueryDto) DefaultSort(sort string) {
	if len(q.Sorts) == 0 || q.Sorts[0] == "" {
		q.Sorts = []string{sort}
	}
}

func ParseBody(c echo.Context, opts ...QueryOption) (*QueryDto, error) {
	payload := &QueryDto{
		Page:  1,
		Limit: 20,
	}

	if err := c.Bind(payload); err != nil {
		return payload, err
	}

	for _, opt := range opts {
		opt(payload)
	}

	return payload, nil
}

func (q *QueryDto) DefaultSelects(s string) {
	q.Selects = strings.Split(s, ",")
	// trim space
	for i, v := range q.Selects {
		q.Selects[i] = strings.TrimSpace(v)
	}
	if len(q.Selects) == 0 {
		q.Selects = []string{"list", "totalCount"}
	}
}
