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
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/comic"
	"github.com/gva/internal/ent/comicchapter"
	"github.com/gva/internal/ent/comicimg"
)

// ComicChapterCreate is the builder for creating a ComicChapter entity.
type ComicChapterCreate struct {
	config
	mutation *ComicChapterMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ccc *ComicChapterCreate) SetCreatedAt(t time.Time) *ComicChapterCreate {
	ccc.mutation.SetCreatedAt(t)
	return ccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableCreatedAt(t *time.Time) *ComicChapterCreate {
	if t != nil {
		ccc.SetCreatedAt(*t)
	}
	return ccc
}

// SetUpdatedAt sets the "updated_at" field.
func (ccc *ComicChapterCreate) SetUpdatedAt(t time.Time) *ComicChapterCreate {
	ccc.mutation.SetUpdatedAt(t)
	return ccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableUpdatedAt(t *time.Time) *ComicChapterCreate {
	if t != nil {
		ccc.SetUpdatedAt(*t)
	}
	return ccc
}

// SetChapter sets the "chapter" field.
func (ccc *ComicChapterCreate) SetChapter(u uint) *ComicChapterCreate {
	ccc.mutation.SetChapter(u)
	return ccc
}

// SetTitle sets the "title" field.
func (ccc *ComicChapterCreate) SetTitle(s string) *ComicChapterCreate {
	ccc.mutation.SetTitle(s)
	return ccc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableTitle(s *string) *ComicChapterCreate {
	if s != nil {
		ccc.SetTitle(*s)
	}
	return ccc
}

// SetVolumn sets the "volumn" field.
func (ccc *ComicChapterCreate) SetVolumn(s string) *ComicChapterCreate {
	ccc.mutation.SetVolumn(s)
	return ccc
}

// SetNillableVolumn sets the "volumn" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableVolumn(s *string) *ComicChapterCreate {
	if s != nil {
		ccc.SetVolumn(*s)
	}
	return ccc
}

// SetLang sets the "lang" field.
func (ccc *ComicChapterCreate) SetLang(s string) *ComicChapterCreate {
	ccc.mutation.SetLang(s)
	return ccc
}

// SetUpCount sets the "up_count" field.
func (ccc *ComicChapterCreate) SetUpCount(u uint) *ComicChapterCreate {
	ccc.mutation.SetUpCount(u)
	return ccc
}

// SetNillableUpCount sets the "up_count" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableUpCount(u *uint) *ComicChapterCreate {
	if u != nil {
		ccc.SetUpCount(*u)
	}
	return ccc
}

// SetDownCount sets the "down_count" field.
func (ccc *ComicChapterCreate) SetDownCount(u uint) *ComicChapterCreate {
	ccc.mutation.SetDownCount(u)
	return ccc
}

// SetNillableDownCount sets the "down_count" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableDownCount(u *uint) *ComicChapterCreate {
	if u != nil {
		ccc.SetDownCount(*u)
	}
	return ccc
}

// SetIsLastChapter sets the "is_last_chapter" field.
func (ccc *ComicChapterCreate) SetIsLastChapter(b bool) *ComicChapterCreate {
	ccc.mutation.SetIsLastChapter(b)
	return ccc
}

// SetNillableIsLastChapter sets the "is_last_chapter" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableIsLastChapter(b *bool) *ComicChapterCreate {
	if b != nil {
		ccc.SetIsLastChapter(*b)
	}
	return ccc
}

