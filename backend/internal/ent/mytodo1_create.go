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
	"github.com/gva/internal/ent/mytodo1"
)

// MyTodo1Create is the builder for creating a MyTodo1 entity.
type MyTodo1Create struct {
	config
	mutation *MyTodo1Mutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mt *MyTodo1Create) SetCreatedAt(t time.Time) *MyTodo1Create {
	mt.mutation.SetCreatedAt(t)
	return mt
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mt *MyTodo1Create) SetNillableCreatedAt(t *time.Time) *MyTodo1Create {
	if t != nil {
		mt.SetCreatedAt(*t)
	}
	return mt
}

// SetUpdatedAt sets the "updated_at" field.
func (mt *MyTodo1Create) SetUpdatedAt(t time.Time) *MyTodo1Create {
	mt.mutation.SetUpdatedAt(t)
	return mt
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mt *MyTodo1Create) SetNillableUpdatedAt(t *time.Time) *MyTodo1Create {
	if t != nil {
		mt.SetUpdatedAt(*t)
	}
	return mt
}

// SetDeletedAt sets the "deleted_at" field.
func (mt *MyTodo1Create) SetDeletedAt(i int) *MyTodo1Create {
	mt.mutation.SetDeletedAt(i)
	return mt
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mt *MyTodo1Create) SetNillableDeletedAt(i *int) *MyTodo1Create {
	if i != nil {
		mt.SetDeletedAt(*i)
	}
	return mt
}

// SetName sets the "name" field.
func (mt *MyTodo1Create) SetName(s string) *MyTodo1Create {
	mt.mutation.SetName(s)
	return mt
}

// SetID sets the "id" field.
func (mt *MyTodo1Create) SetID(x xid.ID) *MyTodo1Create {
	mt.mutation.SetID(x)
	return mt
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mt *MyTodo1Create) SetNillableID(x *xid.ID) *MyTodo1Create {
	if x != nil {
		mt.SetID(*x)
	}
	return mt
}

// Mutation returns the MyTodo1Mutation object of the builder.
func (mt *MyTodo1Create) Mutation() *MyTodo1Mutation {
	return mt.mutation
}

// Save creates the MyTodo1 in the database.
func (mt *MyTodo1Create) Save(ctx context.Context) (*MyTodo1, error) {
	if err := mt.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mt.sqlSave, mt.mutation, mt.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mt *MyTodo1Create) SaveX(ctx context.Context) *MyTodo1 {
	v, err := mt.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mt *MyTodo1Create) Exec(ctx context.Context) error {
	_, err := mt.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mt *MyTodo1Create) ExecX(ctx context.Context) {
	if err := mt.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mt *MyTodo1Create) defaults() error {
	if _, ok := mt.mutation.CreatedAt(); !ok {
		if mytodo1.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized mytodo1.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := mytodo1.DefaultCreatedAt()
		mt.mutation.SetCreatedAt(v)
	}
	if _, ok := mt.mutation.UpdatedAt(); !ok {
		if mytodo1.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized mytodo1.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := mytodo1.DefaultUpdatedAt()
		mt.mutation.SetUpdatedAt(v)
	}
	if _, ok := mt.mutation.DeletedAt(); !ok {
		v := mytodo1.DefaultDeletedAt
		mt.mutation.SetDeletedAt(v)
	}
	if _, ok := mt.mutation.ID(); !ok {
		if mytodo1.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized mytodo1.DefaultID (forgotten import ent/runtime?)")
		}
		v := mytodo1.DefaultID()
		mt.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mt *MyTodo1Create) check() error {
	if _, ok := mt.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MyTodo1.created_at"`)}
	}
	if _, ok := mt.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MyTodo1.updated_at"`)}
	}
	if _, ok := mt.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "MyTodo1.deleted_at"`)}
	}
	if _, ok := mt.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MyTodo1.name"`)}
	}
	return nil
}

