// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/predicate"
	"github.com/gva/internal/ent/role"

	"github.com/gva/internal/ent/internal"
)

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks     []Hook
	mutation  *AdminMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AdminUpdate builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AdminUpdate) SetCreatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableCreatedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AdminUpdate) SetUpdatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetIsEnable sets the "is_enable" field.
func (au *AdminUpdate) SetIsEnable(b bool) *AdminUpdate {
	au.mutation.SetIsEnable(b)
	return au
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (au *AdminUpdate) SetNillableIsEnable(b *bool) *AdminUpdate {
	if b != nil {
		au.SetIsEnable(*b)
	}
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AdminUpdate) SetDeletedAt(i int) *AdminUpdate {
	au.mutation.ResetDeletedAt()
	au.mutation.SetDeletedAt(i)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableDeletedAt(i *int) *AdminUpdate {
	if i != nil {
		au.SetDeletedAt(*i)
	}
	return au
}

// AddDeletedAt adds i to the "deleted_at" field.
func (au *AdminUpdate) AddDeletedAt(i int) *AdminUpdate {
	au.mutation.AddDeletedAt(i)
	return au
}

// SetUsername sets the "username" field.
func (au *AdminUpdate) SetUsername(s string) *AdminUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *AdminUpdate) SetNillableUsername(s *string) *AdminUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// SetPassword sets the "password" field.
func (au *AdminUpdate) SetPassword(s string) *AdminUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AdminUpdate) SetNillablePassword(s *string) *AdminUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// SetWhitelistIps sets the "whitelist_ips" field.
func (au *AdminUpdate) SetWhitelistIps(s []string) *AdminUpdate {
	au.mutation.SetWhitelistIps(s)
	return au
}

// AppendWhitelistIps appends s to the "whitelist_ips" field.
func (au *AdminUpdate) AppendWhitelistIps(s []string) *AdminUpdate {
	au.mutation.AppendWhitelistIps(s)
	return au
}

