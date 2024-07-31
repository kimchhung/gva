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
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dc *DepartmentCreate) SetCreatedAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableCreatedAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DepartmentCreate) SetUpdatedAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableUpdatedAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DepartmentCreate) SetDeletedAt(i int) *DepartmentCreate {
	dc.mutation.SetDeletedAt(i)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableDeletedAt(i *int) *DepartmentCreate {
	if i != nil {
		dc.SetDeletedAt(*i)
	}
	return dc
}

// SetIsEnable sets the "is_enable" field.
func (dc *DepartmentCreate) SetIsEnable(b bool) *DepartmentCreate {
	dc.mutation.SetIsEnable(b)
	return dc
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableIsEnable(b *bool) *DepartmentCreate {
	if b != nil {
		dc.SetIsEnable(*b)
	}
	return dc
}

// SetNameID sets the "name_id" field.
func (dc *DepartmentCreate) SetNameID(s string) *DepartmentCreate {
	dc.mutation.SetNameID(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DepartmentCreate) SetName(s string) *DepartmentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetPid sets the "pid" field.
func (dc *DepartmentCreate) SetPid(px pxid.ID) *DepartmentCreate {
	dc.mutation.SetPid(px)
	return dc
}

// SetNillablePid sets the "pid" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillablePid(px *pxid.ID) *DepartmentCreate {
	if px != nil {
		dc.SetPid(*px)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DepartmentCreate) SetID(px pxid.ID) *DepartmentCreate {
	dc.mutation.SetID(px)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableID(px *pxid.ID) *DepartmentCreate {
	if px != nil {
		dc.SetID(*px)
	}
	return dc
}

// SetParentID sets the "parent" edge to the Department entity by ID.
func (dc *DepartmentCreate) SetParentID(id pxid.ID) *DepartmentCreate {
	dc.mutation.SetParentID(id)
	return dc
}

// SetNillableParentID sets the "parent" edge to the Department entity by ID if the given value is not nil.
func (dc *DepartmentCreate) SetNillableParentID(id *pxid.ID) *DepartmentCreate {
	if id != nil {
		dc = dc.SetParentID(*id)
	}
	return dc
}

// SetParent sets the "parent" edge to the Department entity.
func (dc *DepartmentCreate) SetParent(d *Department) *DepartmentCreate {
	return dc.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Department entity by IDs.
func (dc *DepartmentCreate) AddChildIDs(ids ...pxid.ID) *DepartmentCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Department entity.
func (dc *DepartmentCreate) AddChildren(d ...*Department) *DepartmentCreate {
	ids := make([]pxid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the Admin entity by IDs.
func (dc *DepartmentCreate) AddMemberIDs(ids ...pxid.ID) *DepartmentCreate {
	dc.mutation.AddMemberIDs(ids...)
	return dc
}

// AddMembers adds the "members" edges to the Admin entity.
func (dc *DepartmentCreate) AddMembers(a ...*Admin) *DepartmentCreate {
	ids := make([]pxid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return dc.AddMemberIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		if department.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized department.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := department.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		if department.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized department.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := department.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		v := department.DefaultDeletedAt
		dc.mutation.SetDeletedAt(v)
	}
	if _, ok := dc.mutation.IsEnable(); !ok {
		v := department.DefaultIsEnable
		dc.mutation.SetIsEnable(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		if department.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized department.DefaultID (forgotten import ent/runtime?)")
		}
		v := department.DefaultID()
		dc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Department.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Department.updated_at"`)}
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Department.deleted_at"`)}
	}
	if _, ok := dc.mutation.IsEnable(); !ok {
		return &ValidationError{Name: "is_enable", err: errors.New(`ent: missing required field "Department.is_enable"`)}
	}
	if _, ok := dc.mutation.NameID(); !ok {
		return &ValidationError{Name: "name_id", err: errors.New(`ent: missing required field "Department.name_id"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Department.name"`)}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pxid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeString))
	)
	_spec.Schema = dc.schemaConfig.Department
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(department.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(department.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.SetField(department.FieldDeletedAt, field.TypeInt, value)
		_node.DeletedAt = value
	}
	if value, ok := dc.mutation.IsEnable(); ok {
		_spec.SetField(department.FieldIsEnable, field.TypeBool, value)
		_node.IsEnable = value
	}
	if value, ok := dc.mutation.NameID(); ok {
		_spec.SetField(department.FieldNameID, field.TypeString, value)
		_node.NameID = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := dc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		edge.Schema = dc.schemaConfig.Department
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Pid = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		edge.Schema = dc.schemaConfig.Department
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.MembersTable,
			Columns: []string{department.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeString),
			},
		}
		edge.Schema = dc.schemaConfig.Admin
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
//	client.Department.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DepartmentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dc *DepartmentCreate) OnConflict(opts ...sql.ConflictOption) *DepartmentUpsertOne {
	dc.conflict = opts
	return &DepartmentUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DepartmentCreate) OnConflictColumns(columns ...string) *DepartmentUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DepartmentUpsertOne{
		create: dc,
	}
}

type (
	// DepartmentUpsertOne is the builder for "upsert"-ing
	//  one Department node.
	DepartmentUpsertOne struct {
		create *DepartmentCreate
	}

	// DepartmentUpsert is the "OnConflict" setter.
	DepartmentUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *DepartmentUpsert) SetCreatedAt(v time.Time) *DepartmentUpsert {
	u.Set(department.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateCreatedAt() *DepartmentUpsert {
	u.SetExcluded(department.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DepartmentUpsert) SetUpdatedAt(v time.Time) *DepartmentUpsert {
	u.Set(department.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateUpdatedAt() *DepartmentUpsert {
	u.SetExcluded(department.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DepartmentUpsert) SetDeletedAt(v int) *DepartmentUpsert {
	u.Set(department.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateDeletedAt() *DepartmentUpsert {
	u.SetExcluded(department.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DepartmentUpsert) AddDeletedAt(v int) *DepartmentUpsert {
	u.Add(department.FieldDeletedAt, v)
	return u
}

// SetIsEnable sets the "is_enable" field.
func (u *DepartmentUpsert) SetIsEnable(v bool) *DepartmentUpsert {
	u.Set(department.FieldIsEnable, v)
	return u
}

// UpdateIsEnable sets the "is_enable" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateIsEnable() *DepartmentUpsert {
	u.SetExcluded(department.FieldIsEnable)
	return u
}

// SetNameID sets the "name_id" field.
func (u *DepartmentUpsert) SetNameID(v string) *DepartmentUpsert {
	u.Set(department.FieldNameID, v)
	return u
}

// UpdateNameID sets the "name_id" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateNameID() *DepartmentUpsert {
	u.SetExcluded(department.FieldNameID)
	return u
}

// SetName sets the "name" field.
func (u *DepartmentUpsert) SetName(v string) *DepartmentUpsert {
	u.Set(department.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateName() *DepartmentUpsert {
	u.SetExcluded(department.FieldName)
	return u
}

// SetPid sets the "pid" field.
func (u *DepartmentUpsert) SetPid(v pxid.ID) *DepartmentUpsert {
	u.Set(department.FieldPid, v)
	return u
}

// UpdatePid sets the "pid" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdatePid() *DepartmentUpsert {
	u.SetExcluded(department.FieldPid)
	return u
}

// ClearPid clears the value of the "pid" field.
func (u *DepartmentUpsert) ClearPid() *DepartmentUpsert {
	u.SetNull(department.FieldPid)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(department.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DepartmentUpsertOne) UpdateNewValues() *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(department.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Department.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DepartmentUpsertOne) Ignore() *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DepartmentUpsertOne) DoNothing() *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DepartmentCreate.OnConflict
// documentation for more info.
func (u *DepartmentUpsertOne) Update(set func(*DepartmentUpsert)) *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DepartmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DepartmentUpsertOne) SetCreatedAt(v time.Time) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateCreatedAt() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DepartmentUpsertOne) SetUpdatedAt(v time.Time) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateUpdatedAt() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DepartmentUpsertOne) SetDeletedAt(v int) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DepartmentUpsertOne) AddDeletedAt(v int) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateDeletedAt() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetIsEnable sets the "is_enable" field.
func (u *DepartmentUpsertOne) SetIsEnable(v bool) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetIsEnable(v)
	})
}

// UpdateIsEnable sets the "is_enable" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateIsEnable() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateIsEnable()
	})
}

// SetNameID sets the "name_id" field.
func (u *DepartmentUpsertOne) SetNameID(v string) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetNameID(v)
	})
}

