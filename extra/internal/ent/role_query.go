// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
	"github.com/kimchhung/gva/extra/internal/ent/permission"
	"github.com/kimchhung/gva/extra/internal/ent/predicate"
	"github.com/kimchhung/gva/extra/internal/ent/role"
	"github.com/kimchhung/gva/extra/internal/ent/route"
)

// RoleQuery is the builder for querying Role entities.
type RoleQuery struct {
	config
	ctx             *QueryContext
	order           []role.OrderOption
	inters          []Interceptor
	predicates      []predicate.Role
	withAdmins      *AdminQuery
	withPermissions *PermissionQuery
	withRoutes      *RouteQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleQuery builder.
func (rq *RoleQuery) Where(ps ...predicate.Role) *RoleQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RoleQuery) Limit(limit int) *RoleQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RoleQuery) Offset(offset int) *RoleQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoleQuery) Unique(unique bool) *RoleQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RoleQuery) Order(o ...role.OrderOption) *RoleQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryAdmins chains the current query on the "admins" edge.
func (rq *RoleQuery) QueryAdmins() *AdminQuery {
	query := (&AdminClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.AdminsTable, role.AdminsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermissions chains the current query on the "permissions" edge.
func (rq *RoleQuery) QueryPermissions() *PermissionQuery {
	query := (&PermissionClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.PermissionsTable, role.PermissionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutes chains the current query on the "routes" edge.
func (rq *RoleQuery) QueryRoutes() *RouteQuery {
	query := (&RouteClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(route.Table, route.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.RoutesTable, role.RoutesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Role entity from the query.
// Returns a *NotFoundError when no Role was found.
func (rq *RoleQuery) First(ctx context.Context) (*Role, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{role.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoleQuery) FirstX(ctx context.Context) *Role {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Role ID from the query.
// Returns a *NotFoundError when no Role ID was found.
func (rq *RoleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{role.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoleQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Role entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Role entity is found.
// Returns a *NotFoundError when no Role entities are found.
func (rq *RoleQuery) Only(ctx context.Context) (*Role, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{role.Label}
	default:
		return nil, &NotSingularError{role.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoleQuery) OnlyX(ctx context.Context) *Role {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Role ID in the query.
// Returns a *NotSingularError when more than one Role ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{role.Label}
	default:
		err = &NotSingularError{role.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoleQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Roles.
func (rq *RoleQuery) All(ctx context.Context) ([]*Role, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Role, *RoleQuery]()
	return withInterceptors[[]*Role](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoleQuery) AllX(ctx context.Context) []*Role {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Role IDs.
func (rq *RoleQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(role.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoleQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RoleQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoleQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoleQuery) Exist(ctx context.Context) (bool, error) {
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
func (rq *RoleQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoleQuery) Clone() *RoleQuery {
	if rq == nil {
		return nil
	}
	return &RoleQuery{
		config:          rq.config,
		ctx:             rq.ctx.Clone(),
		order:           append([]role.OrderOption{}, rq.order...),
		inters:          append([]Interceptor{}, rq.inters...),
		predicates:      append([]predicate.Role{}, rq.predicates...),
		withAdmins:      rq.withAdmins.Clone(),
		withPermissions: rq.withPermissions.Clone(),
		withRoutes:      rq.withRoutes.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithAdmins tells the query-builder to eager-load the nodes that are connected to
// the "admins" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithAdmins(opts ...func(*AdminQuery)) *RoleQuery {
	query := (&AdminClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withAdmins = query
	return rq
}

// WithPermissions tells the query-builder to eager-load the nodes that are connected to
// the "permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithPermissions(opts ...func(*PermissionQuery)) *RoleQuery {
	query := (&PermissionClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withPermissions = query
	return rq
}

// WithRoutes tells the query-builder to eager-load the nodes that are connected to
// the "routes" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithRoutes(opts ...func(*RouteQuery)) *RoleQuery {
	query := (&RouteClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withRoutes = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" rql:"name=createdAt,column=created_at,filter,sort"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Role.Query().
//		GroupBy(role.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RoleQuery) GroupBy(field string, fields ...string) *RoleGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoleGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = role.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" rql:"name=createdAt,column=created_at,filter,sort"`
//	}
//
//	client.Role.Query().
//		Select(role.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *RoleQuery) Select(fields ...string) *RoleSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RoleSelect{RoleQuery: rq}
	sbuild.label = role.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoleSelect configured with the given aggregations.
func (rq *RoleQuery) Aggregate(fns ...AggregateFunc) *RoleSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoleQuery) prepareQuery(ctx context.Context) error {
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
		if !role.ValidColumn(f) {
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

func (rq *RoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Role, error) {
	var (
		nodes       = []*Role{}
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withAdmins != nil,
			rq.withPermissions != nil,
			rq.withRoutes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Role).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Role{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := rq.withAdmins; query != nil {
		if err := rq.loadAdmins(ctx, query, nodes,
			func(n *Role) { n.Edges.Admins = []*Admin{} },
			func(n *Role, e *Admin) { n.Edges.Admins = append(n.Edges.Admins, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withPermissions; query != nil {
		if err := rq.loadPermissions(ctx, query, nodes,
			func(n *Role) { n.Edges.Permissions = []*Permission{} },
			func(n *Role, e *Permission) { n.Edges.Permissions = append(n.Edges.Permissions, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withRoutes; query != nil {
		if err := rq.loadRoutes(ctx, query, nodes,
			func(n *Role) { n.Edges.Routes = []*Route{} },
			func(n *Role, e *Route) { n.Edges.Routes = append(n.Edges.Routes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RoleQuery) loadAdmins(ctx context.Context, query *AdminQuery, nodes []*Role, init func(*Role), assign func(*Role, *Admin)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Role)
	nids := make(map[int]map[*Role]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(role.AdminsTable)
		s.Join(joinT).On(s.C(admin.FieldID), joinT.C(role.AdminsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(role.AdminsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(role.AdminsPrimaryKey[1]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Role]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Admin](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "admins" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RoleQuery) loadPermissions(ctx context.Context, query *PermissionQuery, nodes []*Role, init func(*Role), assign func(*Role, *Permission)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Role)
	nids := make(map[int]map[*Role]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(role.PermissionsTable)
		s.Join(joinT).On(s.C(permission.FieldID), joinT.C(role.PermissionsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(role.PermissionsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(role.PermissionsPrimaryKey[0]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Role]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Permission](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "permissions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RoleQuery) loadRoutes(ctx context.Context, query *RouteQuery, nodes []*Role, init func(*Role), assign func(*Role, *Route)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Role)
	nids := make(map[int]map[*Role]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(role.RoutesTable)
		s.Join(joinT).On(s.C(route.FieldID), joinT.C(role.RoutesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(role.RoutesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(role.RoutesPrimaryKey[0]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Role]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Route](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "routes" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (rq *RoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for i := range fields {
			if fields[i] != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
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

func (rq *RoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(role.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = role.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
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

// RoleGroupBy is the group-by builder for Role entities.
type RoleGroupBy struct {
	selector
	build *RoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoleGroupBy) Aggregate(fns ...AggregateFunc) *RoleGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleQuery, *RoleGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RoleGroupBy) sqlScan(ctx context.Context, root *RoleQuery, v any) error {
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

// RoleSelect is the builder for selecting fields of Role entities.
type RoleSelect struct {
	*RoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoleSelect) Aggregate(fns ...AggregateFunc) *RoleSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleQuery, *RoleSelect](ctx, rs.RoleQuery, rs, rs.inters, v)
}

func (rs *RoleSelect) sqlScan(ctx context.Context, root *RoleQuery, v any) error {
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
