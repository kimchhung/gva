// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/extra/internal/ent/admin"
	"github.com/kimchhung/gva/extra/internal/ent/permission"
	"github.com/kimchhung/gva/extra/internal/ent/predicate"
	"github.com/kimchhung/gva/extra/internal/ent/role"
	"github.com/kimchhung/gva/extra/internal/ent/route"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RoleUpdate) SetCreatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCreatedAt(t *time.Time) *RoleUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetIsEnable sets the "is_enable" field.
func (ru *RoleUpdate) SetIsEnable(b bool) *RoleUpdate {
	ru.mutation.SetIsEnable(b)
	return ru
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableIsEnable(b *bool) *RoleUpdate {
	if b != nil {
		ru.SetIsEnable(*b)
	}
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RoleUpdate) SetDeletedAt(i int) *RoleUpdate {
	ru.mutation.ResetDeletedAt()
	ru.mutation.SetDeletedAt(i)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDeletedAt(i *int) *RoleUpdate {
	if i != nil {
		ru.SetDeletedAt(*i)
	}
	return ru
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ru *RoleUpdate) AddDeletedAt(i int) *RoleUpdate {
	ru.mutation.AddDeletedAt(i)
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *RoleUpdate) SetDescription(s string) *RoleUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDescription(s *string) *RoleUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetOrder sets the "order" field.
func (ru *RoleUpdate) SetOrder(i int) *RoleUpdate {
	ru.mutation.ResetOrder()
	ru.mutation.SetOrder(i)
	return ru
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableOrder(i *int) *RoleUpdate {
	if i != nil {
		ru.SetOrder(*i)
	}
	return ru
}

// AddOrder adds i to the "order" field.
func (ru *RoleUpdate) AddOrder(i int) *RoleUpdate {
	ru.mutation.AddOrder(i)
	return ru
}

// SetIsChangeable sets the "is_changeable" field.
func (ru *RoleUpdate) SetIsChangeable(b bool) *RoleUpdate {
	ru.mutation.SetIsChangeable(b)
	return ru
}

// SetNillableIsChangeable sets the "is_changeable" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableIsChangeable(b *bool) *RoleUpdate {
	if b != nil {
		ru.SetIsChangeable(*b)
	}
	return ru
}

// AddAdminIDs adds the "admins" edge to the Admin entity by IDs.
func (ru *RoleUpdate) AddAdminIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddAdminIDs(ids...)
	return ru
}

// AddAdmins adds the "admins" edges to the Admin entity.
func (ru *RoleUpdate) AddAdmins(a ...*Admin) *RoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAdminIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ru *RoleUpdate) AddPermissionIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddPermissionIDs(ids...)
	return ru
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ru *RoleUpdate) AddPermissions(p ...*Permission) *RoleUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPermissionIDs(ids...)
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (ru *RoleUpdate) AddRouteIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddRouteIDs(ids...)
	return ru
}

// AddRoutes adds the "routes" edges to the Route entity.
func (ru *RoleUpdate) AddRoutes(r ...*Route) *RoleUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRouteIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearAdmins clears all "admins" edges to the Admin entity.
func (ru *RoleUpdate) ClearAdmins() *RoleUpdate {
	ru.mutation.ClearAdmins()
	return ru
}

// RemoveAdminIDs removes the "admins" edge to Admin entities by IDs.
func (ru *RoleUpdate) RemoveAdminIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemoveAdminIDs(ids...)
	return ru
}

// RemoveAdmins removes "admins" edges to Admin entities.
func (ru *RoleUpdate) RemoveAdmins(a ...*Admin) *RoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAdminIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (ru *RoleUpdate) ClearPermissions() *RoleUpdate {
	ru.mutation.ClearPermissions()
	return ru
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (ru *RoleUpdate) RemovePermissionIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemovePermissionIDs(ids...)
	return ru
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (ru *RoleUpdate) RemovePermissions(p ...*Permission) *RoleUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePermissionIDs(ids...)
}

// ClearRoutes clears all "routes" edges to the Route entity.
func (ru *RoleUpdate) ClearRoutes() *RoleUpdate {
	ru.mutation.ClearRoutes()
	return ru
}

// RemoveRouteIDs removes the "routes" edge to Route entities by IDs.
func (ru *RoleUpdate) RemoveRouteIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemoveRouteIDs(ids...)
	return ru
}

