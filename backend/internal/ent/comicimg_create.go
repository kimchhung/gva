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
	"github.com/gva/internal/ent/comicchapter"
	"github.com/gva/internal/ent/comicimg"
)

// ComicImgCreate is the builder for creating a ComicImg entity.
type ComicImgCreate struct {
	config
	mutation *ComicImgMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cic *ComicImgCreate) SetCreatedAt(t time.Time) *ComicImgCreate {
	cic.mutation.SetCreatedAt(t)
	return cic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cic *ComicImgCreate) SetNillableCreatedAt(t *time.Time) *ComicImgCreate {
	if t != nil {
		cic.SetCreatedAt(*t)
	}
	return cic
}

// SetUpdatedAt sets the "updated_at" field.
func (cic *ComicImgCreate) SetUpdatedAt(t time.Time) *ComicImgCreate {
	cic.mutation.SetUpdatedAt(t)
	return cic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cic *ComicImgCreate) SetNillableUpdatedAt(t *time.Time) *ComicImgCreate {
	if t != nil {
		cic.SetUpdatedAt(*t)
	}
	return cic
}

// SetB2key sets the "b2key" field.
func (cic *ComicImgCreate) SetB2key(s string) *ComicImgCreate {
	cic.mutation.SetB2key(s)
	return cic
}

// SetHeight sets the "height" field.
func (cic *ComicImgCreate) SetHeight(i int) *ComicImgCreate {
	cic.mutation.SetHeight(i)
	return cic
}

// SetName sets the "name" field.
func (cic *ComicImgCreate) SetName(s string) *ComicImgCreate {
	cic.mutation.SetName(s)
	return cic
}

// SetOptimizedSize sets the "optimized_size" field.
func (cic *ComicImgCreate) SetOptimizedSize(i int64) *ComicImgCreate {
	cic.mutation.SetOptimizedSize(i)
	return cic
}

// SetSize sets the "size" field.
func (cic *ComicImgCreate) SetSize(i int64) *ComicImgCreate {
	cic.mutation.SetSize(i)
	return cic
}

// SetWidth sets the "width" field.
func (cic *ComicImgCreate) SetWidth(i int) *ComicImgCreate {
	cic.mutation.SetWidth(i)
	return cic
}

// SetID sets the "id" field.
func (cic *ComicImgCreate) SetID(s string) *ComicImgCreate {
	cic.mutation.SetID(s)
	return cic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cic *ComicImgCreate) SetNillableID(s *string) *ComicImgCreate {
	if s != nil {
		cic.SetID(*s)
	}
	return cic
}

// SetChapterID sets the "chapter" edge to the ComicChapter entity by ID.
func (cic *ComicImgCreate) SetChapterID(id string) *ComicImgCreate {
	cic.mutation.SetChapterID(id)
	return cic
}

// SetNillableChapterID sets the "chapter" edge to the ComicChapter entity by ID if the given value is not nil.
func (cic *ComicImgCreate) SetNillableChapterID(id *string) *ComicImgCreate {
	if id != nil {
		cic = cic.SetChapterID(*id)
	}
	return cic
}

// SetChapter sets the "chapter" edge to the ComicChapter entity.
func (cic *ComicImgCreate) SetChapter(c *ComicChapter) *ComicImgCreate {
	return cic.SetChapterID(c.ID)
}

// Mutation returns the ComicImgMutation object of the builder.
func (cic *ComicImgCreate) Mutation() *ComicImgMutation {
	return cic.mutation
}

