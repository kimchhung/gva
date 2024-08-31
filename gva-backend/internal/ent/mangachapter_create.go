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
	"github.com/gva/internal/ent/manga"
	"github.com/gva/internal/ent/mangachapter"
)

// MangaChapterCreate is the builder for creating a MangaChapter entity.
type MangaChapterCreate struct {
	config
	mutation *MangaChapterMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mcc *MangaChapterCreate) SetCreatedAt(t time.Time) *MangaChapterCreate {
	mcc.mutation.SetCreatedAt(t)
	return mcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcc *MangaChapterCreate) SetNillableCreatedAt(t *time.Time) *MangaChapterCreate {
	if t != nil {
		mcc.SetCreatedAt(*t)
	}
	return mcc
}

// SetUpdatedAt sets the "updated_at" field.
func (mcc *MangaChapterCreate) SetUpdatedAt(t time.Time) *MangaChapterCreate {
	mcc.mutation.SetUpdatedAt(t)
	return mcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mcc *MangaChapterCreate) SetNillableUpdatedAt(t *time.Time) *MangaChapterCreate {
	if t != nil {
		mcc.SetUpdatedAt(*t)
	}
	return mcc
}

// SetMangaID sets the "manga_id" field.
func (mcc *MangaChapterCreate) SetMangaID(px pxid.ID) *MangaChapterCreate {
	mcc.mutation.SetMangaID(px)
	return mcc
}

// SetTitle sets the "title" field.
func (mcc *MangaChapterCreate) SetTitle(s string) *MangaChapterCreate {
	mcc.mutation.SetTitle(s)
	return mcc
}

// SetImgURL sets the "img_url" field.
func (mcc *MangaChapterCreate) SetImgURL(s string) *MangaChapterCreate {
	mcc.mutation.SetImgURL(s)
	return mcc
}

// SetNumber sets the "number" field.
func (mcc *MangaChapterCreate) SetNumber(u uint) *MangaChapterCreate {
	mcc.mutation.SetNumber(u)
	return mcc
}

// SetProviderName sets the "provider_name" field.
func (mcc *MangaChapterCreate) SetProviderName(s string) *MangaChapterCreate {
	mcc.mutation.SetProviderName(s)
	return mcc
}

// SetChapterUpdatedAt sets the "chapter_updated_at" field.
func (mcc *MangaChapterCreate) SetChapterUpdatedAt(t time.Time) *MangaChapterCreate {
	mcc.mutation.SetChapterUpdatedAt(t)
	return mcc
}

// SetID sets the "id" field.
func (mcc *MangaChapterCreate) SetID(px pxid.ID) *MangaChapterCreate {
	mcc.mutation.SetID(px)
	return mcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mcc *MangaChapterCreate) SetNillableID(px *pxid.ID) *MangaChapterCreate {
	if px != nil {
		mcc.SetID(*px)
	}
	return mcc
}

// SetManga sets the "manga" edge to the Manga entity.
func (mcc *MangaChapterCreate) SetManga(m *Manga) *MangaChapterCreate {
	return mcc.SetMangaID(m.ID)
}

// Mutation returns the MangaChapterMutation object of the builder.
func (mcc *MangaChapterCreate) Mutation() *MangaChapterMutation {
	return mcc.mutation
}

// Save creates the MangaChapter in the database.
func (mcc *MangaChapterCreate) Save(ctx context.Context) (*MangaChapter, error) {
	mcc.defaults()
	return withHooks(ctx, mcc.sqlSave, mcc.mutation, mcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MangaChapterCreate) SaveX(ctx context.Context) *MangaChapter {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcc *MangaChapterCreate) Exec(ctx context.Context) error {
	_, err := mcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcc *MangaChapterCreate) ExecX(ctx context.Context) {
	if err := mcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcc *MangaChapterCreate) defaults() {
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		v := mangachapter.DefaultCreatedAt()
		mcc.mutation.SetCreatedAt(v)
	}
	if _, ok := mcc.mutation.UpdatedAt(); !ok {
		v := mangachapter.DefaultUpdatedAt()
		mcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mcc.mutation.ID(); !ok {
		v := mangachapter.DefaultID()
		mcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcc *MangaChapterCreate) check() error {
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MangaChapter.created_at"`)}
	}
	if _, ok := mcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MangaChapter.updated_at"`)}
	}
	if _, ok := mcc.mutation.MangaID(); !ok {
		return &ValidationError{Name: "manga_id", err: errors.New(`ent: missing required field "MangaChapter.manga_id"`)}
	}
	if v, ok := mcc.mutation.MangaID(); ok {
		if err := mangachapter.MangaIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "manga_id", err: fmt.Errorf(`ent: validator failed for field "MangaChapter.manga_id": %w`, err)}
		}
	}
	if _, ok := mcc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "MangaChapter.title"`)}
	}
	if _, ok := mcc.mutation.ImgURL(); !ok {
		return &ValidationError{Name: "img_url", err: errors.New(`ent: missing required field "MangaChapter.img_url"`)}
	}
	if _, ok := mcc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "MangaChapter.number"`)}
	}
	if _, ok := mcc.mutation.ProviderName(); !ok {
		return &ValidationError{Name: "provider_name", err: errors.New(`ent: missing required field "MangaChapter.provider_name"`)}
	}
	if _, ok := mcc.mutation.ChapterUpdatedAt(); !ok {
		return &ValidationError{Name: "chapter_updated_at", err: errors.New(`ent: missing required field "MangaChapter.chapter_updated_at"`)}
	}
	if len(mcc.mutation.MangaIDs()) == 0 {
		return &ValidationError{Name: "manga", err: errors.New(`ent: missing required edge "MangaChapter.manga"`)}
	}
	return nil
}

