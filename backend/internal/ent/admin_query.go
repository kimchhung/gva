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
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/predicate"
	"github.com/gva/internal/ent/role"

	"github.com/gva/internal/ent/internal"
)

// AdminQuery is the builder for querying Admin entities.
type AdminQuery struct {
	config
	ctx            *QueryContext
	order          []admin.OrderOption
	inters         []Interceptor
	predicates     []predicate.Admin
	withRoles      *RoleQuery
	loadTotal      []func(context.Context, []*Admin) error
	modifiers      []func(*sql.Selector)
	withNamedRoles map[string]*RoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminQuery builder.
func (aq *AdminQuery) Where(ps ...predicate.Admin) *AdminQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AdminQuery) Limit(limit int) *AdminQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AdminQuery) Offset(offset int) *AdminQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AdminQuery) Unique(unique bool) *AdminQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AdminQuery) Order(o ...admin.OrderOption) *AdminQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryRoles chains the current query on the "roles" edge.
func (aq *AdminQuery) QueryRoles() *RoleQuery {
	query := (&RoleClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(admin.Table, admin.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, admin.RolesTable, admin.RolesPrimaryKey...),
		)
		schemaConfig := aq.schemaConfig
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.AdminRoles
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Admin entity from the query.
// Returns a *NotFoundError when no Admin was found.
func (aq *AdminQuery) First(ctx context.Context) (*Admin, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{admin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AdminQuery) FirstX(ctx context.Context) *Admin {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Admin ID from the query.
// Returns a *NotFoundError when no Admin ID was found.
func (aq *AdminQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{admin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AdminQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Admin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Admin entity is found.
// Returns a *NotFoundError when no Admin entities are found.
func (aq *AdminQuery) Only(ctx context.Context) (*Admin, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{admin.Label}
	default:
		return nil, &NotSingularError{admin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AdminQuery) OnlyX(ctx context.Context) *Admin {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Admin ID in the query.
// Returns a *NotSingularError when more than one Admin ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AdminQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{admin.Label}
	default:
		err = &NotSingularError{admin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AdminQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Admins.
func (aq *AdminQuery) All(ctx context.Context) ([]*Admin, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Admin, *AdminQuery]()
	return withInterceptors[[]*Admin](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AdminQuery) AllX(ctx context.Context) []*Admin {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Admin IDs.
func (aq *AdminQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(admin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AdminQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AdminQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AdminQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AdminQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AdminQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AdminQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AdminQuery) Clone() *AdminQuery {
	if aq == nil {
		return nil
	}
	return &AdminQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]admin.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Admin{}, aq.predicates...),
		withRoles:  aq.withRoles.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AdminQuery) WithRoles(opts ...func(*RoleQuery)) *AdminQuery {
	query := (&RoleClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withRoles = query
	return aq
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
//	client.Admin.Query().
//		GroupBy(admin.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AdminQuery) GroupBy(field string, fields ...string) *AdminGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = admin.Label
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
//	client.Admin.Query().
//		Select(admin.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AdminQuery) Select(fields ...string) *AdminSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AdminSelect{AdminQuery: aq}
	sbuild.label = admin.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminSelect configured with the given aggregations.
func (aq *AdminQuery) Aggregate(fns ...AggregateFunc) *AdminSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AdminQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !admin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AdminQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Admin, error) {
	var (
		nodes       = []*Admin{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Admin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Admin{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = aq.schemaConfig.Admin
	ctx = internal.NewSchemaConfigContext(ctx, aq.schemaConfig)
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withRoles; query != nil {
		if err := aq.loadRoles(ctx, query, nodes,
			func(n *Admin) { n.Edges.Roles = []*Role{} },
			func(n *Admin, e *Role) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedRoles {
		if err := aq.loadRoles(ctx, query, nodes,
			func(n *Admin) { n.appendNamedRoles(name) },
			func(n *Admin, e *Role) { n.appendNamedRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AdminQuery) loadRoles(ctx context.Context, query *RoleQuery, nodes []*Admin, init func(*Admin), assign func(*Admin, *Role)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Admin)
	nids := make(map[xid.ID]map[*Admin]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(admin.RolesTable)
		joinT.Schema(aq.schemaConfig.AdminRoles)
		s.Join(joinT).On(s.C(role.FieldID), joinT.C(admin.RolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(admin.RolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(admin.RolesPrimaryKey[0]))
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
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Admin]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Role](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aq *AdminQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Schema = aq.schemaConfig.Admin
	ctx = internal.NewSchemaConfigContext(ctx, aq.schemaConfig)
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AdminQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeString))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admin.FieldID)
		for i := range fields {
			if fields[i] != admin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AdminQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(admin.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = admin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(aq.schemaConfig.Admin)
	ctx = internal.NewSchemaConfigContext(ctx, aq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (aq *AdminQuery) ForUpdate(opts ...sql.LockOption) *AdminQuery {
	if aq.driver.Dialect() == dialect.Postgres {
		aq.Unique(false)
	}
	aq.modifiers = append(aq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return aq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (aq *AdminQuery) ForShare(opts ...sql.LockOption) *AdminQuery {
	if aq.driver.Dialect() == dialect.Postgres {
		aq.Unique(false)
	}
	aq.modifiers = append(aq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *AdminQuery) Modify(modifiers ...func(s *sql.Selector)) *AdminSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// WithNamedRoles tells the query-builder to eager-load the nodes that are connected to the "roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AdminQuery) WithNamedRoles(name string, opts ...func(*RoleQuery)) *AdminQuery {
	query := (&RoleClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedRoles == nil {
		aq.withNamedRoles = make(map[string]*RoleQuery)
	}
	aq.withNamedRoles[name] = query
	return aq
}

// AdminGroupBy is the group-by builder for Admin entities.
type AdminGroupBy struct {
	selector
	build *AdminQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AdminGroupBy) Aggregate(fns ...AggregateFunc) *AdminGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AdminGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminQuery, *AdminGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AdminGroupBy) sqlScan(ctx context.Context, root *AdminQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminSelect is the builder for selecting fields of Admin entities.
type AdminSelect struct {
	*AdminQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AdminSelect) Aggregate(fns ...AggregateFunc) *AdminSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AdminSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminQuery, *AdminSelect](ctx, as.AdminQuery, as, as.inters, v)
}

func (as *AdminSelect) sqlScan(ctx context.Context, root *AdminQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *AdminSelect) Modify(modifiers ...func(s *sql.Selector)) *AdminSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}
