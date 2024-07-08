// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/menu"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/role"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoleCreate) SetCreatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoleCreate) SetUpdatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetIsEnable sets the "is_enable" field.
func (rc *RoleCreate) SetIsEnable(b bool) *RoleCreate {
	rc.mutation.SetIsEnable(b)
	return rc
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (rc *RoleCreate) SetNillableIsEnable(b *bool) *RoleCreate {
	if b != nil {
		rc.SetIsEnable(*b)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *RoleCreate) SetDeletedAt(i int) *RoleCreate {
	rc.mutation.SetDeletedAt(i)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDeletedAt(i *int) *RoleCreate {
	if i != nil {
		rc.SetDeletedAt(*i)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoleCreate) SetDescription(s string) *RoleCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetOrder sets the "order" field.
func (rc *RoleCreate) SetOrder(i int) *RoleCreate {
	rc.mutation.SetOrder(i)
	return rc
}

// SetIsChangeable sets the "is_changeable" field.
func (rc *RoleCreate) SetIsChangeable(b bool) *RoleCreate {
	rc.mutation.SetIsChangeable(b)
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(x xid.ID) *RoleCreate {
	rc.mutation.SetID(x)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RoleCreate) SetNillableID(x *xid.ID) *RoleCreate {
	if x != nil {
		rc.SetID(*x)
	}
	return rc
}

// AddAdminIDs adds the "admins" edge to the Admin entity by IDs.
func (rc *RoleCreate) AddAdminIDs(ids ...xid.ID) *RoleCreate {
	rc.mutation.AddAdminIDs(ids...)
	return rc
}

// AddAdmins adds the "admins" edges to the Admin entity.
func (rc *RoleCreate) AddAdmins(a ...*Admin) *RoleCreate {
	ids := make([]xid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAdminIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (rc *RoleCreate) AddPermissionIDs(ids ...xid.ID) *RoleCreate {
	rc.mutation.AddPermissionIDs(ids...)
	return rc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (rc *RoleCreate) AddPermissions(p ...*Permission) *RoleCreate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionIDs(ids...)
}

// AddRouteIDs adds the "routes" edge to the Menu entity by IDs.
func (rc *RoleCreate) AddRouteIDs(ids ...xid.ID) *RoleCreate {
	rc.mutation.AddRouteIDs(ids...)
	return rc
}

// AddRoutes adds the "routes" edges to the Menu entity.
func (rc *RoleCreate) AddRoutes(m ...*Menu) *RoleCreate {
	ids := make([]xid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddRouteIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		if role.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized role.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := role.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		if role.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized role.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := role.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.IsEnable(); !ok {
		v := role.DefaultIsEnable
		rc.mutation.SetIsEnable(v)
	}
	if _, ok := rc.mutation.DeletedAt(); !ok {
		v := role.DefaultDeletedAt
		rc.mutation.SetDeletedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		if role.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized role.DefaultID (forgotten import ent/runtime?)")
		}
		v := role.DefaultID()
		rc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Role.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Role.updated_at"`)}
	}
	if _, ok := rc.mutation.IsEnable(); !ok {
		return &ValidationError{Name: "is_enable", err: errors.New(`ent: missing required field "Role.is_enable"`)}
	}
	if _, ok := rc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Role.deleted_at"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Role.name"`)}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Role.description"`)}
	}
	if _, ok := rc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Role.order"`)}
	}
	if _, ok := rc.mutation.IsChangeable(); !ok {
		return &ValidationError{Name: "is_changeable", err: errors.New(`ent: missing required field "Role.is_changeable"`)}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeString))
	)
	_spec.Schema = rc.schemaConfig.Role
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.IsEnable(); ok {
		_spec.SetField(role.FieldIsEnable, field.TypeBool, value)
		_node.IsEnable = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeInt, value)
		_node.DeletedAt = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Order(); ok {
		_spec.SetField(role.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := rc.mutation.IsChangeable(); ok {
		_spec.SetField(role.FieldIsChangeable, field.TypeBool, value)
		_node.IsChangeable = value
	}
	if nodes := rc.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeString),
			},
		}
		edge.Schema = rc.schemaConfig.AdminRoles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeString),
			},
		}
		edge.Schema = rc.schemaConfig.RolePermissions
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		edge.Schema = rc.schemaConfig.RoleRoutes
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RoleCreate) OnConflict(opts ...sql.ConflictOption) *RoleUpsertOne {
	rc.conflict = opts
	return &RoleUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoleCreate) OnConflictColumns(columns ...string) *RoleUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertOne{
		create: rc,
	}
}