func (mcc *MangaChapterCreate) sqlSave(ctx context.Context) (*MangaChapter, error) {
	if err := mcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
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
	mcc.mutation.id = &_node.ID
	mcc.mutation.done = true
	return _node, nil
}

func (mcc *MangaChapterCreate) createSpec() (*MangaChapter, *sqlgraph.CreateSpec) {
	var (
		_node = &MangaChapter{config: mcc.config}
		_spec = sqlgraph.NewCreateSpec(mangachapter.Table, sqlgraph.NewFieldSpec(mangachapter.FieldID, field.TypeString))
	)
	_spec.Schema = mcc.schemaConfig.MangaChapter
	_spec.OnConflict = mcc.conflict
	if id, ok := mcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mcc.mutation.CreatedAt(); ok {
		_spec.SetField(mangachapter.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mcc.mutation.UpdatedAt(); ok {
		_spec.SetField(mangachapter.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mcc.mutation.Title(); ok {
		_spec.SetField(mangachapter.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := mcc.mutation.ImgURL(); ok {
		_spec.SetField(mangachapter.FieldImgURL, field.TypeString, value)
		_node.ImgURL = value
	}
	if value, ok := mcc.mutation.Number(); ok {
		_spec.SetField(mangachapter.FieldNumber, field.TypeUint, value)
		_node.Number = value
	}
	if value, ok := mcc.mutation.ProviderName(); ok {
		_spec.SetField(mangachapter.FieldProviderName, field.TypeString, value)
		_node.ProviderName = value
	}
	if value, ok := mcc.mutation.ChapterUpdatedAt(); ok {
		_spec.SetField(mangachapter.FieldChapterUpdatedAt, field.TypeTime, value)
		_node.ChapterUpdatedAt = value
	}
	if nodes := mcc.mutation.MangaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mangachapter.MangaTable,
			Columns: []string{mangachapter.MangaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = mcc.schemaConfig.MangaChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MangaID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MangaChapter.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MangaChapterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcc *MangaChapterCreate) OnConflict(opts ...sql.ConflictOption) *MangaChapterUpsertOne {
	mcc.conflict = opts
	return &MangaChapterUpsertOne{
		create: mcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MangaChapter.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcc *MangaChapterCreate) OnConflictColumns(columns ...string) *MangaChapterUpsertOne {
	mcc.conflict = append(mcc.conflict, sql.ConflictColumns(columns...))
	return &MangaChapterUpsertOne{
		create: mcc,
	}
}

type (
	// MangaChapterUpsertOne is the builder for "upsert"-ing
	//  one MangaChapter node.
	MangaChapterUpsertOne struct {
		create *MangaChapterCreate
	}

	// MangaChapterUpsert is the "OnConflict" setter.
	MangaChapterUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MangaChapterUpsert) SetCreatedAt(v time.Time) *MangaChapterUpsert {
	u.Set(mangachapter.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateCreatedAt() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MangaChapterUpsert) SetUpdatedAt(v time.Time) *MangaChapterUpsert {
	u.Set(mangachapter.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateUpdatedAt() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldUpdatedAt)
	return u
}

// SetMangaID sets the "manga_id" field.
func (u *MangaChapterUpsert) SetMangaID(v pxid.ID) *MangaChapterUpsert {
	u.Set(mangachapter.FieldMangaID, v)
	return u
}

// UpdateMangaID sets the "manga_id" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateMangaID() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldMangaID)
	return u
}

// SetTitle sets the "title" field.
func (u *MangaChapterUpsert) SetTitle(v string) *MangaChapterUpsert {
	u.Set(mangachapter.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateTitle() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldTitle)
	return u
}

// SetImgURL sets the "img_url" field.
func (u *MangaChapterUpsert) SetImgURL(v string) *MangaChapterUpsert {
	u.Set(mangachapter.FieldImgURL, v)
	return u
}

// UpdateImgURL sets the "img_url" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateImgURL() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldImgURL)
	return u
}

// SetNumber sets the "number" field.
func (u *MangaChapterUpsert) SetNumber(v uint) *MangaChapterUpsert {
	u.Set(mangachapter.FieldNumber, v)
	return u
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateNumber() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldNumber)
	return u
}

// AddNumber adds v to the "number" field.
func (u *MangaChapterUpsert) AddNumber(v uint) *MangaChapterUpsert {
	u.Add(mangachapter.FieldNumber, v)
	return u
}

// SetProviderName sets the "provider_name" field.
func (u *MangaChapterUpsert) SetProviderName(v string) *MangaChapterUpsert {
	u.Set(mangachapter.FieldProviderName, v)
	return u
}

// UpdateProviderName sets the "provider_name" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateProviderName() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldProviderName)
	return u
}

// SetChapterUpdatedAt sets the "chapter_updated_at" field.
func (u *MangaChapterUpsert) SetChapterUpdatedAt(v time.Time) *MangaChapterUpsert {
	u.Set(mangachapter.FieldChapterUpdatedAt, v)
	return u
}

// UpdateChapterUpdatedAt sets the "chapter_updated_at" field to the value that was provided on create.
func (u *MangaChapterUpsert) UpdateChapterUpdatedAt() *MangaChapterUpsert {
	u.SetExcluded(mangachapter.FieldChapterUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MangaChapter.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mangachapter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MangaChapterUpsertOne) UpdateNewValues() *MangaChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mangachapter.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MangaChapter.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MangaChapterUpsertOne) Ignore() *MangaChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MangaChapterUpsertOne) DoNothing() *MangaChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MangaChapterCreate.OnConflict
// documentation for more info.
func (u *MangaChapterUpsertOne) Update(set func(*MangaChapterUpsert)) *MangaChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MangaChapterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MangaChapterUpsertOne) SetCreatedAt(v time.Time) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateCreatedAt() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MangaChapterUpsertOne) SetUpdatedAt(v time.Time) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateUpdatedAt() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetMangaID sets the "manga_id" field.
func (u *MangaChapterUpsertOne) SetMangaID(v pxid.ID) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetMangaID(v)
	})
}

