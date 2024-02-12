// Code generated by ent, DO NOT EDIT.

package ent

import (
	"gva/internal/ent/todo2"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// Todo2Create is the builder for creating a Todo2 entity.
type Todo2Create struct {
	config
	mutation *Todo2Mutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (t *Todo2Create) SetCreatedAt(value time.Time) *Todo2Create {
	t.mutation.SetCreatedAt(value)
	return t
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (t *Todo2Create) SetNillableCreatedAt(value *time.Time) *Todo2Create {
	if value != nil {
		t.SetCreatedAt(*value)
	}
	return t
}

// SetUpdatedAt sets the "updated_at" field.
func (t *Todo2Create) SetUpdatedAt(value time.Time) *Todo2Create {
	t.mutation.SetUpdatedAt(value)
	return t
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (t *Todo2Create) SetNillableUpdatedAt(value *time.Time) *Todo2Create {
	if value != nil {
		t.SetUpdatedAt(*value)
	}
	return t
}

// SetDeletedAt sets the "deleted_at" field.
func (t *Todo2Create) SetDeletedAt(value time.Time) *Todo2Create {
	t.mutation.SetDeletedAt(value)
	return t
}

// Mutation returns the Todo2Mutation object of the builder.
func (t *Todo2Create) Mutation() *Todo2Mutation {
	return t.mutation
}

// Save creates the Todo2 in the database.
func (t *Todo2Create) Save(ctx context.Context) (*Todo2, error) {
	t.defaults()
	return withHooks(ctx, t.sqlSave, t.mutation, t.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (t *Todo2Create) SaveX(ctx context.Context) *Todo2 {
	v, err := t.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (t *Todo2Create) Exec(ctx context.Context) error {
	_, err := t.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (t *Todo2Create) ExecX(ctx context.Context) {
	if err := t.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (t *Todo2Create) defaults() {
	if _, ok := t.mutation.CreatedAt(); !ok {
		v := todo2.DefaultCreatedAt()
		t.mutation.SetCreatedAt(v)
	}
	if _, ok := t.mutation.UpdatedAt(); !ok {
		v := todo2.DefaultUpdatedAt()
		t.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (t *Todo2Create) check() error {
	if _, ok := t.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Todo2.created_at"`)}
	}
	if _, ok := t.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Todo2.updated_at"`)}
	}
	if _, ok := t.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Todo2.deleted_at"`)}
	}
	return nil
}

func (t *Todo2Create) sqlSave(ctx context.Context) (*Todo2, error) {
	if err := t.check(); err != nil {
		return nil, err
	}
	_node, _spec := t.createSpec()
	if err := sqlgraph.CreateNode(ctx, t.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	t.mutation.id = &_node.ID
	t.mutation.done = true
	return _node, nil
}

func (t *Todo2Create) createSpec() (*Todo2, *sqlgraph.CreateSpec) {
	var (
		_node = &Todo2{config: t.config}
		_spec = sqlgraph.NewCreateSpec(todo2.Table, sqlgraph.NewFieldSpec(todo2.FieldID, field.TypeInt))
	)
	if value, ok := t.mutation.CreatedAt(); ok {
		_spec.SetField(todo2.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := t.mutation.UpdatedAt(); ok {
		_spec.SetField(todo2.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := t.mutation.DeletedAt(); ok {
		_spec.SetField(todo2.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// Todo2CreateBulk is the builder for creating many Todo2 entities in bulk.
type Todo2CreateBulk struct {
	config
	err      error
	builders []*Todo2Create
}

// Save creates the Todo2 entities in the database.
func (tb *Todo2CreateBulk) Save(ctx context.Context) ([]*Todo2, error) {
	if tb.err != nil {
		return nil, tb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tb.builders))
	nodes := make([]*Todo2, len(tb.builders))
	mutators := make([]Mutator, len(tb.builders))
	for i := range tb.builders {
		func(i int, root context.Context) {
			builder := tb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Todo2Mutation)
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
					_, err = mutators[i+1].Mutate(root, tb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tb *Todo2CreateBulk) SaveX(ctx context.Context) []*Todo2 {
	v, err := tb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tb *Todo2CreateBulk) Exec(ctx context.Context) error {
	_, err := tb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tb *Todo2CreateBulk) ExecX(ctx context.Context) {
	if err := tb.Exec(ctx); err != nil {
		panic(err)
	}
}
