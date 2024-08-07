// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/predicate"
	"github.com/gva/internal/ent/region"

	"github.com/gva/internal/ent/internal"
)

// RegionQuery is the builder for querying Region entities.
type RegionQuery struct {
	config
	ctx               *QueryContext
	order             []region.OrderOption
	inters            []Interceptor
	predicates        []predicate.Region
	withParent        *RegionQuery
	withChildren      *RegionQuery
	loadTotal         []func(context.Context, []*Region) error
	modifiers         []func(*sql.Selector)
	withNamedChildren map[string]*RegionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RegionQuery builder.
func (rq *RegionQuery) Where(ps ...predicate.Region) *RegionQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RegionQuery) Limit(limit int) *RegionQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RegionQuery) Offset(offset int) *RegionQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RegionQuery) Unique(unique bool) *RegionQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RegionQuery) Order(o ...region.OrderOption) *RegionQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryParent chains the current query on the "parent" edge.
func (rq *RegionQuery) QueryParent() *RegionQuery {
	query := (&RegionClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(region.Table, region.FieldID, selector),
			sqlgraph.To(region.Table, region.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, region.ParentTable, region.ParentColumn),
		)
		schemaConfig := rq.schemaConfig
		step.To.Schema = schemaConfig.Region
		step.Edge.Schema = schemaConfig.Region
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (rq *RegionQuery) QueryChildren() *RegionQuery {
	query := (&RegionClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(region.Table, region.FieldID, selector),
			sqlgraph.To(region.Table, region.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, region.ChildrenTable, region.ChildrenColumn),
		)
		schemaConfig := rq.schemaConfig
		step.To.Schema = schemaConfig.Region
		step.Edge.Schema = schemaConfig.Region
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Region entity from the query.
// Returns a *NotFoundError when no Region was found.
func (rq *RegionQuery) First(ctx context.Context) (*Region, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{region.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RegionQuery) FirstX(ctx context.Context) *Region {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Region ID from the query.
// Returns a *NotFoundError when no Region ID was found.
func (rq *RegionQuery) FirstID(ctx context.Context) (id pxid.ID, err error) {
	var ids []pxid.ID
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{region.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RegionQuery) FirstIDX(ctx context.Context) pxid.ID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Region entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Region entity is found.
// Returns a *NotFoundError when no Region entities are found.
func (rq *RegionQuery) Only(ctx context.Context) (*Region, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{region.Label}
	default:
		return nil, &NotSingularError{region.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RegionQuery) OnlyX(ctx context.Context) *Region {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Region ID in the query.
// Returns a *NotSingularError when more than one Region ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RegionQuery) OnlyID(ctx context.Context) (id pxid.ID, err error) {
	var ids []pxid.ID
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{region.Label}
	default:
		err = &NotSingularError{region.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RegionQuery) OnlyIDX(ctx context.Context) pxid.ID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Regions.
func (rq *RegionQuery) All(ctx context.Context) ([]*Region, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Region, *RegionQuery]()
	return withInterceptors[[]*Region](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RegionQuery) AllX(ctx context.Context) []*Region {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Region IDs.
func (rq *RegionQuery) IDs(ctx context.Context) (ids []pxid.ID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(region.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RegionQuery) IDsX(ctx context.Context) []pxid.ID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RegionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RegionQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RegionQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RegionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RegionQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RegionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RegionQuery) Clone() *RegionQuery {
	if rq == nil {
		return nil
	}
	return &RegionQuery{
		config:       rq.config,
		ctx:          rq.ctx.Clone(),
		order:        append([]region.OrderOption{}, rq.order...),
		inters:       append([]Interceptor{}, rq.inters...),
		predicates:   append([]predicate.Region{}, rq.predicates...),
		withParent:   rq.withParent.Clone(),
		withChildren: rq.withChildren.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RegionQuery) WithParent(opts ...func(*RegionQuery)) *RegionQuery {
	query := (&RegionClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withParent = query
	return rq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RegionQuery) WithChildren(opts ...func(*RegionQuery)) *RegionQuery {
	query := (&RegionClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withChildren = query
	return rq
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
//	client.Region.Query().
//		GroupBy(region.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RegionQuery) GroupBy(field string, fields ...string) *RegionGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RegionGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = region.Label
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
//	client.Region.Query().
//		Select(region.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *RegionQuery) Select(fields ...string) *RegionSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RegionSelect{RegionQuery: rq}
	sbuild.label = region.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RegionSelect configured with the given aggregations.
func (rq *RegionQuery) Aggregate(fns ...AggregateFunc) *RegionSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RegionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !region.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RegionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Region, error) {
	var (
		nodes       = []*Region{}
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withParent != nil,
			rq.withChildren != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Region).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Region{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = rq.schemaConfig.Region
	ctx = internal.NewSchemaConfigContext(ctx, rq.schemaConfig)
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withParent; query != nil {
		if err := rq.loadParent(ctx, query, nodes, nil,
			func(n *Region, e *Region) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withChildren; query != nil {
		if err := rq.loadChildren(ctx, query, nodes,
			func(n *Region) { n.Edges.Children = []*Region{} },
			func(n *Region, e *Region) {
				n.Edges.Children = append(n.Edges.Children, e)
				if !e.Edges.loadedTypes[0] {
					e.Edges.Parent = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedChildren {
		if err := rq.loadChildren(ctx, query, nodes,
			func(n *Region) { n.appendNamedChildren(name) },
			func(n *Region, e *Region) {
				n.appendNamedChildren(name, e)
				if !e.Edges.loadedTypes[0] {
					e.Edges.Parent = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for i := range rq.loadTotal {
		if err := rq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RegionQuery) loadParent(ctx context.Context, query *RegionQuery, nodes []*Region, init func(*Region), assign func(*Region, *Region)) error {
	ids := make([]pxid.ID, 0, len(nodes))
	nodeids := make(map[pxid.ID][]*Region)
	for i := range nodes {
		if nodes[i].Pid == nil {
			continue
		}
		fk := *nodes[i].Pid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(region.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RegionQuery) loadChildren(ctx context.Context, query *RegionQuery, nodes []*Region, init func(*Region), assign func(*Region, *Region)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pxid.ID]*Region)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(region.FieldPid)
	}
	query.Where(predicate.Region(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(region.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.Pid
		if fk == nil {
			return fmt.Errorf(`foreign-key "pid" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "pid" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RegionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Schema = rq.schemaConfig.Region
	ctx = internal.NewSchemaConfigContext(ctx, rq.schemaConfig)
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RegionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(region.Table, region.Columns, sqlgraph.NewFieldSpec(region.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, region.FieldID)
		for i := range fields {
			if fields[i] != region.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rq.withParent != nil {
			_spec.Node.AddColumnOnce(region.FieldPid)
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RegionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(region.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = region.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(rq.schemaConfig.Region)
	ctx = internal.NewSchemaConfigContext(ctx, rq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range rq.modifiers {
		m(selector)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rq *RegionQuery) ForUpdate(opts ...sql.LockOption) *RegionQuery {
	if rq.driver.Dialect() == dialect.Postgres {
		rq.Unique(false)
	}
	rq.modifiers = append(rq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rq *RegionQuery) ForShare(opts ...sql.LockOption) *RegionQuery {
	if rq.driver.Dialect() == dialect.Postgres {
		rq.Unique(false)
	}
	rq.modifiers = append(rq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rq *RegionQuery) Modify(modifiers ...func(s *sql.Selector)) *RegionSelect {
	rq.modifiers = append(rq.modifiers, modifiers...)
	return rq.Select()
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RegionQuery) WithNamedChildren(name string, opts ...func(*RegionQuery)) *RegionQuery {
	query := (&RegionClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedChildren == nil {
		rq.withNamedChildren = make(map[string]*RegionQuery)
	}
	rq.withNamedChildren[name] = query
	return rq
}

// RegionGroupBy is the group-by builder for Region entities.
type RegionGroupBy struct {
	selector
	build *RegionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RegionGroupBy) Aggregate(fns ...AggregateFunc) *RegionGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RegionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RegionQuery, *RegionGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RegionGroupBy) sqlScan(ctx context.Context, root *RegionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RegionSelect is the builder for selecting fields of Region entities.
type RegionSelect struct {
	*RegionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RegionSelect) Aggregate(fns ...AggregateFunc) *RegionSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RegionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RegionQuery, *RegionSelect](ctx, rs.RegionQuery, rs, rs.inters, v)
}

func (rs *RegionSelect) sqlScan(ctx context.Context, root *RegionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *RegionSelect) Modify(modifiers ...func(s *sql.Selector)) *RegionSelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