// Save creates the ComicImg in the database.
func (cic *ComicImgCreate) Save(ctx context.Context) (*ComicImg, error) {
	cic.defaults()
	return withHooks(ctx, cic.sqlSave, cic.mutation, cic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cic *ComicImgCreate) SaveX(ctx context.Context) *ComicImg {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *ComicImgCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *ComicImgCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cic *ComicImgCreate) defaults() {
	if _, ok := cic.mutation.CreatedAt(); !ok {
		v := comicimg.DefaultCreatedAt()
		cic.mutation.SetCreatedAt(v)
	}
	if _, ok := cic.mutation.UpdatedAt(); !ok {
		v := comicimg.DefaultUpdatedAt()
		cic.mutation.SetUpdatedAt(v)
	}
	if _, ok := cic.mutation.ID(); !ok {
		v := comicimg.DefaultID()
		cic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cic *ComicImgCreate) check() error {
	if _, ok := cic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ComicImg.created_at"`)}
	}
	if _, ok := cic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ComicImg.updated_at"`)}
	}
	if _, ok := cic.mutation.B2key(); !ok {
		return &ValidationError{Name: "b2key", err: errors.New(`ent: missing required field "ComicImg.b2key"`)}
	}
	if _, ok := cic.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "ComicImg.height"`)}
	}
	if _, ok := cic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ComicImg.name"`)}
	}
	if _, ok := cic.mutation.OptimizedSize(); !ok {
		return &ValidationError{Name: "optimized_size", err: errors.New(`ent: missing required field "ComicImg.optimized_size"`)}
	}
	if _, ok := cic.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "ComicImg.size"`)}
	}
	if _, ok := cic.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "ComicImg.width"`)}
	}
	return nil
}

func (cic *ComicImgCreate) sqlSave(ctx context.Context) (*ComicImg, error) {
	if err := cic.check(); err != nil {
		return nil, err
	}
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ComicImg.ID type: %T", _spec.ID.Value)
		}
	}
	cic.mutation.id = &_node.ID
	cic.mutation.done = true
	return _node, nil
}

func (cic *ComicImgCreate) createSpec() (*ComicImg, *sqlgraph.CreateSpec) {
	var (
		_node = &ComicImg{config: cic.config}
		_spec = sqlgraph.NewCreateSpec(comicimg.Table, sqlgraph.NewFieldSpec(comicimg.FieldID, field.TypeString))
	)
	_spec.Schema = cic.schemaConfig.ComicImg
	_spec.OnConflict = cic.conflict
	if id, ok := cic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cic.mutation.CreatedAt(); ok {
		_spec.SetField(comicimg.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cic.mutation.UpdatedAt(); ok {
		_spec.SetField(comicimg.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cic.mutation.B2key(); ok {
		_spec.SetField(comicimg.FieldB2key, field.TypeString, value)
		_node.B2key = value
	}
	if value, ok := cic.mutation.Height(); ok {
		_spec.SetField(comicimg.FieldHeight, field.TypeInt, value)
		_node.Height = value
	}
	if value, ok := cic.mutation.Name(); ok {
		_spec.SetField(comicimg.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cic.mutation.OptimizedSize(); ok {
		_spec.SetField(comicimg.FieldOptimizedSize, field.TypeInt64, value)
		_node.OptimizedSize = value
	}
	if value, ok := cic.mutation.Size(); ok {
		_spec.SetField(comicimg.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := cic.mutation.Width(); ok {
		_spec.SetField(comicimg.FieldWidth, field.TypeInt, value)
		_node.Width = value
	}
	if nodes := cic.mutation.ChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comicimg.ChapterTable,
			Columns: []string{comicimg.ChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cic.schemaConfig.ComicImg
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.comic_chapter_imgs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ComicImg.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ComicImgUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cic *ComicImgCreate) OnConflict(opts ...sql.ConflictOption) *ComicImgUpsertOne {
	cic.conflict = opts
	return &ComicImgUpsertOne{
		create: cic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ComicImg.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cic *ComicImgCreate) OnConflictColumns(columns ...string) *ComicImgUpsertOne {
	cic.conflict = append(cic.conflict, sql.ConflictColumns(columns...))
	return &ComicImgUpsertOne{
		create: cic,
	}
}

type (
	// ComicImgUpsertOne is the builder for "upsert"-ing
	//  one ComicImg node.
	ComicImgUpsertOne struct {
		create *ComicImgCreate
	}

	// ComicImgUpsert is the "OnConflict" setter.
	ComicImgUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ComicImgUpsert) SetCreatedAt(v time.Time) *ComicImgUpsert {
	u.Set(comicimg.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateCreatedAt() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ComicImgUpsert) SetUpdatedAt(v time.Time) *ComicImgUpsert {
	u.Set(comicimg.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateUpdatedAt() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldUpdatedAt)
	return u
}

// SetB2key sets the "b2key" field.
func (u *ComicImgUpsert) SetB2key(v string) *ComicImgUpsert {
	u.Set(comicimg.FieldB2key, v)
	return u
}

// UpdateB2key sets the "b2key" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateB2key() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldB2key)
	return u
}

// SetHeight sets the "height" field.
func (u *ComicImgUpsert) SetHeight(v int) *ComicImgUpsert {
	u.Set(comicimg.FieldHeight, v)
	return u
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateHeight() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldHeight)
	return u
}

// AddHeight adds v to the "height" field.
func (u *ComicImgUpsert) AddHeight(v int) *ComicImgUpsert {
	u.Add(comicimg.FieldHeight, v)
	return u
}

// SetName sets the "name" field.
func (u *ComicImgUpsert) SetName(v string) *ComicImgUpsert {
	u.Set(comicimg.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateName() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldName)
	return u
}

// SetOptimizedSize sets the "optimized_size" field.
func (u *ComicImgUpsert) SetOptimizedSize(v int64) *ComicImgUpsert {
	u.Set(comicimg.FieldOptimizedSize, v)
	return u
}

// UpdateOptimizedSize sets the "optimized_size" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateOptimizedSize() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldOptimizedSize)
	return u
}

// AddOptimizedSize adds v to the "optimized_size" field.
func (u *ComicImgUpsert) AddOptimizedSize(v int64) *ComicImgUpsert {
	u.Add(comicimg.FieldOptimizedSize, v)
	return u
}

// SetSize sets the "size" field.
func (u *ComicImgUpsert) SetSize(v int64) *ComicImgUpsert {
	u.Set(comicimg.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateSize() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *ComicImgUpsert) AddSize(v int64) *ComicImgUpsert {
	u.Add(comicimg.FieldSize, v)
	return u
}

// SetWidth sets the "width" field.
func (u *ComicImgUpsert) SetWidth(v int) *ComicImgUpsert {
	u.Set(comicimg.FieldWidth, v)
	return u
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *ComicImgUpsert) UpdateWidth() *ComicImgUpsert {
	u.SetExcluded(comicimg.FieldWidth)
	return u
}

// AddWidth adds v to the "width" field.
func (u *ComicImgUpsert) AddWidth(v int) *ComicImgUpsert {
	u.Add(comicimg.FieldWidth, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ComicImg.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(comicimg.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ComicImgUpsertOne) UpdateNewValues() *ComicImgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(comicimg.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ComicImg.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ComicImgUpsertOne) Ignore() *ComicImgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ComicImgUpsertOne) DoNothing() *ComicImgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ComicImgCreate.OnConflict
// documentation for more info.
func (u *ComicImgUpsertOne) Update(set func(*ComicImgUpsert)) *ComicImgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ComicImgUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ComicImgUpsertOne) SetCreatedAt(v time.Time) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateCreatedAt() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ComicImgUpsertOne) SetUpdatedAt(v time.Time) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateUpdatedAt() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetB2key sets the "b2key" field.
func (u *ComicImgUpsertOne) SetB2key(v string) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetB2key(v)
	})
}

// UpdateB2key sets the "b2key" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateB2key() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateB2key()
	})
}

// SetHeight sets the "height" field.
func (u *ComicImgUpsertOne) SetHeight(v int) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetHeight(v)
	})
}

