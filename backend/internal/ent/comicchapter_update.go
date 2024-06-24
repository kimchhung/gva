// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/backend/internal/ent/comic"
	"github.com/kimchhung/gva/backend/internal/ent/comicchapter"
	"github.com/kimchhung/gva/backend/internal/ent/comicimg"
	"github.com/kimchhung/gva/backend/internal/ent/predicate"

	"github.com/kimchhung/gva/backend/internal/ent/internal"
)

// ComicChapterUpdate is the builder for updating ComicChapter entities.
type ComicChapterUpdate struct {
	config
	hooks     []Hook
	mutation  *ComicChapterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ComicChapterUpdate builder.
func (ccu *ComicChapterUpdate) Where(ps ...predicate.ComicChapter) *ComicChapterUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *ComicChapterUpdate) SetCreatedAt(t time.Time) *ComicChapterUpdate {
	ccu.mutation.SetCreatedAt(t)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableCreatedAt(t *time.Time) *ComicChapterUpdate {
	if t != nil {
		ccu.SetCreatedAt(*t)
	}
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *ComicChapterUpdate) SetUpdatedAt(t time.Time) *ComicChapterUpdate {
	ccu.mutation.SetUpdatedAt(t)
	return ccu
}

// SetChapter sets the "chapter" field.
func (ccu *ComicChapterUpdate) SetChapter(u uint) *ComicChapterUpdate {
	ccu.mutation.ResetChapter()
	ccu.mutation.SetChapter(u)
	return ccu
}

// SetNillableChapter sets the "chapter" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableChapter(u *uint) *ComicChapterUpdate {
	if u != nil {
		ccu.SetChapter(*u)
	}
	return ccu
}

// AddChapter adds u to the "chapter" field.
func (ccu *ComicChapterUpdate) AddChapter(u int) *ComicChapterUpdate {
	ccu.mutation.AddChapter(u)
	return ccu
}

// SetTitle sets the "title" field.
func (ccu *ComicChapterUpdate) SetTitle(s string) *ComicChapterUpdate {
	ccu.mutation.SetTitle(s)
	return ccu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableTitle(s *string) *ComicChapterUpdate {
	if s != nil {
		ccu.SetTitle(*s)
	}
	return ccu
}

// ClearTitle clears the value of the "title" field.
func (ccu *ComicChapterUpdate) ClearTitle() *ComicChapterUpdate {
	ccu.mutation.ClearTitle()
	return ccu
}

// SetVolumn sets the "volumn" field.
func (ccu *ComicChapterUpdate) SetVolumn(s string) *ComicChapterUpdate {
	ccu.mutation.SetVolumn(s)
	return ccu
}

// SetNillableVolumn sets the "volumn" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableVolumn(s *string) *ComicChapterUpdate {
	if s != nil {
		ccu.SetVolumn(*s)
	}
	return ccu
}

// ClearVolumn clears the value of the "volumn" field.
func (ccu *ComicChapterUpdate) ClearVolumn() *ComicChapterUpdate {
	ccu.mutation.ClearVolumn()
	return ccu
}

// SetLang sets the "lang" field.
func (ccu *ComicChapterUpdate) SetLang(s string) *ComicChapterUpdate {
	ccu.mutation.SetLang(s)
	return ccu
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableLang(s *string) *ComicChapterUpdate {
	if s != nil {
		ccu.SetLang(*s)
	}
	return ccu
}

// SetUpCount sets the "up_count" field.
func (ccu *ComicChapterUpdate) SetUpCount(u uint) *ComicChapterUpdate {
	ccu.mutation.ResetUpCount()
	ccu.mutation.SetUpCount(u)
	return ccu
}

// SetNillableUpCount sets the "up_count" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableUpCount(u *uint) *ComicChapterUpdate {
	if u != nil {
		ccu.SetUpCount(*u)
	}
	return ccu
}

// AddUpCount adds u to the "up_count" field.
func (ccu *ComicChapterUpdate) AddUpCount(u int) *ComicChapterUpdate {
	ccu.mutation.AddUpCount(u)
	return ccu
}

// SetDownCount sets the "down_count" field.
func (ccu *ComicChapterUpdate) SetDownCount(u uint) *ComicChapterUpdate {
	ccu.mutation.ResetDownCount()
	ccu.mutation.SetDownCount(u)
	return ccu
}

// SetNillableDownCount sets the "down_count" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableDownCount(u *uint) *ComicChapterUpdate {
	if u != nil {
		ccu.SetDownCount(*u)
	}
	return ccu
}