// UpdateNameID sets the "name_id" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateNameID() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateNameID()
	})
}

// SetName sets the "name" field.
func (u *DepartmentUpsertOne) SetName(v string) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateName() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateName()
	})
}

// SetPid sets the "pid" field.
func (u *DepartmentUpsertOne) SetPid(v pxid.ID) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetPid(v)
	})
}

// UpdatePid sets the "pid" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdatePid() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdatePid()
	})
}

// ClearPid clears the value of the "pid" field.
func (u *DepartmentUpsertOne) ClearPid() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.ClearPid()
	})
}

// Exec executes the query.
func (u *DepartmentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DepartmentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DepartmentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DepartmentUpsertOne) ID(ctx context.Context) (id pxid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DepartmentUpsertOne.ID is not supported by MySQL driver. Use DepartmentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DepartmentUpsertOne) IDX(ctx context.Context) pxid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	err      error
	builders []*DepartmentCreate
	conflict []sql.ConflictOption
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Department.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DepartmentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dcb *DepartmentCreateBulk) OnConflict(opts ...sql.ConflictOption) *DepartmentUpsertBulk {
	dcb.conflict = opts
	return &DepartmentUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DepartmentCreateBulk) OnConflictColumns(columns ...string) *DepartmentUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DepartmentUpsertBulk{
		create: dcb,
	}
}

