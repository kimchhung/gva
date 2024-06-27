// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/types"
	"github.com/gva/internal/ent/comic"
	"github.com/gva/internal/ent/comicchapter"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ComicUpdate is the builder for updating Comic entities.
type ComicUpdate struct {
	config
	hooks     []Hook
	mutation  *ComicMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ComicUpdate builder.
func (cu *ComicUpdate) Where(ps ...predicate.Comic) *ComicUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *ComicUpdate) SetCreatedAt(t time.Time) *ComicUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableCreatedAt(t *time.Time) *ComicUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ComicUpdate) SetUpdatedAt(t time.Time) *ComicUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetChapter sets the "chapter" field.
func (cu *ComicUpdate) SetChapter(u uint) *ComicUpdate {
	cu.mutation.ResetChapter()
	cu.mutation.SetChapter(u)
	return cu
}

// SetNillableChapter sets the "chapter" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableChapter(u *uint) *ComicUpdate {
	if u != nil {
		cu.SetChapter(*u)
	}
	return cu
}

// AddChapter adds u to the "chapter" field.
func (cu *ComicUpdate) AddChapter(u int) *ComicUpdate {
	cu.mutation.AddChapter(u)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ComicUpdate) SetTitle(s string) *ComicUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableTitle(s *string) *ComicUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// SetSlug sets the "slug" field.
func (cu *ComicUpdate) SetSlug(s string) *ComicUpdate {
	cu.mutation.SetSlug(s)
	return cu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableSlug(s *string) *ComicUpdate {
	if s != nil {
		cu.SetSlug(*s)
	}
	return cu
}

// SetCovers sets the "covers" field.
func (cu *ComicUpdate) SetCovers(ti []types.CoverImg) *ComicUpdate {
	cu.mutation.SetCovers(ti)
	return cu
}

// AppendCovers appends ti to the "covers" field.
func (cu *ComicUpdate) AppendCovers(ti []types.CoverImg) *ComicUpdate {
	cu.mutation.AppendCovers(ti)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ComicUpdate) SetStatus(s string) *ComicUpdate {
	cu.mutation.SetStatus(s)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableStatus(s *string) *ComicUpdate {
	if s != nil {
		cu.SetStatus(*s)
	}
	return cu
}

// SetIsTranslateCompleted sets the "isTranslateCompleted" field.
func (cu *ComicUpdate) SetIsTranslateCompleted(b bool) *ComicUpdate {
	cu.mutation.SetIsTranslateCompleted(b)
	return cu
}

// SetNillableIsTranslateCompleted sets the "isTranslateCompleted" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableIsTranslateCompleted(b *bool) *ComicUpdate {
	if b != nil {
		cu.SetIsTranslateCompleted(*b)
	}
	return cu
}

// SetUpCount sets the "up_count" field.
func (cu *ComicUpdate) SetUpCount(u uint) *ComicUpdate {
	cu.mutation.ResetUpCount()
	cu.mutation.SetUpCount(u)
	return cu
}

// SetNillableUpCount sets the "up_count" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableUpCount(u *uint) *ComicUpdate {
	if u != nil {
		cu.SetUpCount(*u)
	}
	return cu
}

// AddUpCount adds u to the "up_count" field.
func (cu *ComicUpdate) AddUpCount(u int) *ComicUpdate {
	cu.mutation.AddUpCount(u)
	return cu
}

// SetFinalChapterID sets the "final_chapter_id" field.
func (cu *ComicUpdate) SetFinalChapterID(s string) *ComicUpdate {
	cu.mutation.SetFinalChapterID(s)
	return cu
}

// SetNillableFinalChapterID sets the "final_chapter_id" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableFinalChapterID(s *string) *ComicUpdate {
	if s != nil {
		cu.SetFinalChapterID(*s)
	}
	return cu
}

// ClearFinalChapterID clears the value of the "final_chapter_id" field.
func (cu *ComicUpdate) ClearFinalChapterID() *ComicUpdate {
	cu.mutation.ClearFinalChapterID()
	return cu
}

// SetLastChapterID sets the "last_chapter_id" field.
func (cu *ComicUpdate) SetLastChapterID(s string) *ComicUpdate {
	cu.mutation.SetLastChapterID(s)
	return cu
}

// SetNillableLastChapterID sets the "last_chapter_id" field if the given value is not nil.
func (cu *ComicUpdate) SetNillableLastChapterID(s *string) *ComicUpdate {
	if s != nil {
		cu.SetLastChapterID(*s)
	}
	return cu
}

// ClearLastChapterID clears the value of the "last_chapter_id" field.
func (cu *ComicUpdate) ClearLastChapterID() *ComicUpdate {
	cu.mutation.ClearLastChapterID()
	return cu
}