// AddDownCount adds u to the "down_count" field.
func (ccu *ComicChapterUpdate) AddDownCount(u int) *ComicChapterUpdate {
	ccu.mutation.AddDownCount(u)
	return ccu
}

// SetIsLastChapter sets the "is_last_chapter" field.
func (ccu *ComicChapterUpdate) SetIsLastChapter(b bool) *ComicChapterUpdate {
	ccu.mutation.SetIsLastChapter(b)
	return ccu
}

// SetNillableIsLastChapter sets the "is_last_chapter" field if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableIsLastChapter(b *bool) *ComicChapterUpdate {
	if b != nil {
		ccu.SetIsLastChapter(*b)
	}
	return ccu
}

// AddImgIDs adds the "imgs" edge to the ComicImg entity by IDs.
func (ccu *ComicChapterUpdate) AddImgIDs(ids ...string) *ComicChapterUpdate {
	ccu.mutation.AddImgIDs(ids...)
	return ccu
}

// AddImgs adds the "imgs" edges to the ComicImg entity.
func (ccu *ComicChapterUpdate) AddImgs(c ...*ComicImg) *ComicChapterUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccu.AddImgIDs(ids...)
}

// SetComicID sets the "comic" edge to the Comic entity by ID.
func (ccu *ComicChapterUpdate) SetComicID(id string) *ComicChapterUpdate {
	ccu.mutation.SetComicID(id)
	return ccu
}

// SetNillableComicID sets the "comic" edge to the Comic entity by ID if the given value is not nil.
func (ccu *ComicChapterUpdate) SetNillableComicID(id *string) *ComicChapterUpdate {
	if id != nil {
		ccu = ccu.SetComicID(*id)
	}
	return ccu
}

// SetComic sets the "comic" edge to the Comic entity.
func (ccu *ComicChapterUpdate) SetComic(c *Comic) *ComicChapterUpdate {
	return ccu.SetComicID(c.ID)
}

// Mutation returns the ComicChapterMutation object of the builder.
func (ccu *ComicChapterUpdate) Mutation() *ComicChapterMutation {
	return ccu.mutation
}

// ClearImgs clears all "imgs" edges to the ComicImg entity.
func (ccu *ComicChapterUpdate) ClearImgs() *ComicChapterUpdate {
	ccu.mutation.ClearImgs()
	return ccu
}

// RemoveImgIDs removes the "imgs" edge to ComicImg entities by IDs.
func (ccu *ComicChapterUpdate) RemoveImgIDs(ids ...string) *ComicChapterUpdate {
	ccu.mutation.RemoveImgIDs(ids...)
	return ccu
}

// RemoveImgs removes "imgs" edges to ComicImg entities.
func (ccu *ComicChapterUpdate) RemoveImgs(c ...*ComicImg) *ComicChapterUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccu.RemoveImgIDs(ids...)
}