// SetDisplayName sets the "display_name" field.
func (au *AdminUpdate) SetDisplayName(s string) *AdminUpdate {
	au.mutation.SetDisplayName(s)
	return au
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (au *AdminUpdate) SetNillableDisplayName(s *string) *AdminUpdate {
	if s != nil {
		au.SetDisplayName(*s)
	}
	return au
}

// ClearDisplayName clears the value of the "display_name" field.
func (au *AdminUpdate) ClearDisplayName() *AdminUpdate {
	au.mutation.ClearDisplayName()
	return au
}

// SetDepartmentID sets the "department_id" field.
func (au *AdminUpdate) SetDepartmentID(x xid.ID) *AdminUpdate {
	au.mutation.SetDepartmentID(x)
	return au
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (au *AdminUpdate) SetNillableDepartmentID(x *xid.ID) *AdminUpdate {
	if x != nil {
		au.SetDepartmentID(*x)
	}
	return au
}

// ClearDepartmentID clears the value of the "department_id" field.
func (au *AdminUpdate) ClearDepartmentID() *AdminUpdate {
	au.mutation.ClearDepartmentID()
	return au
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (au *AdminUpdate) AddRoleIDs(ids ...xid.ID) *AdminUpdate {
	au.mutation.AddRoleIDs(ids...)
	return au
}

// AddRoles adds the "roles" edges to the Role entity.
func (au *AdminUpdate) AddRoles(r ...*Role) *AdminUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.AddRoleIDs(ids...)
}

// SetDepartment sets the "department" edge to the Department entity.
func (au *AdminUpdate) SetDepartment(d *Department) *AdminUpdate {
	return au.SetDepartmentID(d.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (au *AdminUpdate) ClearRoles() *AdminUpdate {
	au.mutation.ClearRoles()
	return au
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (au *AdminUpdate) RemoveRoleIDs(ids ...xid.ID) *AdminUpdate {
	au.mutation.RemoveRoleIDs(ids...)
	return au
}

// RemoveRoles removes "roles" edges to Role entities.
func (au *AdminUpdate) RemoveRoles(r ...*Role) *AdminUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.RemoveRoleIDs(ids...)
}

// ClearDepartment clears the "department" edge to the Department entity.
func (au *AdminUpdate) ClearDepartment() *AdminUpdate {
	au.mutation.ClearDepartment()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	if err := au.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AdminUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if admin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized admin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := admin.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AdminUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdminUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.IsEnable(); ok {
		_spec.SetField(admin.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedDeletedAt(); ok {
		_spec.AddField(admin.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.WhitelistIps(); ok {
		_spec.SetField(admin.FieldWhitelistIps, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedWhitelistIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, admin.FieldWhitelistIps, value)
		})
	}
	if value, ok := au.mutation.DisplayName(); ok {
		_spec.SetField(admin.FieldDisplayName, field.TypeString, value)
	}
	if au.mutation.DisplayNameCleared() {
		_spec.ClearField(admin.FieldDisplayName, field.TypeString)
	}
	if au.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = au.schemaConfig.AdminRoles
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedRolesIDs(); len(nodes) > 0 && !au.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = au.schemaConfig.AdminRoles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = au.schemaConfig.AdminRoles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.DepartmentTable,
			Columns: []string{admin.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		edge.Schema = au.schemaConfig.Admin
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.DepartmentTable,
			Columns: []string{admin.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		edge.Schema = au.schemaConfig.Admin
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = au.schemaConfig.Admin
	ctx = internal.NewSchemaConfigContext(ctx, au.schemaConfig)
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AdminMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (auo *AdminUpdateOne) SetCreatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableCreatedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AdminUpdateOne) SetUpdatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetIsEnable sets the "is_enable" field.
func (auo *AdminUpdateOne) SetIsEnable(b bool) *AdminUpdateOne {
	auo.mutation.SetIsEnable(b)
	return auo
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableIsEnable(b *bool) *AdminUpdateOne {
	if b != nil {
		auo.SetIsEnable(*b)
	}
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AdminUpdateOne) SetDeletedAt(i int) *AdminUpdateOne {
	auo.mutation.ResetDeletedAt()
	auo.mutation.SetDeletedAt(i)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableDeletedAt(i *int) *AdminUpdateOne {
	if i != nil {
		auo.SetDeletedAt(*i)
	}
	return auo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (auo *AdminUpdateOne) AddDeletedAt(i int) *AdminUpdateOne {
	auo.mutation.AddDeletedAt(i)
	return auo
}

// SetUsername sets the "username" field.
func (auo *AdminUpdateOne) SetUsername(s string) *AdminUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableUsername(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// SetPassword sets the "password" field.
func (auo *AdminUpdateOne) SetPassword(s string) *AdminUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillablePassword(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// SetWhitelistIps sets the "whitelist_ips" field.
func (auo *AdminUpdateOne) SetWhitelistIps(s []string) *AdminUpdateOne {
	auo.mutation.SetWhitelistIps(s)
	return auo
}

// AppendWhitelistIps appends s to the "whitelist_ips" field.
func (auo *AdminUpdateOne) AppendWhitelistIps(s []string) *AdminUpdateOne {
	auo.mutation.AppendWhitelistIps(s)
	return auo
}

// SetDisplayName sets the "display_name" field.
func (auo *AdminUpdateOne) SetDisplayName(s string) *AdminUpdateOne {
	auo.mutation.SetDisplayName(s)
	return auo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableDisplayName(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetDisplayName(*s)
	}
	return auo
}

// ClearDisplayName clears the value of the "display_name" field.
func (auo *AdminUpdateOne) ClearDisplayName() *AdminUpdateOne {
	auo.mutation.ClearDisplayName()
	return auo
}

// SetDepartmentID sets the "department_id" field.
func (auo *AdminUpdateOne) SetDepartmentID(x xid.ID) *AdminUpdateOne {
	auo.mutation.SetDepartmentID(x)
	return auo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableDepartmentID(x *xid.ID) *AdminUpdateOne {
	if x != nil {
		auo.SetDepartmentID(*x)
	}
	return auo
}

// ClearDepartmentID clears the value of the "department_id" field.
func (auo *AdminUpdateOne) ClearDepartmentID() *AdminUpdateOne {
	auo.mutation.ClearDepartmentID()
	return auo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (auo *AdminUpdateOne) AddRoleIDs(ids ...xid.ID) *AdminUpdateOne {
	auo.mutation.AddRoleIDs(ids...)
	return auo
}

// AddRoles adds the "roles" edges to the Role entity.
func (auo *AdminUpdateOne) AddRoles(r ...*Role) *AdminUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.AddRoleIDs(ids...)
}

// SetDepartment sets the "department" edge to the Department entity.
func (auo *AdminUpdateOne) SetDepartment(d *Department) *AdminUpdateOne {
	return auo.SetDepartmentID(d.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (auo *AdminUpdateOne) ClearRoles() *AdminUpdateOne {
	auo.mutation.ClearRoles()
	return auo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (auo *AdminUpdateOne) RemoveRoleIDs(ids ...xid.ID) *AdminUpdateOne {
	auo.mutation.RemoveRoleIDs(ids...)
	return auo
}

// RemoveRoles removes "roles" edges to Role entities.
func (auo *AdminUpdateOne) RemoveRoles(r ...*Role) *AdminUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.RemoveRoleIDs(ids...)
}

// ClearDepartment clears the "department" edge to the Department entity.
func (auo *AdminUpdateOne) ClearDepartment() *AdminUpdateOne {
	auo.mutation.ClearDepartment()
	return auo
}

// Where appends a list predicates to the AdminUpdate builder.
func (auo *AdminUpdateOne) Where(ps ...predicate.Admin) *AdminUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdminUpdateOne) Select(field string, fields ...string) *AdminUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Admin entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AdminUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if admin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized admin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := admin.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AdminUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdminUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	_spec := sqlgraph.NewUpdateSpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Admin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admin.FieldID)
		for _, f := range fields {
			if !admin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != admin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.IsEnable(); ok {
		_spec.SetField(admin.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(admin.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.WhitelistIps(); ok {
		_spec.SetField(admin.FieldWhitelistIps, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedWhitelistIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, admin.FieldWhitelistIps, value)
		})
	}
	if value, ok := auo.mutation.DisplayName(); ok {
		_spec.SetField(admin.FieldDisplayName, field.TypeString, value)
	}
	if auo.mutation.DisplayNameCleared() {
		_spec.ClearField(admin.FieldDisplayName, field.TypeString)
	}
	if auo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = auo.schemaConfig.AdminRoles
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !auo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = auo.schemaConfig.AdminRoles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = auo.schemaConfig.AdminRoles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.DepartmentTable,
			Columns: []string{admin.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		edge.Schema = auo.schemaConfig.Admin
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.DepartmentTable,
			Columns: []string{admin.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		edge.Schema = auo.schemaConfig.Admin
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = auo.schemaConfig.Admin
	ctx = internal.NewSchemaConfigContext(ctx, auo.schemaConfig)
	_spec.AddModifiers(auo.modifiers...)
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
