// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/genre"
	"github.com/gva/internal/ent/manga"
	"github.com/gva/internal/ent/mangachapter"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/predicate"
	"github.com/gva/internal/ent/role"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AdminFunc type is an adapter to allow the use of ordinary function as a Querier.
type AdminFunc func(context.Context, *ent.AdminQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AdminFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AdminQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AdminQuery", q)
}

// The TraverseAdmin type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAdmin func(context.Context, *ent.AdminQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAdmin) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAdmin) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AdminQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AdminQuery", q)
}

// The DepartmentFunc type is an adapter to allow the use of ordinary function as a Querier.
type DepartmentFunc func(context.Context, *ent.DepartmentQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DepartmentFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DepartmentQuery", q)
}

// The TraverseDepartment type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDepartment func(context.Context, *ent.DepartmentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDepartment) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDepartment) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DepartmentQuery", q)
}

// The GenreFunc type is an adapter to allow the use of ordinary function as a Querier.
type GenreFunc func(context.Context, *ent.GenreQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f GenreFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.GenreQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.GenreQuery", q)
}

// The TraverseGenre type is an adapter to allow the use of ordinary function as Traverser.
type TraverseGenre func(context.Context, *ent.GenreQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseGenre) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseGenre) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenreQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.GenreQuery", q)
}

// The MangaFunc type is an adapter to allow the use of ordinary function as a Querier.
type MangaFunc func(context.Context, *ent.MangaQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MangaFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MangaQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MangaQuery", q)
}

// The TraverseManga type is an adapter to allow the use of ordinary function as Traverser.
type TraverseManga func(context.Context, *ent.MangaQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseManga) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseManga) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MangaQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MangaQuery", q)
}

// The MangaChapterFunc type is an adapter to allow the use of ordinary function as a Querier.
type MangaChapterFunc func(context.Context, *ent.MangaChapterQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MangaChapterFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MangaChapterQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MangaChapterQuery", q)
}

// The TraverseMangaChapter type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMangaChapter func(context.Context, *ent.MangaChapterQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMangaChapter) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMangaChapter) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MangaChapterQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MangaChapterQuery", q)
}

// The PermissionFunc type is an adapter to allow the use of ordinary function as a Querier.
type PermissionFunc func(context.Context, *ent.PermissionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PermissionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PermissionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PermissionQuery", q)
}

// The TraversePermission type is an adapter to allow the use of ordinary function as Traverser.
type TraversePermission func(context.Context, *ent.PermissionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePermission) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePermission) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PermissionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PermissionQuery", q)
}

// The RoleFunc type is an adapter to allow the use of ordinary function as a Querier.
type RoleFunc func(context.Context, *ent.RoleQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f RoleFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.RoleQuery", q)
}

// The TraverseRole type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRole func(context.Context, *ent.RoleQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRole) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRole) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.RoleQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.AdminQuery:
		return &query[*ent.AdminQuery, predicate.Admin, admin.OrderOption]{typ: ent.TypeAdmin, tq: q}, nil
	case *ent.DepartmentQuery:
		return &query[*ent.DepartmentQuery, predicate.Department, department.OrderOption]{typ: ent.TypeDepartment, tq: q}, nil
	case *ent.GenreQuery:
		return &query[*ent.GenreQuery, predicate.Genre, genre.OrderOption]{typ: ent.TypeGenre, tq: q}, nil
	case *ent.MangaQuery:
		return &query[*ent.MangaQuery, predicate.Manga, manga.OrderOption]{typ: ent.TypeManga, tq: q}, nil
	case *ent.MangaChapterQuery:
		return &query[*ent.MangaChapterQuery, predicate.MangaChapter, mangachapter.OrderOption]{typ: ent.TypeMangaChapter, tq: q}, nil
	case *ent.PermissionQuery:
		return &query[*ent.PermissionQuery, predicate.Permission, permission.OrderOption]{typ: ent.TypePermission, tq: q}, nil
	case *ent.RoleQuery:
		return &query[*ent.RoleQuery, predicate.Role, role.OrderOption]{typ: ent.TypeRole, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