// ClearComic clears the "comic" edge to the Comic entity.
func (ccu *ComicChapterUpdate) ClearComic() *ComicChapterUpdate {
	ccu.mutation.ClearComic()
	return ccu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *ComicChapterUpdate) Save(ctx context.Context) (int, error) {
	ccu.defaults()
	return withHooks(ctx, ccu.sqlSave, ccu.mutation, ccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *ComicChapterUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *ComicChapterUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *ComicChapterUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccu *ComicChapterUpdate) defaults() {
	if _, ok := ccu.mutation.UpdatedAt(); !ok {
		v := comicchapter.UpdateDefaultUpdatedAt()
		ccu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ccu *ComicChapterUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ComicChapterUpdate {
	ccu.modifiers = append(ccu.modifiers, modifiers...)
	return ccu
}

func (ccu *ComicChapterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(comicchapter.Table, comicchapter.Columns, sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString))
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.SetField(comicchapter.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.SetField(comicchapter.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ccu.mutation.Chapter(); ok {
		_spec.SetField(comicchapter.FieldChapter, field.TypeUint, value)
	}
	if value, ok := ccu.mutation.AddedChapter(); ok {
		_spec.AddField(comicchapter.FieldChapter, field.TypeUint, value)
	}
	if value, ok := ccu.mutation.Title(); ok {
		_spec.SetField(comicchapter.FieldTitle, field.TypeString, value)
	}
	if ccu.mutation.TitleCleared() {
		_spec.ClearField(comicchapter.FieldTitle, field.TypeString)
	}
	if value, ok := ccu.mutation.Volumn(); ok {
		_spec.SetField(comicchapter.FieldVolumn, field.TypeString, value)
	}
	if ccu.mutation.VolumnCleared() {
		_spec.ClearField(comicchapter.FieldVolumn, field.TypeString)
	}
	if value, ok := ccu.mutation.Lang(); ok {
		_spec.SetField(comicchapter.FieldLang, field.TypeString, value)
	}
	if value, ok := ccu.mutation.UpCount(); ok {
		_spec.SetField(comicchapter.FieldUpCount, field.TypeUint, value)
	}
	if value, ok := ccu.mutation.AddedUpCount(); ok {
		_spec.AddField(comicchapter.FieldUpCount, field.TypeUint, value)
	}
	if value, ok := ccu.mutation.DownCount(); ok {
		_spec.SetField(comicchapter.FieldDownCount, field.TypeUint, value)
	}
	if value, ok := ccu.mutation.AddedDownCount(); ok {
		_spec.AddField(comicchapter.FieldDownCount, field.TypeUint, value)
	}
	if value, ok := ccu.mutation.IsLastChapter(); ok {
		_spec.SetField(comicchapter.FieldIsLastChapter, field.TypeBool, value)
	}
	if ccu.mutation.ImgsCleared() {
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
		edge.Schema = ccu.schemaConfig.ComicImg
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.RemovedImgsIDs(); len(nodes) > 0 && !ccu.mutation.ImgsCleared() {
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
		edge.Schema = ccu.schemaConfig.ComicImg
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.ImgsIDs(); len(nodes) > 0 {
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
		edge.Schema = ccu.schemaConfig.ComicImg
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ccu.mutation.ComicCleared() {
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
		edge.Schema = ccu.schemaConfig.ComicChapter
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.ComicIDs(); len(nodes) > 0 {
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
		edge.Schema = ccu.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = ccu.schemaConfig.ComicChapter
	ctx = internal.NewSchemaConfigContext(ctx, ccu.schemaConfig)
	_spec.AddModifiers(ccu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comicchapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccu.mutation.done = true
	return n, nil
}

// ComicChapterUpdateOne is the builder for updating a single ComicChapter entity.
type ComicChapterUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ComicChapterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *ComicChapterUpdateOne) SetCreatedAt(t time.Time) *ComicChapterUpdateOne {
	ccuo.mutation.SetCreatedAt(t)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableCreatedAt(t *time.Time) *ComicChapterUpdateOne {
	if t != nil {
		ccuo.SetCreatedAt(*t)
	}
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *ComicChapterUpdateOne) SetUpdatedAt(t time.Time) *ComicChapterUpdateOne {
	ccuo.mutation.SetUpdatedAt(t)
	return ccuo
}

// SetChapter sets the "chapter" field.
func (ccuo *ComicChapterUpdateOne) SetChapter(u uint) *ComicChapterUpdateOne {
	ccuo.mutation.ResetChapter()
	ccuo.mutation.SetChapter(u)
	return ccuo
}

// SetNillableChapter sets the "chapter" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableChapter(u *uint) *ComicChapterUpdateOne {
	if u != nil {
		ccuo.SetChapter(*u)
	}
	return ccuo
}

// AddChapter adds u to the "chapter" field.
func (ccuo *ComicChapterUpdateOne) AddChapter(u int) *ComicChapterUpdateOne {
	ccuo.mutation.AddChapter(u)
	return ccuo
}

// SetTitle sets the "title" field.
func (ccuo *ComicChapterUpdateOne) SetTitle(s string) *ComicChapterUpdateOne {
	ccuo.mutation.SetTitle(s)
	return ccuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableTitle(s *string) *ComicChapterUpdateOne {
	if s != nil {
		ccuo.SetTitle(*s)
	}
	return ccuo
}

// ClearTitle clears the value of the "title" field.
func (ccuo *ComicChapterUpdateOne) ClearTitle() *ComicChapterUpdateOne {
	ccuo.mutation.ClearTitle()
	return ccuo
}

// SetVolumn sets the "volumn" field.
func (ccuo *ComicChapterUpdateOne) SetVolumn(s string) *ComicChapterUpdateOne {
	ccuo.mutation.SetVolumn(s)
	return ccuo
}

// SetNillableVolumn sets the "volumn" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableVolumn(s *string) *ComicChapterUpdateOne {
	if s != nil {
		ccuo.SetVolumn(*s)
	}
	return ccuo
}

// ClearVolumn clears the value of the "volumn" field.
func (ccuo *ComicChapterUpdateOne) ClearVolumn() *ComicChapterUpdateOne {
	ccuo.mutation.ClearVolumn()
	return ccuo
}

// SetLang sets the "lang" field.
func (ccuo *ComicChapterUpdateOne) SetLang(s string) *ComicChapterUpdateOne {
	ccuo.mutation.SetLang(s)
	return ccuo
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableLang(s *string) *ComicChapterUpdateOne {
	if s != nil {
		ccuo.SetLang(*s)
	}
	return ccuo
}

// SetUpCount sets the "up_count" field.
func (ccuo *ComicChapterUpdateOne) SetUpCount(u uint) *ComicChapterUpdateOne {
	ccuo.mutation.ResetUpCount()
	ccuo.mutation.SetUpCount(u)
	return ccuo
}

// SetNillableUpCount sets the "up_count" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableUpCount(u *uint) *ComicChapterUpdateOne {
	if u != nil {
		ccuo.SetUpCount(*u)
	}
	return ccuo
}

// AddUpCount adds u to the "up_count" field.
func (ccuo *ComicChapterUpdateOne) AddUpCount(u int) *ComicChapterUpdateOne {
	ccuo.mutation.AddUpCount(u)
	return ccuo
}

// SetDownCount sets the "down_count" field.
func (ccuo *ComicChapterUpdateOne) SetDownCount(u uint) *ComicChapterUpdateOne {
	ccuo.mutation.ResetDownCount()
	ccuo.mutation.SetDownCount(u)
	return ccuo
}

// SetNillableDownCount sets the "down_count" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableDownCount(u *uint) *ComicChapterUpdateOne {
	if u != nil {
		ccuo.SetDownCount(*u)
	}
	return ccuo
}

// AddDownCount adds u to the "down_count" field.
func (ccuo *ComicChapterUpdateOne) AddDownCount(u int) *ComicChapterUpdateOne {
	ccuo.mutation.AddDownCount(u)
	return ccuo
}

// SetIsLastChapter sets the "is_last_chapter" field.
func (ccuo *ComicChapterUpdateOne) SetIsLastChapter(b bool) *ComicChapterUpdateOne {
	ccuo.mutation.SetIsLastChapter(b)
	return ccuo
}

// SetNillableIsLastChapter sets the "is_last_chapter" field if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableIsLastChapter(b *bool) *ComicChapterUpdateOne {
	if b != nil {
		ccuo.SetIsLastChapter(*b)
	}
	return ccuo
}

// AddImgIDs adds the "imgs" edge to the ComicImg entity by IDs.
func (ccuo *ComicChapterUpdateOne) AddImgIDs(ids ...string) *ComicChapterUpdateOne {
	ccuo.mutation.AddImgIDs(ids...)
	return ccuo
}

// AddImgs adds the "imgs" edges to the ComicImg entity.
func (ccuo *ComicChapterUpdateOne) AddImgs(c ...*ComicImg) *ComicChapterUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccuo.AddImgIDs(ids...)
}

// SetComicID sets the "comic" edge to the Comic entity by ID.
func (ccuo *ComicChapterUpdateOne) SetComicID(id string) *ComicChapterUpdateOne {
	ccuo.mutation.SetComicID(id)
	return ccuo
}

// SetNillableComicID sets the "comic" edge to the Comic entity by ID if the given value is not nil.
func (ccuo *ComicChapterUpdateOne) SetNillableComicID(id *string) *ComicChapterUpdateOne {
	if id != nil {
		ccuo = ccuo.SetComicID(*id)
	}
	return ccuo
}

// SetComic sets the "comic" edge to the Comic entity.
func (ccuo *ComicChapterUpdateOne) SetComic(c *Comic) *ComicChapterUpdateOne {
	return ccuo.SetComicID(c.ID)
}

// Mutation returns the ComicChapterMutation object of the builder.
func (ccuo *ComicChapterUpdateOne) Mutation() *ComicChapterMutation {
	return ccuo.mutation
}

// ClearImgs clears all "imgs" edges to the ComicImg entity.
func (ccuo *ComicChapterUpdateOne) ClearImgs() *ComicChapterUpdateOne {
	ccuo.mutation.ClearImgs()
	return ccuo
}

// RemoveImgIDs removes the "imgs" edge to ComicImg entities by IDs.
func (ccuo *ComicChapterUpdateOne) RemoveImgIDs(ids ...string) *ComicChapterUpdateOne {
	ccuo.mutation.RemoveImgIDs(ids...)
	return ccuo
}

// RemoveImgs removes "imgs" edges to ComicImg entities.
func (ccuo *ComicChapterUpdateOne) RemoveImgs(c ...*ComicImg) *ComicChapterUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccuo.RemoveImgIDs(ids...)
}

// ClearComic clears the "comic" edge to the Comic entity.
func (ccuo *ComicChapterUpdateOne) ClearComic() *ComicChapterUpdateOne {
	ccuo.mutation.ClearComic()
	return ccuo
}

// Where appends a list predicates to the ComicChapterUpdate builder.
func (ccuo *ComicChapterUpdateOne) Where(ps ...predicate.ComicChapter) *ComicChapterUpdateOne {
	ccuo.mutation.Where(ps...)
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *ComicChapterUpdateOne) Select(field string, fields ...string) *ComicChapterUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated ComicChapter entity.
func (ccuo *ComicChapterUpdateOne) Save(ctx context.Context) (*ComicChapter, error) {
	ccuo.defaults()
	return withHooks(ctx, ccuo.sqlSave, ccuo.mutation, ccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *ComicChapterUpdateOne) SaveX(ctx context.Context) *ComicChapter {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *ComicChapterUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *ComicChapterUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccuo *ComicChapterUpdateOne) defaults() {
	if _, ok := ccuo.mutation.UpdatedAt(); !ok {
		v := comicchapter.UpdateDefaultUpdatedAt()
		ccuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ccuo *ComicChapterUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ComicChapterUpdateOne {
	ccuo.modifiers = append(ccuo.modifiers, modifiers...)
	return ccuo
}

func (ccuo *ComicChapterUpdateOne) sqlSave(ctx context.Context) (_node *ComicChapter, err error) {
	_spec := sqlgraph.NewUpdateSpec(comicchapter.Table, comicchapter.Columns, sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString))
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ComicChapter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comicchapter.FieldID)
		for _, f := range fields {
			if !comicchapter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comicchapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.SetField(comicchapter.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comicchapter.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ccuo.mutation.Chapter(); ok {
		_spec.SetField(comicchapter.FieldChapter, field.TypeUint, value)
	}
	if value, ok := ccuo.mutation.AddedChapter(); ok {
		_spec.AddField(comicchapter.FieldChapter, field.TypeUint, value)
	}
	if value, ok := ccuo.mutation.Title(); ok {
		_spec.SetField(comicchapter.FieldTitle, field.TypeString, value)
	}
	if ccuo.mutation.TitleCleared() {
		_spec.ClearField(comicchapter.FieldTitle, field.TypeString)
	}
	if value, ok := ccuo.mutation.Volumn(); ok {
		_spec.SetField(comicchapter.FieldVolumn, field.TypeString, value)
	}
	if ccuo.mutation.VolumnCleared() {
		_spec.ClearField(comicchapter.FieldVolumn, field.TypeString)
	}
	if value, ok := ccuo.mutation.Lang(); ok {
		_spec.SetField(comicchapter.FieldLang, field.TypeString, value)
	}
	if value, ok := ccuo.mutation.UpCount(); ok {
		_spec.SetField(comicchapter.FieldUpCount, field.TypeUint, value)
	}
	if value, ok := ccuo.mutation.AddedUpCount(); ok {
		_spec.AddField(comicchapter.FieldUpCount, field.TypeUint, value)
	}
	if value, ok := ccuo.mutation.DownCount(); ok {
		_spec.SetField(comicchapter.FieldDownCount, field.TypeUint, value)
	}
	if value, ok := ccuo.mutation.AddedDownCount(); ok {
		_spec.AddField(comicchapter.FieldDownCount, field.TypeUint, value)
	}
	if value, ok := ccuo.mutation.IsLastChapter(); ok {
		_spec.SetField(comicchapter.FieldIsLastChapter, field.TypeBool, value)
	}
	if ccuo.mutation.ImgsCleared() {
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
		edge.Schema = ccuo.schemaConfig.ComicImg
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.RemovedImgsIDs(); len(nodes) > 0 && !ccuo.mutation.ImgsCleared() {
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
		edge.Schema = ccuo.schemaConfig.ComicImg
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.ImgsIDs(); len(nodes) > 0 {
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
		edge.Schema = ccuo.schemaConfig.ComicImg
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ccuo.mutation.ComicCleared() {
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
		edge.Schema = ccuo.schemaConfig.ComicChapter
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.ComicIDs(); len(nodes) > 0 {
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
		edge.Schema = ccuo.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = ccuo.schemaConfig.ComicChapter
	ctx = internal.NewSchemaConfigContext(ctx, ccuo.schemaConfig)
	_spec.AddModifiers(ccuo.modifiers...)
	_node = &ComicChapter{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comicchapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccuo.mutation.done = true
	return _node, nil
}