// AddChapterIDs adds the "chapters" edge to the ComicChapter entity by IDs.
func (cu *ComicUpdate) AddChapterIDs(ids ...string) *ComicUpdate {
	cu.mutation.AddChapterIDs(ids...)
	return cu
}

// AddChapters adds the "chapters" edges to the ComicChapter entity.
func (cu *ComicUpdate) AddChapters(c ...*ComicChapter) *ComicUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChapterIDs(ids...)
}

// SetLastChapter sets the "last_chapter" edge to the ComicChapter entity.
func (cu *ComicUpdate) SetLastChapter(c *ComicChapter) *ComicUpdate {
	return cu.SetLastChapterID(c.ID)
}

// SetFinalChapter sets the "final_chapter" edge to the ComicChapter entity.
func (cu *ComicUpdate) SetFinalChapter(c *ComicChapter) *ComicUpdate {
	return cu.SetFinalChapterID(c.ID)
}

// Mutation returns the ComicMutation object of the builder.
func (cu *ComicUpdate) Mutation() *ComicMutation {
	return cu.mutation
}

// ClearChapters clears all "chapters" edges to the ComicChapter entity.
func (cu *ComicUpdate) ClearChapters() *ComicUpdate {
	cu.mutation.ClearChapters()
	return cu
}

// RemoveChapterIDs removes the "chapters" edge to ComicChapter entities by IDs.
func (cu *ComicUpdate) RemoveChapterIDs(ids ...string) *ComicUpdate {
	cu.mutation.RemoveChapterIDs(ids...)
	return cu
}

// RemoveChapters removes "chapters" edges to ComicChapter entities.
func (cu *ComicUpdate) RemoveChapters(c ...*ComicChapter) *ComicUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChapterIDs(ids...)
}

// ClearLastChapter clears the "last_chapter" edge to the ComicChapter entity.
func (cu *ComicUpdate) ClearLastChapter() *ComicUpdate {
	cu.mutation.ClearLastChapter()
	return cu
}