// UpdateMangaID sets the "manga_id" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateMangaID() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateMangaID()
	})
}

// SetTitle sets the "title" field.
func (u *MangaChapterUpsertOne) SetTitle(v string) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateTitle() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateTitle()
	})
}

// SetImgURL sets the "img_url" field.
func (u *MangaChapterUpsertOne) SetImgURL(v string) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetImgURL(v)
	})
}

// UpdateImgURL sets the "img_url" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateImgURL() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateImgURL()
	})
}

// SetNumber sets the "number" field.
func (u *MangaChapterUpsertOne) SetNumber(v uint) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *MangaChapterUpsertOne) AddNumber(v uint) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateNumber() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateNumber()
	})
}

// SetProviderName sets the "provider_name" field.
func (u *MangaChapterUpsertOne) SetProviderName(v string) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetProviderName(v)
	})
}

// UpdateProviderName sets the "provider_name" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateProviderName() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateProviderName()
	})
}

// SetChapterUpdatedAt sets the "chapter_updated_at" field.
func (u *MangaChapterUpsertOne) SetChapterUpdatedAt(v time.Time) *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetChapterUpdatedAt(v)
	})
}

// UpdateChapterUpdatedAt sets the "chapter_updated_at" field to the value that was provided on create.
func (u *MangaChapterUpsertOne) UpdateChapterUpdatedAt() *MangaChapterUpsertOne {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateChapterUpdatedAt()
	})
}