type (
	// RoleUpsertOne is the builder for "upsert"-ing
	//  one Role node.
	RoleUpsertOne struct {
		create *RoleCreate
	}

	// RoleUpsert is the "OnConflict" setter.
	RoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *RoleUpsert) SetCreatedAt(v time.Time) *RoleUpsert {
	u.Set(role.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateCreatedAt() *RoleUpsert {
	u.SetExcluded(role.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsert) SetUpdatedAt(v time.Time) *RoleUpsert {
	u.Set(role.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateUpdatedAt() *RoleUpsert {
	u.SetExcluded(role.FieldUpdatedAt)
	return u
}

// SetIsEnable sets the "is_enable" field.
func (u *RoleUpsert) SetIsEnable(v bool) *RoleUpsert {
	u.Set(role.FieldIsEnable, v)
	return u
}

// UpdateIsEnable sets the "is_enable" field to the value that was provided on create.
func (u *RoleUpsert) UpdateIsEnable() *RoleUpsert {
	u.SetExcluded(role.FieldIsEnable)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsert) SetDeletedAt(v int) *RoleUpsert {
	u.Set(role.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDeletedAt() *RoleUpsert {
	u.SetExcluded(role.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RoleUpsert) AddDeletedAt(v int) *RoleUpsert {
	u.Add(role.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *RoleUpsert) SetName(v string) *RoleUpsert {
	u.Set(role.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsert) UpdateName() *RoleUpsert {
	u.SetExcluded(role.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *RoleUpsert) SetDescription(v string) *RoleUpsert {
	u.Set(role.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDescription() *RoleUpsert {
	u.SetExcluded(role.FieldDescription)
	return u
}

// SetOrder sets the "order" field.
func (u *RoleUpsert) SetOrder(v int) *RoleUpsert {
	u.Set(role.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *RoleUpsert) UpdateOrder() *RoleUpsert {
	u.SetExcluded(role.FieldOrder)
	return u
}

// AddOrder adds v to the "order" field.
func (u *RoleUpsert) AddOrder(v int) *RoleUpsert {
	u.Add(role.FieldOrder, v)
	return u
}

// SetIsChangeable sets the "is_changeable" field.
func (u *RoleUpsert) SetIsChangeable(v bool) *RoleUpsert {
	u.Set(role.FieldIsChangeable, v)
	return u
}

// UpdateIsChangeable sets the "is_changeable" field to the value that was provided on create.
func (u *RoleUpsert) UpdateIsChangeable() *RoleUpsert {
	u.SetExcluded(role.FieldIsChangeable)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertOne) UpdateNewValues() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(role.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoleUpsertOne) Ignore() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertOne) DoNothing() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreate.OnConflict
// documentation for more info.
func (u *RoleUpsertOne) Update(set func(*RoleUpsert)) *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *RoleUpsertOne) SetCreatedAt(v time.Time) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateCreatedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsertOne) SetUpdatedAt(v time.Time) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateUpdatedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetIsEnable sets the "is_enable" field.
func (u *RoleUpsertOne) SetIsEnable(v bool) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetIsEnable(v)
	})
}

// UpdateIsEnable sets the "is_enable" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateIsEnable() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateIsEnable()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsertOne) SetDeletedAt(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RoleUpsertOne) AddDeletedAt(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDeletedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *RoleUpsertOne) SetName(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateName() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertOne) SetDescription(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// SetOrder sets the "order" field.
func (u *RoleUpsertOne) SetOrder(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *RoleUpsertOne) AddOrder(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateOrder() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateOrder()
	})
}

// SetIsChangeable sets the "is_changeable" field.
func (u *RoleUpsertOne) SetIsChangeable(v bool) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetIsChangeable(v)
	})
}

// UpdateIsChangeable sets the "is_changeable" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateIsChangeable() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateIsChangeable()
	})
}

// Exec executes the query.
func (u *RoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoleUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: RoleUpsertOne.ID is not supported by MySQL driver. Use RoleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoleUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	err      error
	builders []*RoleCreate
	conflict []sql.ConflictOption
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoleUpsertBulk {
	rcb.conflict = opts
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflictColumns(columns ...string) *RoleUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// RoleUpsertBulk is the builder for "upsert"-ing
// a bulk of Role nodes.
type RoleUpsertBulk struct {
	create *RoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertBulk) UpdateNewValues() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(role.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoleUpsertBulk) Ignore() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertBulk) DoNothing() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreateBulk.OnConflict
// documentation for more info.
func (u *RoleUpsertBulk) Update(set func(*RoleUpsert)) *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *RoleUpsertBulk) SetCreatedAt(v time.Time) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateCreatedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsertBulk) SetUpdatedAt(v time.Time) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateUpdatedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetIsEnable sets the "is_enable" field.
func (u *RoleUpsertBulk) SetIsEnable(v bool) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetIsEnable(v)
	})
}

// UpdateIsEnable sets the "is_enable" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateIsEnable() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateIsEnable()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsertBulk) SetDeletedAt(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RoleUpsertBulk) AddDeletedAt(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDeletedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *RoleUpsertBulk) SetName(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateName() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertBulk) SetDescription(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// SetOrder sets the "order" field.
func (u *RoleUpsertBulk) SetOrder(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *RoleUpsertBulk) AddOrder(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateOrder() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateOrder()
	})
}

// SetIsChangeable sets the "is_changeable" field.
func (u *RoleUpsertBulk) SetIsChangeable(v bool) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetIsChangeable(v)
	})
}

// UpdateIsChangeable sets the "is_changeable" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateIsChangeable() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateIsChangeable()
	})
}

// Exec executes the query.
func (u *RoleUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