// DepartmentUpsertBulk is the builder for "upsert"-ing
// a bulk of Department nodes.
type DepartmentUpsertBulk struct {
	create *DepartmentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(department.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DepartmentUpsertBulk) UpdateNewValues() *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(department.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DepartmentUpsertBulk) Ignore() *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DepartmentUpsertBulk) DoNothing() *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DepartmentCreateBulk.OnConflict
// documentation for more info.
func (u *DepartmentUpsertBulk) Update(set func(*DepartmentUpsert)) *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DepartmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DepartmentUpsertBulk) SetCreatedAt(v time.Time) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateCreatedAt() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DepartmentUpsertBulk) SetUpdatedAt(v time.Time) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateUpdatedAt() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DepartmentUpsertBulk) SetDeletedAt(v int) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DepartmentUpsertBulk) AddDeletedAt(v int) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateDeletedAt() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetIsEnable sets the "is_enable" field.
func (u *DepartmentUpsertBulk) SetIsEnable(v bool) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetIsEnable(v)
	})
}

// UpdateIsEnable sets the "is_enable" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateIsEnable() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateIsEnable()
	})
}

// SetNameID sets the "name_id" field.
func (u *DepartmentUpsertBulk) SetNameID(v string) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetNameID(v)
	})
}

// UpdateNameID sets the "name_id" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateNameID() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateNameID()
	})
}

// SetName sets the "name" field.
func (u *DepartmentUpsertBulk) SetName(v string) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateName() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateName()
	})
}

// SetPid sets the "pid" field.
func (u *DepartmentUpsertBulk) SetPid(v pxid.ID) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetPid(v)
	})
}

// UpdatePid sets the "pid" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdatePid() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdatePid()
	})
}

// ClearPid clears the value of the "pid" field.
func (u *DepartmentUpsertBulk) ClearPid() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.ClearPid()
	})
}

// Exec executes the query.
func (u *DepartmentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DepartmentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DepartmentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DepartmentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