// AddHeight adds v to the "height" field.
func (u *ComicImgUpsertOne) AddHeight(v int) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddHeight(v)
	})
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateHeight() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateHeight()
	})
}

// SetName sets the "name" field.
func (u *ComicImgUpsertOne) SetName(v string) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateName() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateName()
	})
}

// SetOptimizedSize sets the "optimized_size" field.
func (u *ComicImgUpsertOne) SetOptimizedSize(v int64) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetOptimizedSize(v)
	})
}

// AddOptimizedSize adds v to the "optimized_size" field.
func (u *ComicImgUpsertOne) AddOptimizedSize(v int64) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddOptimizedSize(v)
	})
}

// UpdateOptimizedSize sets the "optimized_size" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateOptimizedSize() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateOptimizedSize()
	})
}

// SetSize sets the "size" field.
func (u *ComicImgUpsertOne) SetSize(v int64) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *ComicImgUpsertOne) AddSize(v int64) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateSize() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateSize()
	})
}

// SetWidth sets the "width" field.
func (u *ComicImgUpsertOne) SetWidth(v int) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetWidth(v)
	})
}

// AddWidth adds v to the "width" field.
func (u *ComicImgUpsertOne) AddWidth(v int) *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *ComicImgUpsertOne) UpdateWidth() *ComicImgUpsertOne {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateWidth()
	})
}

