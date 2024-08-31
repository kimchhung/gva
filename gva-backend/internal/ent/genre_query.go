// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/genre"
	"github.com/gva/internal/ent/manga"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// GenreQuery is the builder for querying Genre entities.
type GenreQuery struct {
	config
	ctx             *QueryContext
	order           []genre.OrderOption
	inters          []Interceptor
	predicates      []predicate.Genre
	withMangas      *MangaQuery
	loadTotal       []func(context.Context, []*Genre) error
	modifiers       []func(*sql.Selector)
	withNamedMangas map[string]*MangaQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GenreQuery builder.
func (gq *GenreQuery) Where(ps ...predicate.Genre) *GenreQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GenreQuery) Limit(limit int) *GenreQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GenreQuery) Offset(offset int) *GenreQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GenreQuery) Unique(unique bool) *GenreQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GenreQuery) Order(o ...genre.OrderOption) *GenreQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryMangas chains the current query on the "mangas" edge.
func (gq *GenreQuery) QueryMangas() *MangaQuery {
	query := (&MangaClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(genre.Table, genre.FieldID, selector),
			sqlgraph.To(manga.Table, manga.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, genre.MangasTable, genre.MangasPrimaryKey...),
		)
		schemaConfig := gq.schemaConfig
		step.To.Schema = schemaConfig.Manga
		step.Edge.Schema = schemaConfig.GenreMangas
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Genre entity from the query.
// Returns a *NotFoundError when no Genre was found.
func (gq *GenreQuery) First(ctx context.Context) (*Genre, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{genre.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GenreQuery) FirstX(ctx context.Context) *Genre {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Genre ID from the query.
// Returns a *NotFoundError when no Genre ID was found.
func (gq *GenreQuery) FirstID(ctx context.Context) (id pxid.ID, err error) {
	var ids []pxid.ID
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{genre.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GenreQuery) FirstIDX(ctx context.Context) pxid.ID {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Genre entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Genre entity is found.
// Returns a *NotFoundError when no Genre entities are found.
func (gq *GenreQuery) Only(ctx context.Context) (*Genre, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{genre.Label}
	default:
		return nil, &NotSingularError{genre.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GenreQuery) OnlyX(ctx context.Context) *Genre {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Genre ID in the query.
// Returns a *NotSingularError when more than one Genre ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GenreQuery) OnlyID(ctx context.Context) (id pxid.ID, err error) {
	var ids []pxid.ID
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{genre.Label}
	default:
		err = &NotSingularError{genre.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GenreQuery) OnlyIDX(ctx context.Context) pxid.ID {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Genres.
func (gq *GenreQuery) All(ctx context.Context) ([]*Genre, error) {
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryAll)
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Genre, *GenreQuery]()
	return withInterceptors[[]*Genre](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GenreQuery) AllX(ctx context.Context) []*Genre {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Genre IDs.
func (gq *GenreQuery) IDs(ctx context.Context) (ids []pxid.ID, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryIDs)
	if err = gq.Select(genre.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GenreQuery) IDsX(ctx context.Context) []pxid.ID {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GenreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryCount)
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GenreQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GenreQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GenreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryExist)
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GenreQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GenreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GenreQuery) Clone() *GenreQuery {
	if gq == nil {
		return nil
	}
	return &GenreQuery{
		config:     gq.config,
		ctx:        gq.ctx.Clone(),
		order:      append([]genre.OrderOption{}, gq.order...),
		inters:     append([]Interceptor{}, gq.inters...),
		predicates: append([]predicate.Genre{}, gq.predicates...),
		withMangas: gq.withMangas.Clone(),
		// clone intermediate query.
		sql:       gq.sql.Clone(),
		path:      gq.path,
		modifiers: append([]func(*sql.Selector){}, gq.modifiers...),
	}
}

// WithMangas tells the query-builder to eager-load the nodes that are connected to
// the "mangas" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GenreQuery) WithMangas(opts ...func(*MangaQuery)) *GenreQuery {
	query := (&MangaClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withMangas = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Genre.Query().
//		GroupBy(genre.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GenreQuery) GroupBy(field string, fields ...string) *GenreGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GenreGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = genre.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
//	}
//
//	client.Genre.Query().
//		Select(genre.FieldCreatedAt).
//		Scan(ctx, &v)
func (gq *GenreQuery) Select(fields ...string) *GenreSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GenreSelect{GenreQuery: gq}
	sbuild.label = genre.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GenreSelect configured with the given aggregations.
func (gq *GenreQuery) Aggregate(fns ...AggregateFunc) *GenreSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GenreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !genre.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GenreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Genre, error) {
	var (
		nodes       = []*Genre{}
		_spec       = gq.querySpec()
		loadedTypes = [1]bool{
			gq.withMangas != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Genre).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Genre{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = gq.schemaConfig.Genre
	ctx = internal.NewSchemaConfigContext(ctx, gq.schemaConfig)
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withMangas; query != nil {
		if err := gq.loadMangas(ctx, query, nodes,
			func(n *Genre) { n.Edges.Mangas = []*Manga{} },
			func(n *Genre, e *Manga) { n.Edges.Mangas = append(n.Edges.Mangas, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range gq.withNamedMangas {
		if err := gq.loadMangas(ctx, query, nodes,
			func(n *Genre) { n.appendNamedMangas(name) },
			func(n *Genre, e *Manga) { n.appendNamedMangas(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range gq.loadTotal {
		if err := gq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GenreQuery) loadMangas(ctx context.Context, query *MangaQuery, nodes []*Genre, init func(*Genre), assign func(*Genre, *Manga)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pxid.ID]*Genre)
	nids := make(map[pxid.ID]map[*Genre]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(genre.MangasTable)
		joinT.Schema(gq.schemaConfig.GenreMangas)
		s.Join(joinT).On(s.C(manga.FieldID), joinT.C(genre.MangasPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(genre.MangasPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(genre.MangasPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(pxid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*pxid.ID)
				inValue := *values[1].(*pxid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Genre]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Manga](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "mangas" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (gq *GenreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Schema = gq.schemaConfig.Genre
	ctx = internal.NewSchemaConfigContext(ctx, gq.schemaConfig)
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GenreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(genre.Table, genre.Columns, sqlgraph.NewFieldSpec(genre.FieldID, field.TypeString))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, genre.FieldID)
		for i := range fields {
			if fields[i] != genre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GenreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(genre.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = genre.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(gq.schemaConfig.Genre)
	ctx = internal.NewSchemaConfigContext(ctx, gq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gq *GenreQuery) ForUpdate(opts ...sql.LockOption) *GenreQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gq *GenreQuery) ForShare(opts ...sql.LockOption) *GenreQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gq *GenreQuery) Modify(modifiers ...func(s *sql.Selector)) *GenreSelect {
	gq.modifiers = append(gq.modifiers, modifiers...)
	return gq.Select()
}

// WithNamedMangas tells the query-builder to eager-load the nodes that are connected to the "mangas"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (gq *GenreQuery) WithNamedMangas(name string, opts ...func(*MangaQuery)) *GenreQuery {
	query := (&MangaClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if gq.withNamedMangas == nil {
		gq.withNamedMangas = make(map[string]*MangaQuery)
	}
	gq.withNamedMangas[name] = query
	return gq
}

// GenreGroupBy is the group-by builder for Genre entities.
type GenreGroupBy struct {
	selector
	build *GenreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GenreGroupBy) Aggregate(fns ...AggregateFunc) *GenreGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GenreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, ent.OpQueryGroupBy)
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GenreQuery, *GenreGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GenreGroupBy) sqlScan(ctx context.Context, root *GenreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GenreSelect is the builder for selecting fields of Genre entities.
type GenreSelect struct {
	*GenreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GenreSelect) Aggregate(fns ...AggregateFunc) *GenreSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GenreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, ent.OpQuerySelect)
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GenreQuery, *GenreSelect](ctx, gs.GenreQuery, gs, gs.inters, v)
}

func (gs *GenreSelect) sqlScan(ctx context.Context, root *GenreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gs *GenreSelect) Modify(modifiers ...func(s *sql.Selector)) *GenreSelect {
	gs.modifiers = append(gs.modifiers, modifiers...)
	return gs
}