package pagi

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/backend-echo/internal/ent/intercept"
)

type InterceptorOption func(*sql.Selector)

func WithFilter(filterExp string, filterArgs []interface{}) InterceptorOption {
	return func(s *sql.Selector) {
		if filterExp != "" {
			s.Where(sql.ExprP(filterExp, filterArgs...))
		}
	}
}

func WithSelect(selects ...string) InterceptorOption {
	return func(s *sql.Selector) {
		if len(selects) > 0 {
			s.Select(selects...)
		}
	}
}

func WithSort(sort ...string) InterceptorOption {
	return func(s *sql.Selector) {
		if len(sort) > 0 {
			for _, _s := range sort {
				parts := strings.Split(_s, " ")

				s.OrderExprFunc(func(b *sql.Builder) {
					b.Ident(parts[0])
					if len(parts) == 2 {
						b.WriteString(" " + parts[1])
					}
				})
			}
		}
	}
}

func WithLimitOffset(limit int, offset int) InterceptorOption {
	return func(s *sql.Selector) {
		s.Limit(limit).Offset(offset)
	}
}

func WithInterceptor[T any](query T, opts ...InterceptorOption) T {
	q, err := intercept.NewQuery(query)
	if err != nil {
		panic(fmt.Errorf("can't user interceptor %v", err))
	}

	for _, opt := range opts {
		q.WhereP(opt)
	}

	return query
}
