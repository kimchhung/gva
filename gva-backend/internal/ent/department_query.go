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
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// DepartmentQuery is the builder for querying Department entities.
type DepartmentQuery struct {
	config
	ctx               *QueryContext
	order             []department.OrderOption
	inters            []Interceptor
	predicates        []predicate.Department
	withParent        *DepartmentQuery
	withChildren      *DepartmentQuery
	withMembers       *AdminQuery
	loadTotal         []func(context.Context, []*Department) error
	modifiers         []func(*sql.Selector)
	withNamedChildren map[string]*DepartmentQuery
	withNamedMembers  map[string]*AdminQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DepartmentQuery builder.
func (dq *DepartmentQuery) Where(ps ...predicate.Department) *DepartmentQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DepartmentQuery) Limit(limit int) *DepartmentQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DepartmentQuery) Offset(offset int) *DepartmentQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DepartmentQuery) Unique(unique bool) *DepartmentQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DepartmentQuery) Order(o ...department.OrderOption) *DepartmentQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryParent chains the current query on the "parent" edge.
func (dq *DepartmentQuery) QueryParent() *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, department.ParentTable, department.ParentColumn),
		)
		schemaConfig := dq.schemaConfig
		step.To.Schema = schemaConfig.Department
		step.Edge.Schema = schemaConfig.Department
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (dq *DepartmentQuery) QueryChildren() *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.ChildrenTable, department.ChildrenColumn),
		)
		schemaConfig := dq.schemaConfig
		step.To.Schema = schemaConfig.Department
		step.Edge.Schema = schemaConfig.Department
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMembers chains the current query on the "members" edge.
func (dq *DepartmentQuery) QueryMembers() *AdminQuery {
	query := (&AdminClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.MembersTable, department.MembersColumn),
		)
		schemaConfig := dq.schemaConfig
		step.To.Schema = schemaConfig.Admin
		step.Edge.Schema = schemaConfig.Admin
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Department entity from the query.
// Returns a *NotFoundError when no Department was found.
func (dq *DepartmentQuery) First(ctx context.Context) (*Department, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{department.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DepartmentQuery) FirstX(ctx context.Context) *Department {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Department ID from the query.
// Returns a *NotFoundError when no Department ID was found.
func (dq *DepartmentQuery) FirstID(ctx context.Context) (id pxid.ID, err error) {
	var ids []pxid.ID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{department.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DepartmentQuery) FirstIDX(ctx context.Context) pxid.ID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Department entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Department entity is found.
// Returns a *NotFoundError when no Department entities are found.
func (dq *DepartmentQuery) Only(ctx context.Context) (*Department, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{department.Label}
	default:
		return nil, &NotSingularError{department.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DepartmentQuery) OnlyX(ctx context.Context) *Department {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Department ID in the query.
// Returns a *NotSingularError when more than one Department ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DepartmentQuery) OnlyID(ctx context.Context) (id pxid.ID, err error) {
	var ids []pxid.ID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = &NotSingularError{department.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DepartmentQuery) OnlyIDX(ctx context.Context) pxid.ID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Departments.
func (dq *DepartmentQuery) All(ctx context.Context) ([]*Department, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Department, *DepartmentQuery]()
	return withInterceptors[[]*Department](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DepartmentQuery) AllX(ctx context.Context) []*Department {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Department IDs.
func (dq *DepartmentQuery) IDs(ctx context.Context) (ids []pxid.ID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(department.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DepartmentQuery) IDsX(ctx context.Context) []pxid.ID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DepartmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DepartmentQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DepartmentQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DepartmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DepartmentQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DepartmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DepartmentQuery) Clone() *DepartmentQuery {
	if dq == nil {
		return nil
	}
	return &DepartmentQuery{
		config:       dq.config,
		ctx:          dq.ctx.Clone(),
		order:        append([]department.OrderOption{}, dq.order...),
		inters:       append([]Interceptor{}, dq.inters...),
		predicates:   append([]predicate.Department{}, dq.predicates...),
		withParent:   dq.withParent.Clone(),
		withChildren: dq.withChildren.Clone(),
		withMembers:  dq.withMembers.Clone(),
		// clone intermediate query.
		sql:       dq.sql.Clone(),
		path:      dq.path,
		modifiers: append([]func(*sql.Selector){}, dq.modifiers...),
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithParent(opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withParent = query
	return dq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithChildren(opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withChildren = query
	return dq
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithMembers(opts ...func(*AdminQuery)) *DepartmentQuery {
	query := (&AdminClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withMembers = query
	return dq
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
//	client.Department.Query().
//		GroupBy(department.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DepartmentQuery) GroupBy(field string, fields ...string) *DepartmentGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DepartmentGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = department.Label
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
//	client.Department.Query().
//		Select(department.FieldCreatedAt).
//		Scan(ctx, &v)
func (dq *DepartmentQuery) Select(fields ...string) *DepartmentSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DepartmentSelect{DepartmentQuery: dq}
	sbuild.label = department.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DepartmentSelect configured with the given aggregations.
func (dq *DepartmentQuery) Aggregate(fns ...AggregateFunc) *DepartmentSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DepartmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !department.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DepartmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Department, error) {
	var (
		nodes       = []*Department{}
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withParent != nil,
			dq.withChildren != nil,
			dq.withMembers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Department).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Department{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = dq.schemaConfig.Department
	ctx = internal.NewSchemaConfigContext(ctx, dq.schemaConfig)
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withParent; query != nil {
		if err := dq.loadParent(ctx, query, nodes, nil,
			func(n *Department, e *Department) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withChildren; query != nil {
		if err := dq.loadChildren(ctx, query, nodes,
			func(n *Department) { n.Edges.Children = []*Department{} },
			func(n *Department, e *Department) {
				n.Edges.Children = append(n.Edges.Children, e)
				if !e.Edges.loadedTypes[0] {
					e.Edges.Parent = n
				}
			}); err != nil {
			return nil, err
		}
	}
	if query := dq.withMembers; query != nil {
		if err := dq.loadMembers(ctx, query, nodes,
			func(n *Department) { n.Edges.Members = []*Admin{} },
			func(n *Department, e *Admin) {
				n.Edges.Members = append(n.Edges.Members, e)
				if !e.Edges.loadedTypes[1] {
					e.Edges.Department = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedChildren {
		if err := dq.loadChildren(ctx, query, nodes,
			func(n *Department) { n.appendNamedChildren(name) },
			func(n *Department, e *Department) {
				n.appendNamedChildren(name, e)
				if !e.Edges.loadedTypes[0] {
					e.Edges.Parent = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedMembers {
		if err := dq.loadMembers(ctx, query, nodes,
			func(n *Department) { n.appendNamedMembers(name) },
			func(n *Department, e *Admin) {
				n.appendNamedMembers(name, e)
				if !e.Edges.loadedTypes[1] {
					e.Edges.Department = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DepartmentQuery) loadParent(ctx context.Context, query *DepartmentQuery, nodes []*Department, init func(*Department), assign func(*Department, *Department)) error {
	ids := make([]pxid.ID, 0, len(nodes))
	nodeids := make(map[pxid.ID][]*Department)
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
	query.Where(department.IDIn(ids...))
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
func (dq *DepartmentQuery) loadChildren(ctx context.Context, query *DepartmentQuery, nodes []*Department, init func(*Department), assign func(*Department, *Department)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pxid.ID]*Department)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(department.FieldPid)
	}
	query.Where(predicate.Department(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(department.ChildrenColumn), fks...))
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
func (dq *DepartmentQuery) loadMembers(ctx context.Context, query *AdminQuery, nodes []*Department, init func(*Department), assign func(*Department, *Admin)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pxid.ID]*Department)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(admin.FieldDepartmentID)
	}
	query.Where(predicate.Admin(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(department.MembersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DepartmentID
		if fk == nil {
			return fmt.Errorf(`foreign-key "department_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "department_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DepartmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Schema = dq.schemaConfig.Department
	ctx = internal.NewSchemaConfigContext(ctx, dq.schemaConfig)
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DepartmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(department.Table, department.Columns, sqlgraph.NewFieldSpec(department.FieldID, field.TypeString))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, department.FieldID)
		for i := range fields {
			if fields[i] != department.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dq.withParent != nil {
			_spec.Node.AddColumnOnce(department.FieldPid)
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DepartmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(department.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = department.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(dq.schemaConfig.Department)
	ctx = internal.NewSchemaConfigContext(ctx, dq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range dq.modifiers {
		m(selector)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dq *DepartmentQuery) ForUpdate(opts ...sql.LockOption) *DepartmentQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dq *DepartmentQuery) ForShare(opts ...sql.LockOption) *DepartmentQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dq *DepartmentQuery) Modify(modifiers ...func(s *sql.Selector)) *DepartmentSelect {
	dq.modifiers = append(dq.modifiers, modifiers...)
	return dq.Select()
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithNamedChildren(name string, opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedChildren == nil {
		dq.withNamedChildren = make(map[string]*DepartmentQuery)
	}
	dq.withNamedChildren[name] = query
	return dq
}

// WithNamedMembers tells the query-builder to eager-load the nodes that are connected to the "members"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithNamedMembers(name string, opts ...func(*AdminQuery)) *DepartmentQuery {
	query := (&AdminClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedMembers == nil {
		dq.withNamedMembers = make(map[string]*AdminQuery)
	}
	dq.withNamedMembers[name] = query
	return dq
}

// DepartmentGroupBy is the group-by builder for Department entities.
type DepartmentGroupBy struct {
	selector
	build *DepartmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DepartmentGroupBy) Aggregate(fns ...AggregateFunc) *DepartmentGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DepartmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DepartmentQuery, *DepartmentGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DepartmentGroupBy) sqlScan(ctx context.Context, root *DepartmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DepartmentSelect is the builder for selecting fields of Department entities.
type DepartmentSelect struct {
	*DepartmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DepartmentSelect) Aggregate(fns ...AggregateFunc) *DepartmentSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DepartmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DepartmentQuery, *DepartmentSelect](ctx, ds.DepartmentQuery, ds, ds.inters, v)
}

func (ds *DepartmentSelect) sqlScan(ctx context.Context, root *DepartmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ds *DepartmentSelect) Modify(modifiers ...func(s *sql.Selector)) *DepartmentSelect {
	ds.modifiers = append(ds.modifiers, modifiers...)
	return ds
}
