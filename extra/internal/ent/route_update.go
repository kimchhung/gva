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
	"github.com/kimchhung/gva/extra/app/database/schema/types"
	"github.com/kimchhung/gva/extra/internal/ent/predicate"
	"github.com/kimchhung/gva/extra/internal/ent/role"
	"github.com/kimchhung/gva/extra/internal/ent/route"
)

// RouteUpdate is the builder for updating Route entities.
type RouteUpdate struct {
	config
	hooks    []Hook
	mutation *RouteMutation
}

// Where appends a list predicates to the RouteUpdate builder.
func (ru *RouteUpdate) Where(ps ...predicate.Route) *RouteUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RouteUpdate) SetCreatedAt(t time.Time) *RouteUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableCreatedAt(t *time.Time) *RouteUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RouteUpdate) SetUpdatedAt(t time.Time) *RouteUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetIsEnable sets the "is_enable" field.
func (ru *RouteUpdate) SetIsEnable(b bool) *RouteUpdate {
	ru.mutation.SetIsEnable(b)
	return ru
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableIsEnable(b *bool) *RouteUpdate {
	if b != nil {
		ru.SetIsEnable(*b)
	}
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RouteUpdate) SetDeletedAt(i int) *RouteUpdate {
	ru.mutation.ResetDeletedAt()
	ru.mutation.SetDeletedAt(i)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableDeletedAt(i *int) *RouteUpdate {
	if i != nil {
		ru.SetDeletedAt(*i)
	}
	return ru
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ru *RouteUpdate) AddDeletedAt(i int) *RouteUpdate {
	ru.mutation.AddDeletedAt(i)
	return ru
}

// SetParentID sets the "parent_id" field.
func (ru *RouteUpdate) SetParentID(i int) *RouteUpdate {
	ru.mutation.SetParentID(i)
	return ru
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableParentID(i *int) *RouteUpdate {
	if i != nil {
		ru.SetParentID(*i)
	}
	return ru
}

// ClearParentID clears the value of the "parent_id" field.
func (ru *RouteUpdate) ClearParentID() *RouteUpdate {
	ru.mutation.ClearParentID()
	return ru
}

// SetPath sets the "path" field.
func (ru *RouteUpdate) SetPath(s string) *RouteUpdate {
	ru.mutation.SetPath(s)
	return ru
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (ru *RouteUpdate) SetNillablePath(s *string) *RouteUpdate {
	if s != nil {
		ru.SetPath(*s)
	}
	return ru
}

// SetComponent sets the "component" field.
func (ru *RouteUpdate) SetComponent(s string) *RouteUpdate {
	ru.mutation.SetComponent(s)
	return ru
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableComponent(s *string) *RouteUpdate {
	if s != nil {
		ru.SetComponent(*s)
	}
	return ru
}

// SetRedirect sets the "redirect" field.
func (ru *RouteUpdate) SetRedirect(s string) *RouteUpdate {
	ru.mutation.SetRedirect(s)
	return ru
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableRedirect(s *string) *RouteUpdate {
	if s != nil {
		ru.SetRedirect(*s)
	}
	return ru
}

// ClearRedirect clears the value of the "redirect" field.
func (ru *RouteUpdate) ClearRedirect() *RouteUpdate {
	ru.mutation.ClearRedirect()
	return ru
}

// SetName sets the "name" field.
func (ru *RouteUpdate) SetName(s string) *RouteUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableName(s *string) *RouteUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetType sets the "type" field.
func (ru *RouteUpdate) SetType(r route.Type) *RouteUpdate {
	ru.mutation.SetType(r)
	return ru
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableType(r *route.Type) *RouteUpdate {
	if r != nil {
		ru.SetType(*r)
	}
	return ru
}

// SetMeta sets the "meta" field.
func (ru *RouteUpdate) SetMeta(tm types.RouteMeta) *RouteUpdate {
	ru.mutation.SetMeta(tm)
	return ru
}

// SetNillableMeta sets the "meta" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableMeta(tm *types.RouteMeta) *RouteUpdate {
	if tm != nil {
		ru.SetMeta(*tm)
	}
	return ru
}

// SetParent sets the "parent" edge to the Route entity.
func (ru *RouteUpdate) SetParent(r *Route) *RouteUpdate {
	return ru.SetParentID(r.ID)
}

// AddChildIDs adds the "children" edge to the Route entity by IDs.
func (ru *RouteUpdate) AddChildIDs(ids ...int) *RouteUpdate {
	ru.mutation.AddChildIDs(ids...)
	return ru
}

// AddChildren adds the "children" edges to the Route entity.
func (ru *RouteUpdate) AddChildren(r ...*Route) *RouteUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddChildIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (ru *RouteUpdate) AddRoleIDs(ids ...int) *RouteUpdate {
	ru.mutation.AddRoleIDs(ids...)
	return ru
}

// AddRoles adds the "roles" edges to the Role entity.
func (ru *RouteUpdate) AddRoles(r ...*Role) *RouteUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRoleIDs(ids...)
}

// Mutation returns the RouteMutation object of the builder.
func (ru *RouteUpdate) Mutation() *RouteMutation {
	return ru.mutation
}

// ClearParent clears the "parent" edge to the Route entity.
func (ru *RouteUpdate) ClearParent() *RouteUpdate {
	ru.mutation.ClearParent()
	return ru
}

// ClearChildren clears all "children" edges to the Route entity.
func (ru *RouteUpdate) ClearChildren() *RouteUpdate {
	ru.mutation.ClearChildren()
	return ru
}

// RemoveChildIDs removes the "children" edge to Route entities by IDs.
func (ru *RouteUpdate) RemoveChildIDs(ids ...int) *RouteUpdate {
	ru.mutation.RemoveChildIDs(ids...)
	return ru
}

// RemoveChildren removes "children" edges to Route entities.
func (ru *RouteUpdate) RemoveChildren(r ...*Route) *RouteUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveChildIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role entity.
func (ru *RouteUpdate) ClearRoles() *RouteUpdate {
	ru.mutation.ClearRoles()
	return ru
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (ru *RouteUpdate) RemoveRoleIDs(ids ...int) *RouteUpdate {
	ru.mutation.RemoveRoleIDs(ids...)
	return ru
}

// RemoveRoles removes "roles" edges to Role entities.
func (ru *RouteUpdate) RemoveRoles(r ...*Role) *RouteUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RouteUpdate) Save(ctx context.Context) (int, error) {
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RouteUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RouteUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RouteUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RouteUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if route.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized route.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := route.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ru *RouteUpdate) check() error {
	if v, ok := ru.mutation.GetType(); ok {
		if err := route.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Route.type": %w`, err)}
		}
	}
	return nil
}

func (ru *RouteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(route.Table, route.Columns, sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(route.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(route.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.IsEnable(); ok {
		_spec.SetField(route.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(route.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedDeletedAt(); ok {
		_spec.AddField(route.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Path(); ok {
		_spec.SetField(route.FieldPath, field.TypeString, value)
	}
	if value, ok := ru.mutation.Component(); ok {
		_spec.SetField(route.FieldComponent, field.TypeString, value)
	}
	if value, ok := ru.mutation.Redirect(); ok {
		_spec.SetField(route.FieldRedirect, field.TypeString, value)
	}
	if ru.mutation.RedirectCleared() {
		_spec.ClearField(route.FieldRedirect, field.TypeString)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(route.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.GetType(); ok {
		_spec.SetField(route.FieldType, field.TypeEnum, value)
	}
	if value, ok := ru.mutation.Meta(); ok {
		_spec.SetField(route.FieldMeta, field.TypeJSON, value)
	}
	if ru.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.ParentTable,
			Columns: []string{route.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.ParentTable,
			Columns: []string{route.ParentColumn},
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
	if ru.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ru.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
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
	if nodes := ru.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
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
	if ru.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRolesIDs(); len(nodes) > 0 && !ru.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{route.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RouteUpdateOne is the builder for updating a single Route entity.
type RouteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RouteMutation
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RouteUpdateOne) SetCreatedAt(t time.Time) *RouteUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableCreatedAt(t *time.Time) *RouteUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RouteUpdateOne) SetUpdatedAt(t time.Time) *RouteUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetIsEnable sets the "is_enable" field.
func (ruo *RouteUpdateOne) SetIsEnable(b bool) *RouteUpdateOne {
	ruo.mutation.SetIsEnable(b)
	return ruo
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableIsEnable(b *bool) *RouteUpdateOne {
	if b != nil {
		ruo.SetIsEnable(*b)
	}
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RouteUpdateOne) SetDeletedAt(i int) *RouteUpdateOne {
	ruo.mutation.ResetDeletedAt()
	ruo.mutation.SetDeletedAt(i)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableDeletedAt(i *int) *RouteUpdateOne {
	if i != nil {
		ruo.SetDeletedAt(*i)
	}
	return ruo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ruo *RouteUpdateOne) AddDeletedAt(i int) *RouteUpdateOne {
	ruo.mutation.AddDeletedAt(i)
	return ruo
}

// SetParentID sets the "parent_id" field.
func (ruo *RouteUpdateOne) SetParentID(i int) *RouteUpdateOne {
	ruo.mutation.SetParentID(i)
	return ruo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableParentID(i *int) *RouteUpdateOne {
	if i != nil {
		ruo.SetParentID(*i)
	}
	return ruo
}

// ClearParentID clears the value of the "parent_id" field.
func (ruo *RouteUpdateOne) ClearParentID() *RouteUpdateOne {
	ruo.mutation.ClearParentID()
	return ruo
}

// SetPath sets the "path" field.
func (ruo *RouteUpdateOne) SetPath(s string) *RouteUpdateOne {
	ruo.mutation.SetPath(s)
	return ruo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillablePath(s *string) *RouteUpdateOne {
	if s != nil {
		ruo.SetPath(*s)
	}
	return ruo
}

// SetComponent sets the "component" field.
func (ruo *RouteUpdateOne) SetComponent(s string) *RouteUpdateOne {
	ruo.mutation.SetComponent(s)
	return ruo
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableComponent(s *string) *RouteUpdateOne {
	if s != nil {
		ruo.SetComponent(*s)
	}
	return ruo
}

// SetRedirect sets the "redirect" field.
func (ruo *RouteUpdateOne) SetRedirect(s string) *RouteUpdateOne {
	ruo.mutation.SetRedirect(s)
	return ruo
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableRedirect(s *string) *RouteUpdateOne {
	if s != nil {
		ruo.SetRedirect(*s)
	}
	return ruo
}

// ClearRedirect clears the value of the "redirect" field.
func (ruo *RouteUpdateOne) ClearRedirect() *RouteUpdateOne {
	ruo.mutation.ClearRedirect()
	return ruo
}

// SetName sets the "name" field.
func (ruo *RouteUpdateOne) SetName(s string) *RouteUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableName(s *string) *RouteUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetType sets the "type" field.
func (ruo *RouteUpdateOne) SetType(r route.Type) *RouteUpdateOne {
	ruo.mutation.SetType(r)
	return ruo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableType(r *route.Type) *RouteUpdateOne {
	if r != nil {
		ruo.SetType(*r)
	}
	return ruo
}

// SetMeta sets the "meta" field.
func (ruo *RouteUpdateOne) SetMeta(tm types.RouteMeta) *RouteUpdateOne {
	ruo.mutation.SetMeta(tm)
	return ruo
}

// SetNillableMeta sets the "meta" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableMeta(tm *types.RouteMeta) *RouteUpdateOne {
	if tm != nil {
		ruo.SetMeta(*tm)
	}
	return ruo
}

// SetParent sets the "parent" edge to the Route entity.
func (ruo *RouteUpdateOne) SetParent(r *Route) *RouteUpdateOne {
	return ruo.SetParentID(r.ID)
}

// AddChildIDs adds the "children" edge to the Route entity by IDs.
func (ruo *RouteUpdateOne) AddChildIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.AddChildIDs(ids...)
	return ruo
}

// AddChildren adds the "children" edges to the Route entity.
func (ruo *RouteUpdateOne) AddChildren(r ...*Route) *RouteUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddChildIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (ruo *RouteUpdateOne) AddRoleIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.AddRoleIDs(ids...)
	return ruo
}

// AddRoles adds the "roles" edges to the Role entity.
func (ruo *RouteUpdateOne) AddRoles(r ...*Role) *RouteUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRoleIDs(ids...)
}

// Mutation returns the RouteMutation object of the builder.
func (ruo *RouteUpdateOne) Mutation() *RouteMutation {
	return ruo.mutation
}

// ClearParent clears the "parent" edge to the Route entity.
func (ruo *RouteUpdateOne) ClearParent() *RouteUpdateOne {
	ruo.mutation.ClearParent()
	return ruo
}

// ClearChildren clears all "children" edges to the Route entity.
func (ruo *RouteUpdateOne) ClearChildren() *RouteUpdateOne {
	ruo.mutation.ClearChildren()
	return ruo
}

// RemoveChildIDs removes the "children" edge to Route entities by IDs.
func (ruo *RouteUpdateOne) RemoveChildIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.RemoveChildIDs(ids...)
	return ruo
}

// RemoveChildren removes "children" edges to Route entities.
func (ruo *RouteUpdateOne) RemoveChildren(r ...*Route) *RouteUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveChildIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role entity.
func (ruo *RouteUpdateOne) ClearRoles() *RouteUpdateOne {
	ruo.mutation.ClearRoles()
	return ruo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (ruo *RouteUpdateOne) RemoveRoleIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.RemoveRoleIDs(ids...)
	return ruo
}

// RemoveRoles removes "roles" edges to Role entities.
func (ruo *RouteUpdateOne) RemoveRoles(r ...*Role) *RouteUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the RouteUpdate builder.
func (ruo *RouteUpdateOne) Where(ps ...predicate.Route) *RouteUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RouteUpdateOne) Select(field string, fields ...string) *RouteUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Route entity.
func (ruo *RouteUpdateOne) Save(ctx context.Context) (*Route, error) {
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RouteUpdateOne) SaveX(ctx context.Context) *Route {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RouteUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RouteUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RouteUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if route.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized route.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := route.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RouteUpdateOne) check() error {
	if v, ok := ruo.mutation.GetType(); ok {
		if err := route.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Route.type": %w`, err)}
		}
	}
	return nil
}

func (ruo *RouteUpdateOne) sqlSave(ctx context.Context) (_node *Route, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(route.Table, route.Columns, sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Route.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, route.FieldID)
		for _, f := range fields {
			if !route.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != route.FieldID {
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
		_spec.SetField(route.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(route.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.IsEnable(); ok {
		_spec.SetField(route.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(route.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(route.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Path(); ok {
		_spec.SetField(route.FieldPath, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Component(); ok {
		_spec.SetField(route.FieldComponent, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Redirect(); ok {
		_spec.SetField(route.FieldRedirect, field.TypeString, value)
	}
	if ruo.mutation.RedirectCleared() {
		_spec.ClearField(route.FieldRedirect, field.TypeString)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(route.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.GetType(); ok {
		_spec.SetField(route.FieldType, field.TypeEnum, value)
	}
	if value, ok := ruo.mutation.Meta(); ok {
		_spec.SetField(route.FieldMeta, field.TypeJSON, value)
	}
	if ruo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.ParentTable,
			Columns: []string{route.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.ParentTable,
			Columns: []string{route.ParentColumn},
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
	if ruo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ruo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
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
	if nodes := ruo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
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
	if ruo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !ruo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Route{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{route.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
