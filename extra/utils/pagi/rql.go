package pagi

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/extra/internal/ent/intercept"
	"github.com/kimchhung/gva/extra/internal/rql"
)

func RQL[T any](query T, p *rql.Params, opts ...func(*sql.Selector)) T {
	q, err := intercept.NewQuery(query)
	if err != nil {
		panic(fmt.Errorf("invalid rql %v", err))
	}

	scripts := []func(*sql.Selector){}
	if p.FilterExp != "" {
		scripts = append(scripts,
			func(s *sql.Selector) { s.Where(sql.ExprP(p.FilterExp.String(), p.FilterArgs...)) },
		)
	}

	if len(p.Sort) > 0 {
		scripts = append(scripts, func(s *sql.Selector) { s.OrderBy(p.Sort...) })
	}

	if len(p.Select) > 0 {
		scripts = append(scripts, func(s *sql.Selector) { s.Select(p.Select...) })
	}

	if len(p.Group) > 0 {
		scripts = append(scripts, func(s *sql.Selector) { s.GroupBy(p.Group...) })
	}

	if len(opts) > 0 {
		scripts = append(scripts, opts...)
	}

	if len(scripts) > 0 {
		q.WhereP(scripts...)
	}

	q.Limit(p.Limit)
	q.Offset(p.Offset)
	return query
}

func RQLCount[T any](query T, p *rql.Params, opts ...func(*sql.Selector)) T {
	q, err := intercept.NewQuery(query)
	if err != nil {
		panic(fmt.Errorf("invalid rql %v", err))
	}

	scripts := []func(*sql.Selector){}
	if p.FilterExp != "" {
		scripts = append(scripts,
			func(s *sql.Selector) { s.Where(sql.ExprP(p.FilterExp.String(), p.FilterArgs...)) },
		)
	}

	if len(p.Sort) > 0 {
		scripts = append(scripts, func(s *sql.Selector) { s.OrderBy(p.Sort...) })
	}

	if len(p.Select) > 0 {
		scripts = append(scripts, func(s *sql.Selector) { s.Select(p.Select...) })
	}

	if len(p.Group) > 0 {
		scripts = append(scripts, func(s *sql.Selector) { s.GroupBy(p.Group...) })
	}

	if len(opts) > 0 {
		scripts = append(scripts, opts...)
	}

	if len(scripts) > 0 {
		q.WhereP(scripts...)
	}

	return query
}