// SetID sets the "id" field.
func (ccc *ComicChapterCreate) SetID(pu pulid.ID) *ComicChapterCreate {
	ccc.mutation.SetID(pu)
	return ccc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableID(pu *pulid.ID) *ComicChapterCreate {
	if pu != nil {
		ccc.SetID(*pu)
	}
	return ccc
}

// AddImgIDs adds the "imgs" edge to the ComicImg entity by IDs.
func (ccc *ComicChapterCreate) AddImgIDs(ids ...pulid.ID) *ComicChapterCreate {
	ccc.mutation.AddImgIDs(ids...)
	return ccc
}

// AddImgs adds the "imgs" edges to the ComicImg entity.
func (ccc *ComicChapterCreate) AddImgs(c ...*ComicImg) *ComicChapterCreate {
	ids := make([]pulid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccc.AddImgIDs(ids...)
}

// SetComicID sets the "comic" edge to the Comic entity by ID.
func (ccc *ComicChapterCreate) SetComicID(id pulid.ID) *ComicChapterCreate {
	ccc.mutation.SetComicID(id)
	return ccc
}

// SetNillableComicID sets the "comic" edge to the Comic entity by ID if the given value is not nil.
func (ccc *ComicChapterCreate) SetNillableComicID(id *pulid.ID) *ComicChapterCreate {
	if id != nil {
		ccc = ccc.SetComicID(*id)
	}
	return ccc
}

// SetComic sets the "comic" edge to the Comic entity.
func (ccc *ComicChapterCreate) SetComic(c *Comic) *ComicChapterCreate {
	return ccc.SetComicID(c.ID)
}

// Mutation returns the ComicChapterMutation object of the builder.
func (ccc *ComicChapterCreate) Mutation() *ComicChapterMutation {
	return ccc.mutation
}

// Save creates the ComicChapter in the database.
func (ccc *ComicChapterCreate) Save(ctx context.Context) (*ComicChapter, error) {
	ccc.defaults()
	return withHooks(ctx, ccc.sqlSave, ccc.mutation, ccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *ComicChapterCreate) SaveX(ctx context.Context) *ComicChapter {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *ComicChapterCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *ComicChapterCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *ComicChapterCreate) defaults() {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		v := comicchapter.DefaultCreatedAt()
		ccc.mutation.SetCreatedAt(v)
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		v := comicchapter.DefaultUpdatedAt()
		ccc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ccc.mutation.UpCount(); !ok {
		v := comicchapter.DefaultUpCount
		ccc.mutation.SetUpCount(v)
	}
	if _, ok := ccc.mutation.DownCount(); !ok {
		v := comicchapter.DefaultDownCount
		ccc.mutation.SetDownCount(v)
	}
	if _, ok := ccc.mutation.IsLastChapter(); !ok {
		v := comicchapter.DefaultIsLastChapter
		ccc.mutation.SetIsLastChapter(v)
	}
	if _, ok := ccc.mutation.ID(); !ok {
		v := comicchapter.DefaultID()
		ccc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *ComicChapterCreate) check() error {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ComicChapter.created_at"`)}
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ComicChapter.updated_at"`)}
	}
	if _, ok := ccc.mutation.Chapter(); !ok {
		return &ValidationError{Name: "chapter", err: errors.New(`ent: missing required field "ComicChapter.chapter"`)}
	}
	if _, ok := ccc.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "ComicChapter.lang"`)}
	}
	if _, ok := ccc.mutation.UpCount(); !ok {
		return &ValidationError{Name: "up_count", err: errors.New(`ent: missing required field "ComicChapter.up_count"`)}
	}
	if _, ok := ccc.mutation.DownCount(); !ok {
		return &ValidationError{Name: "down_count", err: errors.New(`ent: missing required field "ComicChapter.down_count"`)}
	}
	if _, ok := ccc.mutation.IsLastChapter(); !ok {
		return &ValidationError{Name: "is_last_chapter", err: errors.New(`ent: missing required field "ComicChapter.is_last_chapter"`)}
	}
	return nil
}

func (ccc *ComicChapterCreate) sqlSave(ctx context.Context) (*ComicChapter, error) {
	if err := ccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ccc.mutation.id = &_node.ID
	ccc.mutation.done = true
	return _node, nil
}

func (ccc *ComicChapterCreate) createSpec() (*ComicChapter, *sqlgraph.CreateSpec) {
	var (
		_node = &ComicChapter{config: ccc.config}
		_spec = sqlgraph.NewCreateSpec(comicchapter.Table, sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString))
	)
	_spec.Schema = ccc.schemaConfig.ComicChapter
	_spec.OnConflict = ccc.conflict
	if id, ok := ccc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ccc.mutation.CreatedAt(); ok {
		_spec.SetField(comicchapter.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ccc.mutation.UpdatedAt(); ok {
		_spec.SetField(comicchapter.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ccc.mutation.Chapter(); ok {
		_spec.SetField(comicchapter.FieldChapter, field.TypeUint, value)
		_node.Chapter = value
	}
	if value, ok := ccc.mutation.Title(); ok {
		_spec.SetField(comicchapter.FieldTitle, field.TypeString, value)
		_node.Title = &value
	}
	if value, ok := ccc.mutation.Volumn(); ok {
		_spec.SetField(comicchapter.FieldVolumn, field.TypeString, value)
		_node.Volumn = &value
	}
	if value, ok := ccc.mutation.Lang(); ok {
		_spec.SetField(comicchapter.FieldLang, field.TypeString, value)
		_node.Lang = value
	}
	if value, ok := ccc.mutation.UpCount(); ok {
		_spec.SetField(comicchapter.FieldUpCount, field.TypeUint, value)
		_node.UpCount = value
	}
	if value, ok := ccc.mutation.DownCount(); ok {
		_spec.SetField(comicchapter.FieldDownCount, field.TypeUint, value)
		_node.DownCount = value
	}
	if value, ok := ccc.mutation.IsLastChapter(); ok {
		_spec.SetField(comicchapter.FieldIsLastChapter, field.TypeBool, value)
		_node.IsLastChapter = value
	}
	if nodes := ccc.mutation.ImgsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comicchapter.ImgsTable,
			Columns: []string{comicchapter.ImgsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicimg.FieldID, field.TypeString),
			},
		}
		edge.Schema = ccc.schemaConfig.ComicImg
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ccc.mutation.ComicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comicchapter.ComicTable,
			Columns: []string{comicchapter.ComicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comic.FieldID, field.TypeString),
			},
		}
		edge.Schema = ccc.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.comic_chapters = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ComicChapter.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ComicChapterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ccc *ComicChapterCreate) OnConflict(opts ...sql.ConflictOption) *ComicChapterUpsertOne {
	ccc.conflict = opts
	return &ComicChapterUpsertOne{
		create: ccc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ComicChapter.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccc *ComicChapterCreate) OnConflictColumns(columns ...string) *ComicChapterUpsertOne {
	ccc.conflict = append(ccc.conflict, sql.ConflictColumns(columns...))
	return &ComicChapterUpsertOne{
		create: ccc,
	}
}

type (
	// ComicChapterUpsertOne is the builder for "upsert"-ing
	//  one ComicChapter node.
	ComicChapterUpsertOne struct {
		create *ComicChapterCreate
	}

	// ComicChapterUpsert is the "OnConflict" setter.
	ComicChapterUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ComicChapterUpsert) SetCreatedAt(v time.Time) *ComicChapterUpsert {
	u.Set(comicchapter.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateCreatedAt() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ComicChapterUpsert) SetUpdatedAt(v time.Time) *ComicChapterUpsert {
	u.Set(comicchapter.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateUpdatedAt() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldUpdatedAt)
	return u
}

// SetChapter sets the "chapter" field.
func (u *ComicChapterUpsert) SetChapter(v uint) *ComicChapterUpsert {
	u.Set(comicchapter.FieldChapter, v)
	return u
}

// UpdateChapter sets the "chapter" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateChapter() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldChapter)
	return u
}

// AddChapter adds v to the "chapter" field.
func (u *ComicChapterUpsert) AddChapter(v uint) *ComicChapterUpsert {
	u.Add(comicchapter.FieldChapter, v)
	return u
}

// SetTitle sets the "title" field.
func (u *ComicChapterUpsert) SetTitle(v string) *ComicChapterUpsert {
	u.Set(comicchapter.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateTitle() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldTitle)
	return u
}

// ClearTitle clears the value of the "title" field.
func (u *ComicChapterUpsert) ClearTitle() *ComicChapterUpsert {
	u.SetNull(comicchapter.FieldTitle)
	return u
}

// SetVolumn sets the "volumn" field.
func (u *ComicChapterUpsert) SetVolumn(v string) *ComicChapterUpsert {
	u.Set(comicchapter.FieldVolumn, v)
	return u
}

// UpdateVolumn sets the "volumn" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateVolumn() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldVolumn)
	return u
}

// ClearVolumn clears the value of the "volumn" field.
func (u *ComicChapterUpsert) ClearVolumn() *ComicChapterUpsert {
	u.SetNull(comicchapter.FieldVolumn)
	return u
}

// SetLang sets the "lang" field.
func (u *ComicChapterUpsert) SetLang(v string) *ComicChapterUpsert {
	u.Set(comicchapter.FieldLang, v)
	return u
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateLang() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldLang)
	return u
}

// SetUpCount sets the "up_count" field.
func (u *ComicChapterUpsert) SetUpCount(v uint) *ComicChapterUpsert {
	u.Set(comicchapter.FieldUpCount, v)
	return u
}

// UpdateUpCount sets the "up_count" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateUpCount() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldUpCount)
	return u
}

// AddUpCount adds v to the "up_count" field.
func (u *ComicChapterUpsert) AddUpCount(v uint) *ComicChapterUpsert {
	u.Add(comicchapter.FieldUpCount, v)
	return u
}

// SetDownCount sets the "down_count" field.
func (u *ComicChapterUpsert) SetDownCount(v uint) *ComicChapterUpsert {
	u.Set(comicchapter.FieldDownCount, v)
	return u
}

// UpdateDownCount sets the "down_count" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateDownCount() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldDownCount)
	return u
}

// AddDownCount adds v to the "down_count" field.
func (u *ComicChapterUpsert) AddDownCount(v uint) *ComicChapterUpsert {
	u.Add(comicchapter.FieldDownCount, v)
	return u
}

// SetIsLastChapter sets the "is_last_chapter" field.
func (u *ComicChapterUpsert) SetIsLastChapter(v bool) *ComicChapterUpsert {
	u.Set(comicchapter.FieldIsLastChapter, v)
	return u
}

// UpdateIsLastChapter sets the "is_last_chapter" field to the value that was provided on create.
func (u *ComicChapterUpsert) UpdateIsLastChapter() *ComicChapterUpsert {
	u.SetExcluded(comicchapter.FieldIsLastChapter)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ComicChapter.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(comicchapter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ComicChapterUpsertOne) UpdateNewValues() *ComicChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(comicchapter.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ComicChapter.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ComicChapterUpsertOne) Ignore() *ComicChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ComicChapterUpsertOne) DoNothing() *ComicChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ComicChapterCreate.OnConflict
// documentation for more info.
func (u *ComicChapterUpsertOne) Update(set func(*ComicChapterUpsert)) *ComicChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ComicChapterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ComicChapterUpsertOne) SetCreatedAt(v time.Time) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateCreatedAt() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ComicChapterUpsertOne) SetUpdatedAt(v time.Time) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateUpdatedAt() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetChapter sets the "chapter" field.
func (u *ComicChapterUpsertOne) SetChapter(v uint) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetChapter(v)
	})
}

// AddChapter adds v to the "chapter" field.
func (u *ComicChapterUpsertOne) AddChapter(v uint) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.AddChapter(v)
	})
}

// UpdateChapter sets the "chapter" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateChapter() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateChapter()
	})
}

// SetTitle sets the "title" field.
func (u *ComicChapterUpsertOne) SetTitle(v string) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateTitle() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *ComicChapterUpsertOne) ClearTitle() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.ClearTitle()
	})
}

// SetVolumn sets the "volumn" field.
func (u *ComicChapterUpsertOne) SetVolumn(v string) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetVolumn(v)
	})
}

// UpdateVolumn sets the "volumn" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateVolumn() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateVolumn()
	})
}

// ClearVolumn clears the value of the "volumn" field.
func (u *ComicChapterUpsertOne) ClearVolumn() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.ClearVolumn()
	})
}

// SetLang sets the "lang" field.
func (u *ComicChapterUpsertOne) SetLang(v string) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateLang() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateLang()
	})
}

// SetUpCount sets the "up_count" field.
func (u *ComicChapterUpsertOne) SetUpCount(v uint) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetUpCount(v)
	})
}

// AddUpCount adds v to the "up_count" field.
func (u *ComicChapterUpsertOne) AddUpCount(v uint) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.AddUpCount(v)
	})
}

// UpdateUpCount sets the "up_count" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateUpCount() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateUpCount()
	})
}

// SetDownCount sets the "down_count" field.
func (u *ComicChapterUpsertOne) SetDownCount(v uint) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetDownCount(v)
	})
}

// AddDownCount adds v to the "down_count" field.
func (u *ComicChapterUpsertOne) AddDownCount(v uint) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.AddDownCount(v)
	})
}

// UpdateDownCount sets the "down_count" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateDownCount() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateDownCount()
	})
}

// SetIsLastChapter sets the "is_last_chapter" field.
func (u *ComicChapterUpsertOne) SetIsLastChapter(v bool) *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetIsLastChapter(v)
	})
}

// UpdateIsLastChapter sets the "is_last_chapter" field to the value that was provided on create.
func (u *ComicChapterUpsertOne) UpdateIsLastChapter() *ComicChapterUpsertOne {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateIsLastChapter()
	})
}

// Exec executes the query.
func (u *ComicChapterUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ComicChapterCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ComicChapterUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ComicChapterUpsertOne) ID(ctx context.Context) (id pulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ComicChapterUpsertOne.ID is not supported by MySQL driver. Use ComicChapterUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ComicChapterUpsertOne) IDX(ctx context.Context) pulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ComicChapterCreateBulk is the builder for creating many ComicChapter entities in bulk.
type ComicChapterCreateBulk struct {
	config
	err      error
	builders []*ComicChapterCreate
	conflict []sql.ConflictOption
}

// Save creates the ComicChapter entities in the database.
func (cccb *ComicChapterCreateBulk) Save(ctx context.Context) ([]*ComicChapter, error) {
	if cccb.err != nil {
		return nil, cccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*ComicChapter, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ComicChapterMutation)
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
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *ComicChapterCreateBulk) SaveX(ctx context.Context) []*ComicChapter {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *ComicChapterCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *ComicChapterCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ComicChapter.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ComicChapterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cccb *ComicChapterCreateBulk) OnConflict(opts ...sql.ConflictOption) *ComicChapterUpsertBulk {
	cccb.conflict = opts
	return &ComicChapterUpsertBulk{
		create: cccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ComicChapter.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cccb *ComicChapterCreateBulk) OnConflictColumns(columns ...string) *ComicChapterUpsertBulk {
	cccb.conflict = append(cccb.conflict, sql.ConflictColumns(columns...))
	return &ComicChapterUpsertBulk{
		create: cccb,
	}
}

// ComicChapterUpsertBulk is the builder for "upsert"-ing
// a bulk of ComicChapter nodes.
type ComicChapterUpsertBulk struct {
	create *ComicChapterCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ComicChapter.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(comicchapter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ComicChapterUpsertBulk) UpdateNewValues() *ComicChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(comicchapter.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ComicChapter.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ComicChapterUpsertBulk) Ignore() *ComicChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ComicChapterUpsertBulk) DoNothing() *ComicChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ComicChapterCreateBulk.OnConflict
// documentation for more info.
func (u *ComicChapterUpsertBulk) Update(set func(*ComicChapterUpsert)) *ComicChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ComicChapterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ComicChapterUpsertBulk) SetCreatedAt(v time.Time) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateCreatedAt() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ComicChapterUpsertBulk) SetUpdatedAt(v time.Time) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateUpdatedAt() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetChapter sets the "chapter" field.
func (u *ComicChapterUpsertBulk) SetChapter(v uint) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetChapter(v)
	})
}

// AddChapter adds v to the "chapter" field.
func (u *ComicChapterUpsertBulk) AddChapter(v uint) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.AddChapter(v)
	})
}

// UpdateChapter sets the "chapter" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateChapter() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateChapter()
	})
}

// SetTitle sets the "title" field.
func (u *ComicChapterUpsertBulk) SetTitle(v string) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateTitle() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *ComicChapterUpsertBulk) ClearTitle() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.ClearTitle()
	})
}

// SetVolumn sets the "volumn" field.
func (u *ComicChapterUpsertBulk) SetVolumn(v string) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetVolumn(v)
	})
}

// UpdateVolumn sets the "volumn" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateVolumn() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateVolumn()
	})
}

// ClearVolumn clears the value of the "volumn" field.
func (u *ComicChapterUpsertBulk) ClearVolumn() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.ClearVolumn()
	})
}

// SetLang sets the "lang" field.
func (u *ComicChapterUpsertBulk) SetLang(v string) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateLang() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateLang()
	})
}

// SetUpCount sets the "up_count" field.
func (u *ComicChapterUpsertBulk) SetUpCount(v uint) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetUpCount(v)
	})
}

// AddUpCount adds v to the "up_count" field.
func (u *ComicChapterUpsertBulk) AddUpCount(v uint) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.AddUpCount(v)
	})
}

// UpdateUpCount sets the "up_count" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateUpCount() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateUpCount()
	})
}

// SetDownCount sets the "down_count" field.
func (u *ComicChapterUpsertBulk) SetDownCount(v uint) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetDownCount(v)
	})
}

// AddDownCount adds v to the "down_count" field.
func (u *ComicChapterUpsertBulk) AddDownCount(v uint) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.AddDownCount(v)
	})
}

// UpdateDownCount sets the "down_count" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateDownCount() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateDownCount()
	})
}

// SetIsLastChapter sets the "is_last_chapter" field.
func (u *ComicChapterUpsertBulk) SetIsLastChapter(v bool) *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.SetIsLastChapter(v)
	})
}

// UpdateIsLastChapter sets the "is_last_chapter" field to the value that was provided on create.
func (u *ComicChapterUpsertBulk) UpdateIsLastChapter() *ComicChapterUpsertBulk {
	return u.Update(func(s *ComicChapterUpsert) {
		s.UpdateIsLastChapter()
	})
}

// Exec executes the query.
func (u *ComicChapterUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ComicChapterCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ComicChapterCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ComicChapterUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