// Exec executes the query.
func (u *ComicImgUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ComicImgCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ComicImgUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ComicImgUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ComicImgUpsertOne.ID is not supported by MySQL driver. Use ComicImgUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ComicImgUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ComicImgCreateBulk is the builder for creating many ComicImg entities in bulk.
type ComicImgCreateBulk struct {
	config
	err      error
	builders []*ComicImgCreate
	conflict []sql.ConflictOption
}

// Save creates the ComicImg entities in the database.
func (cicb *ComicImgCreateBulk) Save(ctx context.Context) ([]*ComicImg, error) {
	if cicb.err != nil {
		return nil, cicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*ComicImg, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ComicImgMutation)
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
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *ComicImgCreateBulk) SaveX(ctx context.Context) []*ComicImg {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *ComicImgCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *ComicImgCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ComicImg.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ComicImgUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cicb *ComicImgCreateBulk) OnConflict(opts ...sql.ConflictOption) *ComicImgUpsertBulk {
	cicb.conflict = opts
	return &ComicImgUpsertBulk{
		create: cicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ComicImg.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cicb *ComicImgCreateBulk) OnConflictColumns(columns ...string) *ComicImgUpsertBulk {
	cicb.conflict = append(cicb.conflict, sql.ConflictColumns(columns...))
	return &ComicImgUpsertBulk{
		create: cicb,
	}
}

// ComicImgUpsertBulk is the builder for "upsert"-ing
// a bulk of ComicImg nodes.
type ComicImgUpsertBulk struct {
	create *ComicImgCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ComicImg.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(comicimg.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ComicImgUpsertBulk) UpdateNewValues() *ComicImgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(comicimg.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ComicImg.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ComicImgUpsertBulk) Ignore() *ComicImgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ComicImgUpsertBulk) DoNothing() *ComicImgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ComicImgCreateBulk.OnConflict
// documentation for more info.
func (u *ComicImgUpsertBulk) Update(set func(*ComicImgUpsert)) *ComicImgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ComicImgUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ComicImgUpsertBulk) SetCreatedAt(v time.Time) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateCreatedAt() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ComicImgUpsertBulk) SetUpdatedAt(v time.Time) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateUpdatedAt() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetB2key sets the "b2key" field.
func (u *ComicImgUpsertBulk) SetB2key(v string) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetB2key(v)
	})
}

// UpdateB2key sets the "b2key" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateB2key() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateB2key()
	})
}

// SetHeight sets the "height" field.
func (u *ComicImgUpsertBulk) SetHeight(v int) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetHeight(v)
	})
}

// AddHeight adds v to the "height" field.
func (u *ComicImgUpsertBulk) AddHeight(v int) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddHeight(v)
	})
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateHeight() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateHeight()
	})
}

// SetName sets the "name" field.
func (u *ComicImgUpsertBulk) SetName(v string) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateName() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateName()
	})
}

// SetOptimizedSize sets the "optimized_size" field.
func (u *ComicImgUpsertBulk) SetOptimizedSize(v int64) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetOptimizedSize(v)
	})
}

// AddOptimizedSize adds v to the "optimized_size" field.
func (u *ComicImgUpsertBulk) AddOptimizedSize(v int64) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddOptimizedSize(v)
	})
}

// UpdateOptimizedSize sets the "optimized_size" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateOptimizedSize() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateOptimizedSize()
	})
}

// SetSize sets the "size" field.
func (u *ComicImgUpsertBulk) SetSize(v int64) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *ComicImgUpsertBulk) AddSize(v int64) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateSize() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateSize()
	})
}

// SetWidth sets the "width" field.
func (u *ComicImgUpsertBulk) SetWidth(v int) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.SetWidth(v)
	})
}

// AddWidth adds v to the "width" field.
func (u *ComicImgUpsertBulk) AddWidth(v int) *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.AddWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *ComicImgUpsertBulk) UpdateWidth() *ComicImgUpsertBulk {
	return u.Update(func(s *ComicImgUpsert) {
		s.UpdateWidth()
	})
}

// Exec executes the query.
func (u *ComicImgUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ComicImgCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ComicImgCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ComicImgUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