// RemoveRoutes removes "routes" edges to Route entities.
func (ru *RoleUpdate) RemoveRoutes(r ...*Route) *RoleUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRouteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if role.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized role.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := role.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.IsEnable(); ok {
		_spec.SetField(role.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedDeletedAt(); ok {
		_spec.AddField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.Order(); ok {
		_spec.SetField(role.FieldOrder, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedOrder(); ok {
		_spec.AddField(role.FieldOrder, field.TypeInt, value)
	}
	if value, ok := ru.mutation.IsChangeable(); ok {
		_spec.SetField(role.FieldIsChangeable, field.TypeBool, value)
	}
	if ru.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAdminsIDs(); len(nodes) > 0 && !ru.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !ru.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRoutesIDs(); len(nodes) > 0 && !ru.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RoleUpdateOne) SetCreatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCreatedAt(t *time.Time) *RoleUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetIsEnable sets the "is_enable" field.
func (ruo *RoleUpdateOne) SetIsEnable(b bool) *RoleUpdateOne {
	ruo.mutation.SetIsEnable(b)
	return ruo
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableIsEnable(b *bool) *RoleUpdateOne {
	if b != nil {
		ruo.SetIsEnable(*b)
	}
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RoleUpdateOne) SetDeletedAt(i int) *RoleUpdateOne {
	ruo.mutation.ResetDeletedAt()
	ruo.mutation.SetDeletedAt(i)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDeletedAt(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetDeletedAt(*i)
	}
	return ruo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ruo *RoleUpdateOne) AddDeletedAt(i int) *RoleUpdateOne {
	ruo.mutation.AddDeletedAt(i)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RoleUpdateOne) SetDescription(s string) *RoleUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDescription(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetOrder sets the "order" field.
func (ruo *RoleUpdateOne) SetOrder(i int) *RoleUpdateOne {
	ruo.mutation.ResetOrder()
	ruo.mutation.SetOrder(i)
	return ruo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableOrder(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetOrder(*i)
	}
	return ruo
}

// AddOrder adds i to the "order" field.
func (ruo *RoleUpdateOne) AddOrder(i int) *RoleUpdateOne {
	ruo.mutation.AddOrder(i)
	return ruo
}

// SetIsChangeable sets the "is_changeable" field.
func (ruo *RoleUpdateOne) SetIsChangeable(b bool) *RoleUpdateOne {
	ruo.mutation.SetIsChangeable(b)
	return ruo
}

// SetNillableIsChangeable sets the "is_changeable" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableIsChangeable(b *bool) *RoleUpdateOne {
	if b != nil {
		ruo.SetIsChangeable(*b)
	}
	return ruo
}

// AddAdminIDs adds the "admins" edge to the Admin entity by IDs.
func (ruo *RoleUpdateOne) AddAdminIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddAdminIDs(ids...)
	return ruo
}

// AddAdmins adds the "admins" edges to the Admin entity.
func (ruo *RoleUpdateOne) AddAdmins(a ...*Admin) *RoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAdminIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ruo *RoleUpdateOne) AddPermissionIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddPermissionIDs(ids...)
	return ruo
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ruo *RoleUpdateOne) AddPermissions(p ...*Permission) *RoleUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPermissionIDs(ids...)
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (ruo *RoleUpdateOne) AddRouteIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddRouteIDs(ids...)
	return ruo
}

// AddRoutes adds the "routes" edges to the Route entity.
func (ruo *RoleUpdateOne) AddRoutes(r ...*Route) *RoleUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRouteIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearAdmins clears all "admins" edges to the Admin entity.
func (ruo *RoleUpdateOne) ClearAdmins() *RoleUpdateOne {
	ruo.mutation.ClearAdmins()
	return ruo
}

// RemoveAdminIDs removes the "admins" edge to Admin entities by IDs.
func (ruo *RoleUpdateOne) RemoveAdminIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemoveAdminIDs(ids...)
	return ruo
}

// RemoveAdmins removes "admins" edges to Admin entities.
func (ruo *RoleUpdateOne) RemoveAdmins(a ...*Admin) *RoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAdminIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (ruo *RoleUpdateOne) ClearPermissions() *RoleUpdateOne {
	ruo.mutation.ClearPermissions()
	return ruo
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (ruo *RoleUpdateOne) RemovePermissionIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemovePermissionIDs(ids...)
	return ruo
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (ruo *RoleUpdateOne) RemovePermissions(p ...*Permission) *RoleUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePermissionIDs(ids...)
}

// ClearRoutes clears all "routes" edges to the Route entity.
func (ruo *RoleUpdateOne) ClearRoutes() *RoleUpdateOne {
	ruo.mutation.ClearRoutes()
	return ruo
}

// RemoveRouteIDs removes the "routes" edge to Route entities by IDs.
func (ruo *RoleUpdateOne) RemoveRouteIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemoveRouteIDs(ids...)
	return ruo
}

// RemoveRoutes removes "routes" edges to Route entities.
func (ruo *RoleUpdateOne) RemoveRoutes(r ...*Route) *RoleUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRouteIDs(ids...)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if role.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized role.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := role.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.IsEnable(); ok {
		_spec.SetField(role.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Order(); ok {
		_spec.SetField(role.FieldOrder, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedOrder(); ok {
		_spec.AddField(role.FieldOrder, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.IsChangeable(); ok {
		_spec.SetField(role.FieldIsChangeable, field.TypeBool, value)
	}
	if ruo.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAdminsIDs(); len(nodes) > 0 && !ruo.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !ruo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRoutesIDs(); len(nodes) > 0 && !ruo.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
