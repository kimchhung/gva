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
	"github.com/gva/internal/ent/todo"
)

// TodoCreate is the builder for creating a Todo entity.
type TodoCreate struct {
	config
	mutation *TodoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TodoCreate) SetCreatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableCreatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TodoCreate) SetUpdatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableUpdatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TodoCreate) SetDeletedAt(i int) *TodoCreate {
	tc.mutation.SetDeletedAt(i)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableDeletedAt(i *int) *TodoCreate {
	if i != nil {
		tc.SetDeletedAt(*i)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TodoCreate) SetName(s string) *TodoCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TodoCreate) SetID(x xid.ID) *TodoCreate {
	tc.mutation.SetID(x)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TodoCreate) SetNillableID(x *xid.ID) *TodoCreate {
	if x != nil {
		tc.SetID(*x)
	}
	return tc
}

// Mutation returns the TodoMutation object of the builder.
func (tc *TodoCreate) Mutation() *TodoMutation {
	return tc.mutation
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*Todo, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodoCreate) SaveX(ctx context.Context) *Todo {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TodoCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TodoCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TodoCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if todo.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized todo.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := todo.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if todo.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized todo.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := todo.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		v := todo.DefaultDeletedAt
		tc.mutation.SetDeletedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		if todo.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized todo.DefaultID (forgotten import ent/runtime?)")
		}
		v := todo.DefaultID()
		tc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TodoCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Todo.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Todo.updated_at"`)}
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Todo.deleted_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Todo.name"`)}
	}
	return nil
}

func (tc *TodoCreate) sqlSave(ctx context.Context) (*Todo, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TodoCreate) createSpec() (*Todo, *sqlgraph.CreateSpec) {
	var (
		_node = &Todo{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(todo.Table, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	)
	_spec.Schema = tc.schemaConfig.Todo
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(todo.FieldDeletedAt, field.TypeInt, value)
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(todo.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Todo.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TodoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tc *TodoCreate) OnConflict(opts ...sql.ConflictOption) *TodoUpsertOne {
	tc.conflict = opts
	return &TodoUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TodoCreate) OnConflictColumns(columns ...string) *TodoUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TodoUpsertOne{
		create: tc,
	}
}

type (
	// TodoUpsertOne is the builder for "upsert"-ing
	//  one Todo node.
	TodoUpsertOne struct {
		create *TodoCreate
	}

	// TodoUpsert is the "OnConflict" setter.
	TodoUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *TodoUpsert) SetCreatedAt(v time.Time) *TodoUpsert {
	u.Set(todo.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TodoUpsert) UpdateCreatedAt() *TodoUpsert {
	u.SetExcluded(todo.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TodoUpsert) SetUpdatedAt(v time.Time) *TodoUpsert {
	u.Set(todo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TodoUpsert) UpdateUpdatedAt() *TodoUpsert {
	u.SetExcluded(todo.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TodoUpsert) SetDeletedAt(v int) *TodoUpsert {
	u.Set(todo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TodoUpsert) UpdateDeletedAt() *TodoUpsert {
	u.SetExcluded(todo.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TodoUpsert) AddDeletedAt(v int) *TodoUpsert {
	u.Add(todo.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *TodoUpsert) SetName(v string) *TodoUpsert {
	u.Set(todo.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TodoUpsert) UpdateName() *TodoUpsert {
	u.SetExcluded(todo.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(todo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TodoUpsertOne) UpdateNewValues() *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(todo.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Todo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TodoUpsertOne) Ignore() *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TodoUpsertOne) DoNothing() *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TodoCreate.OnConflict
// documentation for more info.
func (u *TodoUpsertOne) Update(set func(*TodoUpsert)) *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TodoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TodoUpsertOne) SetCreatedAt(v time.Time) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateCreatedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TodoUpsertOne) SetUpdatedAt(v time.Time) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateUpdatedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TodoUpsertOne) SetDeletedAt(v int) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TodoUpsertOne) AddDeletedAt(v int) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateDeletedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *TodoUpsertOne) SetName(v string) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateName() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *TodoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TodoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TodoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TodoUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TodoUpsertOne.ID is not supported by MySQL driver. Use TodoUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TodoUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TodoCreateBulk is the builder for creating many Todo entities in bulk.
type TodoCreateBulk struct {
	config
	err      error
	builders []*TodoCreate
	conflict []sql.ConflictOption
}

// Save creates the Todo entities in the database.
func (tcb *TodoCreateBulk) Save(ctx context.Context) ([]*Todo, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Todo, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TodoMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TodoCreateBulk) SaveX(ctx context.Context) []*Todo {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TodoCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TodoCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Todo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TodoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tcb *TodoCreateBulk) OnConflict(opts ...sql.ConflictOption) *TodoUpsertBulk {
	tcb.conflict = opts
	return &TodoUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TodoCreateBulk) OnConflictColumns(columns ...string) *TodoUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TodoUpsertBulk{
		create: tcb,
	}
}

// TodoUpsertBulk is the builder for "upsert"-ing
// a bulk of Todo nodes.
type TodoUpsertBulk struct {
	create *TodoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(todo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TodoUpsertBulk) UpdateNewValues() *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(todo.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TodoUpsertBulk) Ignore() *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TodoUpsertBulk) DoNothing() *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TodoCreateBulk.OnConflict
// documentation for more info.
func (u *TodoUpsertBulk) Update(set func(*TodoUpsert)) *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TodoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TodoUpsertBulk) SetCreatedAt(v time.Time) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateCreatedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TodoUpsertBulk) SetUpdatedAt(v time.Time) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateUpdatedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TodoUpsertBulk) SetDeletedAt(v int) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TodoUpsertBulk) AddDeletedAt(v int) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateDeletedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *TodoUpsertBulk) SetName(v string) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateName() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *TodoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TodoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TodoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TodoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
