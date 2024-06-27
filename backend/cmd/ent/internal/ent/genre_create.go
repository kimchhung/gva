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
	"github.com/gva/internal/ent/genre"
)

// GenreCreate is the builder for creating a Genre entity.
type GenreCreate struct {
	config
	mutation *GenreMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gc *GenreCreate) SetCreatedAt(t time.Time) *GenreCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GenreCreate) SetNillableCreatedAt(t *time.Time) *GenreCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GenreCreate) SetUpdatedAt(t time.Time) *GenreCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GenreCreate) SetNillableUpdatedAt(t *time.Time) *GenreCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GenreCreate) SetName(s string) *GenreCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetType sets the "type" field.
func (gc *GenreCreate) SetType(ge genre.Type) *GenreCreate {
	gc.mutation.SetType(ge)
	return gc
}

// SetID sets the "id" field.
func (gc *GenreCreate) SetID(s string) *GenreCreate {
	gc.mutation.SetID(s)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GenreCreate) SetNillableID(s *string) *GenreCreate {
	if s != nil {
		gc.SetID(*s)
	}
	return gc
}

// Mutation returns the GenreMutation object of the builder.
func (gc *GenreCreate) Mutation() *GenreMutation {
	return gc.mutation
}

// Save creates the Genre in the database.
func (gc *GenreCreate) Save(ctx context.Context) (*Genre, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenreCreate) SaveX(ctx context.Context) *Genre {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GenreCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GenreCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GenreCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := genre.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := genre.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := genre.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GenreCreate) check() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Genre.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Genre.updated_at"`)}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Genre.name"`)}
	}
	if _, ok := gc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Genre.type"`)}
	}
	if v, ok := gc.mutation.GetType(); ok {
		if err := genre.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Genre.type": %w`, err)}
		}
	}
	return nil
}

func (gc *GenreCreate) sqlSave(ctx context.Context) (*Genre, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Genre.ID type: %T", _spec.ID.Value)
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GenreCreate) createSpec() (*Genre, *sqlgraph.CreateSpec) {
	var (
		_node = &Genre{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(genre.Table, sqlgraph.NewFieldSpec(genre.FieldID, field.TypeString))
	)
	_spec.Schema = gc.schemaConfig.Genre
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(genre.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(genre.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(genre.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.GetType(); ok {
		_spec.SetField(genre.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Genre.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GenreUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gc *GenreCreate) OnConflict(opts ...sql.ConflictOption) *GenreUpsertOne {
	gc.conflict = opts
	return &GenreUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Genre.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GenreCreate) OnConflictColumns(columns ...string) *GenreUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GenreUpsertOne{
		create: gc,
	}
}

type (
	// GenreUpsertOne is the builder for "upsert"-ing
	//  one Genre node.
	GenreUpsertOne struct {
		create *GenreCreate
	}

	// GenreUpsert is the "OnConflict" setter.
	GenreUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *GenreUpsert) SetCreatedAt(v time.Time) *GenreUpsert {
	u.Set(genre.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GenreUpsert) UpdateCreatedAt() *GenreUpsert {
	u.SetExcluded(genre.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GenreUpsert) SetUpdatedAt(v time.Time) *GenreUpsert {
	u.Set(genre.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GenreUpsert) UpdateUpdatedAt() *GenreUpsert {
	u.SetExcluded(genre.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *GenreUpsert) SetName(v string) *GenreUpsert {
	u.Set(genre.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GenreUpsert) UpdateName() *GenreUpsert {
	u.SetExcluded(genre.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *GenreUpsert) SetType(v genre.Type) *GenreUpsert {
	u.Set(genre.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GenreUpsert) UpdateType() *GenreUpsert {
	u.SetExcluded(genre.FieldType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Genre.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(genre.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GenreUpsertOne) UpdateNewValues() *GenreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(genre.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Genre.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GenreUpsertOne) Ignore() *GenreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GenreUpsertOne) DoNothing() *GenreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GenreCreate.OnConflict
// documentation for more info.
func (u *GenreUpsertOne) Update(set func(*GenreUpsert)) *GenreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GenreUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GenreUpsertOne) SetCreatedAt(v time.Time) *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GenreUpsertOne) UpdateCreatedAt() *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GenreUpsertOne) SetUpdatedAt(v time.Time) *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GenreUpsertOne) UpdateUpdatedAt() *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *GenreUpsertOne) SetName(v string) *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GenreUpsertOne) UpdateName() *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *GenreUpsertOne) SetType(v genre.Type) *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GenreUpsertOne) UpdateType() *GenreUpsertOne {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *GenreUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GenreCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GenreUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GenreUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GenreUpsertOne.ID is not supported by MySQL driver. Use GenreUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GenreUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GenreCreateBulk is the builder for creating many Genre entities in bulk.
type GenreCreateBulk struct {
	config
	err      error
	builders []*GenreCreate
	conflict []sql.ConflictOption
}

// Save creates the Genre entities in the database.
func (gcb *GenreCreateBulk) Save(ctx context.Context) ([]*Genre, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Genre, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenreMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GenreCreateBulk) SaveX(ctx context.Context) []*Genre {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GenreCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GenreCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Genre.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GenreUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gcb *GenreCreateBulk) OnConflict(opts ...sql.ConflictOption) *GenreUpsertBulk {
	gcb.conflict = opts
	return &GenreUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Genre.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GenreCreateBulk) OnConflictColumns(columns ...string) *GenreUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GenreUpsertBulk{
		create: gcb,
	}
}

// GenreUpsertBulk is the builder for "upsert"-ing
// a bulk of Genre nodes.
type GenreUpsertBulk struct {
	create *GenreCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Genre.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(genre.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GenreUpsertBulk) UpdateNewValues() *GenreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(genre.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Genre.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GenreUpsertBulk) Ignore() *GenreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GenreUpsertBulk) DoNothing() *GenreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GenreCreateBulk.OnConflict
// documentation for more info.
func (u *GenreUpsertBulk) Update(set func(*GenreUpsert)) *GenreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GenreUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GenreUpsertBulk) SetCreatedAt(v time.Time) *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GenreUpsertBulk) UpdateCreatedAt() *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GenreUpsertBulk) SetUpdatedAt(v time.Time) *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GenreUpsertBulk) UpdateUpdatedAt() *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *GenreUpsertBulk) SetName(v string) *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GenreUpsertBulk) UpdateName() *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *GenreUpsertBulk) SetType(v genre.Type) *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GenreUpsertBulk) UpdateType() *GenreUpsertBulk {
	return u.Update(func(s *GenreUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *GenreUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GenreCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GenreCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GenreUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