func (mt *MyTodo1Create) sqlSave(ctx context.Context) (*MyTodo1, error) {
	if err := mt.check(); err != nil {
		return nil, err
	}
	_node, _spec := mt.createSpec()
	if err := sqlgraph.CreateNode(ctx, mt.driver, _spec); err != nil {
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
	mt.mutation.id = &_node.ID
	mt.mutation.done = true
	return _node, nil
}

func (mt *MyTodo1Create) createSpec() (*MyTodo1, *sqlgraph.CreateSpec) {
	var (
		_node = &MyTodo1{config: mt.config}
		_spec = sqlgraph.NewCreateSpec(mytodo1.Table, sqlgraph.NewFieldSpec(mytodo1.FieldID, field.TypeString))
	)
	_spec.Schema = mt.schemaConfig.MyTodo1
	_spec.OnConflict = mt.conflict
	if id, ok := mt.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mt.mutation.CreatedAt(); ok {
		_spec.SetField(mytodo1.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mt.mutation.UpdatedAt(); ok {
		_spec.SetField(mytodo1.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mt.mutation.DeletedAt(); ok {
		_spec.SetField(mytodo1.FieldDeletedAt, field.TypeInt, value)
		_node.DeletedAt = value
	}
	if value, ok := mt.mutation.Name(); ok {
		_spec.SetField(mytodo1.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MyTodo1.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MyTodo1Upsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mt *MyTodo1Create) OnConflict(opts ...sql.ConflictOption) *MyTodo1UpsertOne {
	mt.conflict = opts
	return &MyTodo1UpsertOne{
		create: mt,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MyTodo1.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mt *MyTodo1Create) OnConflictColumns(columns ...string) *MyTodo1UpsertOne {
	mt.conflict = append(mt.conflict, sql.ConflictColumns(columns...))
	return &MyTodo1UpsertOne{
		create: mt,
	}
}

type (
	// MyTodo1UpsertOne is the builder for "upsert"-ing
	//  one MyTodo1 node.
	MyTodo1UpsertOne struct {
		create *MyTodo1Create
	}

	// MyTodo1Upsert is the "OnConflict" setter.
	MyTodo1Upsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MyTodo1Upsert) SetCreatedAt(v time.Time) *MyTodo1Upsert {
	u.Set(mytodo1.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MyTodo1Upsert) UpdateCreatedAt() *MyTodo1Upsert {
	u.SetExcluded(mytodo1.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MyTodo1Upsert) SetUpdatedAt(v time.Time) *MyTodo1Upsert {
	u.Set(mytodo1.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MyTodo1Upsert) UpdateUpdatedAt() *MyTodo1Upsert {
	u.SetExcluded(mytodo1.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MyTodo1Upsert) SetDeletedAt(v int) *MyTodo1Upsert {
	u.Set(mytodo1.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MyTodo1Upsert) UpdateDeletedAt() *MyTodo1Upsert {
	u.SetExcluded(mytodo1.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MyTodo1Upsert) AddDeletedAt(v int) *MyTodo1Upsert {
	u.Add(mytodo1.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *MyTodo1Upsert) SetName(v string) *MyTodo1Upsert {
	u.Set(mytodo1.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MyTodo1Upsert) UpdateName() *MyTodo1Upsert {
	u.SetExcluded(mytodo1.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MyTodo1.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mytodo1.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MyTodo1UpsertOne) UpdateNewValues() *MyTodo1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mytodo1.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MyTodo1.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MyTodo1UpsertOne) Ignore() *MyTodo1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MyTodo1UpsertOne) DoNothing() *MyTodo1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MyTodo1Create.OnConflict
// documentation for more info.
func (u *MyTodo1UpsertOne) Update(set func(*MyTodo1Upsert)) *MyTodo1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MyTodo1Upsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MyTodo1UpsertOne) SetCreatedAt(v time.Time) *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MyTodo1UpsertOne) UpdateCreatedAt() *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MyTodo1UpsertOne) SetUpdatedAt(v time.Time) *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MyTodo1UpsertOne) UpdateUpdatedAt() *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MyTodo1UpsertOne) SetDeletedAt(v int) *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MyTodo1UpsertOne) AddDeletedAt(v int) *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MyTodo1UpsertOne) UpdateDeletedAt() *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *MyTodo1UpsertOne) SetName(v string) *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MyTodo1UpsertOne) UpdateName() *MyTodo1UpsertOne {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *MyTodo1UpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MyTodo1Create.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MyTodo1UpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MyTodo1UpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MyTodo1UpsertOne.ID is not supported by MySQL driver. Use MyTodo1UpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MyTodo1UpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MyTodo1CreateBulk is the builder for creating many MyTodo1 entities in bulk.
type MyTodo1CreateBulk struct {
	config
	err      error
	builders []*MyTodo1Create
	conflict []sql.ConflictOption
}

// Save creates the MyTodo1 entities in the database.
func (mtb *MyTodo1CreateBulk) Save(ctx context.Context) ([]*MyTodo1, error) {
	if mtb.err != nil {
		return nil, mtb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mtb.builders))
	nodes := make([]*MyTodo1, len(mtb.builders))
	mutators := make([]Mutator, len(mtb.builders))
	for i := range mtb.builders {
		func(i int, root context.Context) {
			builder := mtb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MyTodo1Mutation)
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
					_, err = mutators[i+1].Mutate(root, mtb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mtb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtb *MyTodo1CreateBulk) SaveX(ctx context.Context) []*MyTodo1 {
	v, err := mtb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtb *MyTodo1CreateBulk) Exec(ctx context.Context) error {
	_, err := mtb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtb *MyTodo1CreateBulk) ExecX(ctx context.Context) {
	if err := mtb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MyTodo1.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MyTodo1Upsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mtb *MyTodo1CreateBulk) OnConflict(opts ...sql.ConflictOption) *MyTodo1UpsertBulk {
	mtb.conflict = opts
	return &MyTodo1UpsertBulk{
		create: mtb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MyTodo1.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mtb *MyTodo1CreateBulk) OnConflictColumns(columns ...string) *MyTodo1UpsertBulk {
	mtb.conflict = append(mtb.conflict, sql.ConflictColumns(columns...))
	return &MyTodo1UpsertBulk{
		create: mtb,
	}
}

// MyTodo1UpsertBulk is the builder for "upsert"-ing
// a bulk of MyTodo1 nodes.
type MyTodo1UpsertBulk struct {
	create *MyTodo1CreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MyTodo1.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mytodo1.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MyTodo1UpsertBulk) UpdateNewValues() *MyTodo1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mytodo1.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MyTodo1.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MyTodo1UpsertBulk) Ignore() *MyTodo1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MyTodo1UpsertBulk) DoNothing() *MyTodo1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MyTodo1CreateBulk.OnConflict
// documentation for more info.
func (u *MyTodo1UpsertBulk) Update(set func(*MyTodo1Upsert)) *MyTodo1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MyTodo1Upsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MyTodo1UpsertBulk) SetCreatedAt(v time.Time) *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MyTodo1UpsertBulk) UpdateCreatedAt() *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MyTodo1UpsertBulk) SetUpdatedAt(v time.Time) *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MyTodo1UpsertBulk) UpdateUpdatedAt() *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MyTodo1UpsertBulk) SetDeletedAt(v int) *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MyTodo1UpsertBulk) AddDeletedAt(v int) *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MyTodo1UpsertBulk) UpdateDeletedAt() *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *MyTodo1UpsertBulk) SetName(v string) *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MyTodo1UpsertBulk) UpdateName() *MyTodo1UpsertBulk {
	return u.Update(func(s *MyTodo1Upsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *MyTodo1UpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MyTodo1CreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MyTodo1CreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MyTodo1UpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