// ClearFinalChapter clears the "final_chapter" edge to the ComicChapter entity.
func (cu *ComicUpdate) ClearFinalChapter() *ComicUpdate {
	cu.mutation.ClearFinalChapter()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ComicUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ComicUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ComicUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ComicUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ComicUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := comic.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ComicUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ComicUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ComicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(comic.Table, comic.Columns, sqlgraph.NewFieldSpec(comic.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(comic.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Chapter(); ok {
		_spec.SetField(comic.FieldChapter, field.TypeUint, value)
	}
	if value, ok := cu.mutation.AddedChapter(); ok {
		_spec.AddField(comic.FieldChapter, field.TypeUint, value)
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(comic.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Slug(); ok {
		_spec.SetField(comic.FieldSlug, field.TypeString, value)
	}
	if value, ok := cu.mutation.Covers(); ok {
		_spec.SetField(comic.FieldCovers, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedCovers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, comic.FieldCovers, value)
		})
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(comic.FieldStatus, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsTranslateCompleted(); ok {
		_spec.SetField(comic.FieldIsTranslateCompleted, field.TypeBool, value)
	}
	if value, ok := cu.mutation.UpCount(); ok {
		_spec.SetField(comic.FieldUpCount, field.TypeUint, value)
	}
	if value, ok := cu.mutation.AddedUpCount(); ok {
		_spec.AddField(comic.FieldUpCount, field.TypeUint, value)
	}
	if cu.mutation.ChaptersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comic.ChaptersTable,
			Columns: []string{comic.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.ComicChapter
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChaptersIDs(); len(nodes) > 0 && !cu.mutation.ChaptersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comic.ChaptersTable,
			Columns: []string{comic.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comic.ChaptersTable,
			Columns: []string{comic.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.LastChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.LastChapterTable,
			Columns: []string{comic.LastChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.Comic
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LastChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.LastChapterTable,
			Columns: []string{comic.LastChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.Comic
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.FinalChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.FinalChapterTable,
			Columns: []string{comic.FinalChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.Comic
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FinalChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.FinalChapterTable,
			Columns: []string{comic.FinalChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cu.schemaConfig.Comic
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = cu.schemaConfig.Comic
	ctx = internal.NewSchemaConfigContext(ctx, cu.schemaConfig)
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ComicUpdateOne is the builder for updating a single Comic entity.
type ComicUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ComicMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *ComicUpdateOne) SetCreatedAt(t time.Time) *ComicUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableCreatedAt(t *time.Time) *ComicUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ComicUpdateOne) SetUpdatedAt(t time.Time) *ComicUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetChapter sets the "chapter" field.
func (cuo *ComicUpdateOne) SetChapter(u uint) *ComicUpdateOne {
	cuo.mutation.ResetChapter()
	cuo.mutation.SetChapter(u)
	return cuo
}

// SetNillableChapter sets the "chapter" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableChapter(u *uint) *ComicUpdateOne {
	if u != nil {
		cuo.SetChapter(*u)
	}
	return cuo
}

// AddChapter adds u to the "chapter" field.
func (cuo *ComicUpdateOne) AddChapter(u int) *ComicUpdateOne {
	cuo.mutation.AddChapter(u)
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *ComicUpdateOne) SetTitle(s string) *ComicUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableTitle(s *string) *ComicUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// SetSlug sets the "slug" field.
func (cuo *ComicUpdateOne) SetSlug(s string) *ComicUpdateOne {
	cuo.mutation.SetSlug(s)
	return cuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableSlug(s *string) *ComicUpdateOne {
	if s != nil {
		cuo.SetSlug(*s)
	}
	return cuo
}

// SetCovers sets the "covers" field.
func (cuo *ComicUpdateOne) SetCovers(ti []types.CoverImg) *ComicUpdateOne {
	cuo.mutation.SetCovers(ti)
	return cuo
}

// AppendCovers appends ti to the "covers" field.
func (cuo *ComicUpdateOne) AppendCovers(ti []types.CoverImg) *ComicUpdateOne {
	cuo.mutation.AppendCovers(ti)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ComicUpdateOne) SetStatus(s string) *ComicUpdateOne {
	cuo.mutation.SetStatus(s)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableStatus(s *string) *ComicUpdateOne {
	if s != nil {
		cuo.SetStatus(*s)
	}
	return cuo
}

// SetIsTranslateCompleted sets the "isTranslateCompleted" field.
func (cuo *ComicUpdateOne) SetIsTranslateCompleted(b bool) *ComicUpdateOne {
	cuo.mutation.SetIsTranslateCompleted(b)
	return cuo
}

// SetNillableIsTranslateCompleted sets the "isTranslateCompleted" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableIsTranslateCompleted(b *bool) *ComicUpdateOne {
	if b != nil {
		cuo.SetIsTranslateCompleted(*b)
	}
	return cuo
}

// SetUpCount sets the "up_count" field.
func (cuo *ComicUpdateOne) SetUpCount(u uint) *ComicUpdateOne {
	cuo.mutation.ResetUpCount()
	cuo.mutation.SetUpCount(u)
	return cuo
}

// SetNillableUpCount sets the "up_count" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableUpCount(u *uint) *ComicUpdateOne {
	if u != nil {
		cuo.SetUpCount(*u)
	}
	return cuo
}

// AddUpCount adds u to the "up_count" field.
func (cuo *ComicUpdateOne) AddUpCount(u int) *ComicUpdateOne {
	cuo.mutation.AddUpCount(u)
	return cuo
}

// SetFinalChapterID sets the "final_chapter_id" field.
func (cuo *ComicUpdateOne) SetFinalChapterID(s string) *ComicUpdateOne {
	cuo.mutation.SetFinalChapterID(s)
	return cuo
}

// SetNillableFinalChapterID sets the "final_chapter_id" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableFinalChapterID(s *string) *ComicUpdateOne {
	if s != nil {
		cuo.SetFinalChapterID(*s)
	}
	return cuo
}

// ClearFinalChapterID clears the value of the "final_chapter_id" field.
func (cuo *ComicUpdateOne) ClearFinalChapterID() *ComicUpdateOne {
	cuo.mutation.ClearFinalChapterID()
	return cuo
}

// SetLastChapterID sets the "last_chapter_id" field.
func (cuo *ComicUpdateOne) SetLastChapterID(s string) *ComicUpdateOne {
	cuo.mutation.SetLastChapterID(s)
	return cuo
}

// SetNillableLastChapterID sets the "last_chapter_id" field if the given value is not nil.
func (cuo *ComicUpdateOne) SetNillableLastChapterID(s *string) *ComicUpdateOne {
	if s != nil {
		cuo.SetLastChapterID(*s)
	}
	return cuo
}

// ClearLastChapterID clears the value of the "last_chapter_id" field.
func (cuo *ComicUpdateOne) ClearLastChapterID() *ComicUpdateOne {
	cuo.mutation.ClearLastChapterID()
	return cuo
}

// AddChapterIDs adds the "chapters" edge to the ComicChapter entity by IDs.
func (cuo *ComicUpdateOne) AddChapterIDs(ids ...string) *ComicUpdateOne {
	cuo.mutation.AddChapterIDs(ids...)
	return cuo
}

// AddChapters adds the "chapters" edges to the ComicChapter entity.
func (cuo *ComicUpdateOne) AddChapters(c ...*ComicChapter) *ComicUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChapterIDs(ids...)
}

// SetLastChapter sets the "last_chapter" edge to the ComicChapter entity.
func (cuo *ComicUpdateOne) SetLastChapter(c *ComicChapter) *ComicUpdateOne {
	return cuo.SetLastChapterID(c.ID)
}

// SetFinalChapter sets the "final_chapter" edge to the ComicChapter entity.
func (cuo *ComicUpdateOne) SetFinalChapter(c *ComicChapter) *ComicUpdateOne {
	return cuo.SetFinalChapterID(c.ID)
}

// Mutation returns the ComicMutation object of the builder.
func (cuo *ComicUpdateOne) Mutation() *ComicMutation {
	return cuo.mutation
}

// ClearChapters clears all "chapters" edges to the ComicChapter entity.
func (cuo *ComicUpdateOne) ClearChapters() *ComicUpdateOne {
	cuo.mutation.ClearChapters()
	return cuo
}

// RemoveChapterIDs removes the "chapters" edge to ComicChapter entities by IDs.
func (cuo *ComicUpdateOne) RemoveChapterIDs(ids ...string) *ComicUpdateOne {
	cuo.mutation.RemoveChapterIDs(ids...)
	return cuo
}

// RemoveChapters removes "chapters" edges to ComicChapter entities.
func (cuo *ComicUpdateOne) RemoveChapters(c ...*ComicChapter) *ComicUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChapterIDs(ids...)
}

// ClearLastChapter clears the "last_chapter" edge to the ComicChapter entity.
func (cuo *ComicUpdateOne) ClearLastChapter() *ComicUpdateOne {
	cuo.mutation.ClearLastChapter()
	return cuo
}

// ClearFinalChapter clears the "final_chapter" edge to the ComicChapter entity.
func (cuo *ComicUpdateOne) ClearFinalChapter() *ComicUpdateOne {
	cuo.mutation.ClearFinalChapter()
	return cuo
}

// Where appends a list predicates to the ComicUpdate builder.
func (cuo *ComicUpdateOne) Where(ps ...predicate.Comic) *ComicUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ComicUpdateOne) Select(field string, fields ...string) *ComicUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comic entity.
func (cuo *ComicUpdateOne) Save(ctx context.Context) (*Comic, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ComicUpdateOne) SaveX(ctx context.Context) *Comic {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ComicUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ComicUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ComicUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := comic.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ComicUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ComicUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ComicUpdateOne) sqlSave(ctx context.Context) (_node *Comic, err error) {
	_spec := sqlgraph.NewUpdateSpec(comic.Table, comic.Columns, sqlgraph.NewFieldSpec(comic.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comic.FieldID)
		for _, f := range fields {
			if !comic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(comic.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Chapter(); ok {
		_spec.SetField(comic.FieldChapter, field.TypeUint, value)
	}
	if value, ok := cuo.mutation.AddedChapter(); ok {
		_spec.AddField(comic.FieldChapter, field.TypeUint, value)
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(comic.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Slug(); ok {
		_spec.SetField(comic.FieldSlug, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Covers(); ok {
		_spec.SetField(comic.FieldCovers, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedCovers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, comic.FieldCovers, value)
		})
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(comic.FieldStatus, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsTranslateCompleted(); ok {
		_spec.SetField(comic.FieldIsTranslateCompleted, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.UpCount(); ok {
		_spec.SetField(comic.FieldUpCount, field.TypeUint, value)
	}
	if value, ok := cuo.mutation.AddedUpCount(); ok {
		_spec.AddField(comic.FieldUpCount, field.TypeUint, value)
	}
	if cuo.mutation.ChaptersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comic.ChaptersTable,
			Columns: []string{comic.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.ComicChapter
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChaptersIDs(); len(nodes) > 0 && !cuo.mutation.ChaptersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comic.ChaptersTable,
			Columns: []string{comic.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comic.ChaptersTable,
			Columns: []string{comic.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.ComicChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.LastChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.LastChapterTable,
			Columns: []string{comic.LastChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.Comic
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LastChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.LastChapterTable,
			Columns: []string{comic.LastChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.Comic
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.FinalChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.FinalChapterTable,
			Columns: []string{comic.FinalChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.Comic
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FinalChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comic.FinalChapterTable,
			Columns: []string{comic.FinalChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comicchapter.FieldID, field.TypeString),
			},
		}
		edge.Schema = cuo.schemaConfig.Comic
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = cuo.schemaConfig.Comic
	ctx = internal.NewSchemaConfigContext(ctx, cuo.schemaConfig)
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Comic{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}