// Exec executes the query.
func (u *MangaChapterUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MangaChapterCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MangaChapterUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MangaChapterUpsertOne) ID(ctx context.Context) (id pxid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MangaChapterUpsertOne.ID is not supported by MySQL driver. Use MangaChapterUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MangaChapterUpsertOne) IDX(ctx context.Context) pxid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MangaChapterCreateBulk is the builder for creating many MangaChapter entities in bulk.
type MangaChapterCreateBulk struct {
	config
	err      error
	builders []*MangaChapterCreate
	conflict []sql.ConflictOption
}

// Save creates the MangaChapter entities in the database.
func (mccb *MangaChapterCreateBulk) Save(ctx context.Context) ([]*MangaChapter, error) {
	if mccb.err != nil {
		return nil, mccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mccb.builders))
	nodes := make([]*MangaChapter, len(mccb.builders))
	mutators := make([]Mutator, len(mccb.builders))
	for i := range mccb.builders {
		func(i int, root context.Context) {
			builder := mccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MangaChapterMutation)
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
					_, err = mutators[i+1].Mutate(root, mccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mccb *MangaChapterCreateBulk) SaveX(ctx context.Context) []*MangaChapter {
	v, err := mccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccb *MangaChapterCreateBulk) Exec(ctx context.Context) error {
	_, err := mccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccb *MangaChapterCreateBulk) ExecX(ctx context.Context) {
	if err := mccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MangaChapter.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MangaChapterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mccb *MangaChapterCreateBulk) OnConflict(opts ...sql.ConflictOption) *MangaChapterUpsertBulk {
	mccb.conflict = opts
	return &MangaChapterUpsertBulk{
		create: mccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MangaChapter.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mccb *MangaChapterCreateBulk) OnConflictColumns(columns ...string) *MangaChapterUpsertBulk {
	mccb.conflict = append(mccb.conflict, sql.ConflictColumns(columns...))
	return &MangaChapterUpsertBulk{
		create: mccb,
	}
}

// MangaChapterUpsertBulk is the builder for "upsert"-ing
// a bulk of MangaChapter nodes.
type MangaChapterUpsertBulk struct {
	create *MangaChapterCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MangaChapter.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mangachapter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MangaChapterUpsertBulk) UpdateNewValues() *MangaChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mangachapter.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MangaChapter.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MangaChapterUpsertBulk) Ignore() *MangaChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MangaChapterUpsertBulk) DoNothing() *MangaChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MangaChapterCreateBulk.OnConflict
// documentation for more info.
func (u *MangaChapterUpsertBulk) Update(set func(*MangaChapterUpsert)) *MangaChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MangaChapterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MangaChapterUpsertBulk) SetCreatedAt(v time.Time) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateCreatedAt() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MangaChapterUpsertBulk) SetUpdatedAt(v time.Time) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateUpdatedAt() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetMangaID sets the "manga_id" field.
func (u *MangaChapterUpsertBulk) SetMangaID(v pxid.ID) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetMangaID(v)
	})
}

// UpdateMangaID sets the "manga_id" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateMangaID() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateMangaID()
	})
}

// SetTitle sets the "title" field.
func (u *MangaChapterUpsertBulk) SetTitle(v string) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateTitle() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateTitle()
	})
}

// SetImgURL sets the "img_url" field.
func (u *MangaChapterUpsertBulk) SetImgURL(v string) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetImgURL(v)
	})
}

// UpdateImgURL sets the "img_url" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateImgURL() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateImgURL()
	})
}

// SetNumber sets the "number" field.
func (u *MangaChapterUpsertBulk) SetNumber(v uint) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *MangaChapterUpsertBulk) AddNumber(v uint) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateNumber() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateNumber()
	})
}

// SetProviderName sets the "provider_name" field.
func (u *MangaChapterUpsertBulk) SetProviderName(v string) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetProviderName(v)
	})
}

// UpdateProviderName sets the "provider_name" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateProviderName() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateProviderName()
	})
}

// SetChapterUpdatedAt sets the "chapter_updated_at" field.
func (u *MangaChapterUpsertBulk) SetChapterUpdatedAt(v time.Time) *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.SetChapterUpdatedAt(v)
	})
}

// UpdateChapterUpdatedAt sets the "chapter_updated_at" field to the value that was provided on create.
func (u *MangaChapterUpsertBulk) UpdateChapterUpdatedAt() *MangaChapterUpsertBulk {
	return u.Update(func(s *MangaChapterUpsert) {
		s.UpdateChapterUpdatedAt()
	})
}

// Exec executes the query.
func (u *MangaChapterUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MangaChapterCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MangaChapterCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MangaChapterUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}