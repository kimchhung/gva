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
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/manga"
	"github.com/gva/internal/ent/mangachapter"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// MangaChapterUpdate is the builder for updating MangaChapter entities.
type MangaChapterUpdate struct {
	config
	hooks     []Hook
	mutation  *MangaChapterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MangaChapterUpdate builder.
func (mcu *MangaChapterUpdate) Where(ps ...predicate.MangaChapter) *MangaChapterUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetCreatedAt sets the "created_at" field.
func (mcu *MangaChapterUpdate) SetCreatedAt(t time.Time) *MangaChapterUpdate {
	mcu.mutation.SetCreatedAt(t)
	return mcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableCreatedAt(t *time.Time) *MangaChapterUpdate {
	if t != nil {
		mcu.SetCreatedAt(*t)
	}
	return mcu
}

// SetUpdatedAt sets the "updated_at" field.
func (mcu *MangaChapterUpdate) SetUpdatedAt(t time.Time) *MangaChapterUpdate {
	mcu.mutation.SetUpdatedAt(t)
	return mcu
}

// SetMangaID sets the "manga_id" field.
func (mcu *MangaChapterUpdate) SetMangaID(px pxid.ID) *MangaChapterUpdate {
	mcu.mutation.SetMangaID(px)
	return mcu
}

// SetNillableMangaID sets the "manga_id" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableMangaID(px *pxid.ID) *MangaChapterUpdate {
	if px != nil {
		mcu.SetMangaID(*px)
	}
	return mcu
}

// SetTitle sets the "title" field.
func (mcu *MangaChapterUpdate) SetTitle(s string) *MangaChapterUpdate {
	mcu.mutation.SetTitle(s)
	return mcu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableTitle(s *string) *MangaChapterUpdate {
	if s != nil {
		mcu.SetTitle(*s)
	}
	return mcu
}

// SetImgURL sets the "img_url" field.
func (mcu *MangaChapterUpdate) SetImgURL(s string) *MangaChapterUpdate {
	mcu.mutation.SetImgURL(s)
	return mcu
}

// SetNillableImgURL sets the "img_url" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableImgURL(s *string) *MangaChapterUpdate {
	if s != nil {
		mcu.SetImgURL(*s)
	}
	return mcu
}

// SetNumber sets the "number" field.
func (mcu *MangaChapterUpdate) SetNumber(u uint) *MangaChapterUpdate {
	mcu.mutation.ResetNumber()
	mcu.mutation.SetNumber(u)
	return mcu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableNumber(u *uint) *MangaChapterUpdate {
	if u != nil {
		mcu.SetNumber(*u)
	}
	return mcu
}

// AddNumber adds u to the "number" field.
func (mcu *MangaChapterUpdate) AddNumber(u int) *MangaChapterUpdate {
	mcu.mutation.AddNumber(u)
	return mcu
}

// SetProviderName sets the "provider_name" field.
func (mcu *MangaChapterUpdate) SetProviderName(s string) *MangaChapterUpdate {
	mcu.mutation.SetProviderName(s)
	return mcu
}

// SetNillableProviderName sets the "provider_name" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableProviderName(s *string) *MangaChapterUpdate {
	if s != nil {
		mcu.SetProviderName(*s)
	}
	return mcu
}

// SetChapterUpdatedAt sets the "chapter_updated_at" field.
func (mcu *MangaChapterUpdate) SetChapterUpdatedAt(t time.Time) *MangaChapterUpdate {
	mcu.mutation.SetChapterUpdatedAt(t)
	return mcu
}

// SetNillableChapterUpdatedAt sets the "chapter_updated_at" field if the given value is not nil.
func (mcu *MangaChapterUpdate) SetNillableChapterUpdatedAt(t *time.Time) *MangaChapterUpdate {
	if t != nil {
		mcu.SetChapterUpdatedAt(*t)
	}
	return mcu
}

// SetManga sets the "manga" edge to the Manga entity.
func (mcu *MangaChapterUpdate) SetManga(m *Manga) *MangaChapterUpdate {
	return mcu.SetMangaID(m.ID)
}

// Mutation returns the MangaChapterMutation object of the builder.
func (mcu *MangaChapterUpdate) Mutation() *MangaChapterMutation {
	return mcu.mutation
}

// ClearManga clears the "manga" edge to the Manga entity.
func (mcu *MangaChapterUpdate) ClearManga() *MangaChapterUpdate {
	mcu.mutation.ClearManga()
	return mcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MangaChapterUpdate) Save(ctx context.Context) (int, error) {
	mcu.defaults()
	return withHooks(ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MangaChapterUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MangaChapterUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MangaChapterUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcu *MangaChapterUpdate) defaults() {
	if _, ok := mcu.mutation.UpdatedAt(); !ok {
		v := mangachapter.UpdateDefaultUpdatedAt()
		mcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcu *MangaChapterUpdate) check() error {
	if v, ok := mcu.mutation.MangaID(); ok {
		if err := mangachapter.MangaIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "manga_id", err: fmt.Errorf(`ent: validator failed for field "MangaChapter.manga_id": %w`, err)}
		}
	}
	if mcu.mutation.MangaCleared() && len(mcu.mutation.MangaIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MangaChapter.manga"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mcu *MangaChapterUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MangaChapterUpdate {
	mcu.modifiers = append(mcu.modifiers, modifiers...)
	return mcu
}

func (mcu *MangaChapterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(mangachapter.Table, mangachapter.Columns, sqlgraph.NewFieldSpec(mangachapter.FieldID, field.TypeString))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcu.mutation.CreatedAt(); ok {
		_spec.SetField(mangachapter.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mcu.mutation.UpdatedAt(); ok {
		_spec.SetField(mangachapter.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mcu.mutation.Title(); ok {
		_spec.SetField(mangachapter.FieldTitle, field.TypeString, value)
	}
	if value, ok := mcu.mutation.ImgURL(); ok {
		_spec.SetField(mangachapter.FieldImgURL, field.TypeString, value)
	}
	if value, ok := mcu.mutation.Number(); ok {
		_spec.SetField(mangachapter.FieldNumber, field.TypeUint, value)
	}
	if value, ok := mcu.mutation.AddedNumber(); ok {
		_spec.AddField(mangachapter.FieldNumber, field.TypeUint, value)
	}
	if value, ok := mcu.mutation.ProviderName(); ok {
		_spec.SetField(mangachapter.FieldProviderName, field.TypeString, value)
	}
	if value, ok := mcu.mutation.ChapterUpdatedAt(); ok {
		_spec.SetField(mangachapter.FieldChapterUpdatedAt, field.TypeTime, value)
	}
	if mcu.mutation.MangaCleared() {
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
		edge.Schema = mcu.schemaConfig.MangaChapter
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.MangaIDs(); len(nodes) > 0 {
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
		edge.Schema = mcu.schemaConfig.MangaChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = mcu.schemaConfig.MangaChapter
	ctx = internal.NewSchemaConfigContext(ctx, mcu.schemaConfig)
	_spec.AddModifiers(mcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mangachapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MangaChapterUpdateOne is the builder for updating a single MangaChapter entity.
type MangaChapterUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MangaChapterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (mcuo *MangaChapterUpdateOne) SetCreatedAt(t time.Time) *MangaChapterUpdateOne {
	mcuo.mutation.SetCreatedAt(t)
	return mcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableCreatedAt(t *time.Time) *MangaChapterUpdateOne {
	if t != nil {
		mcuo.SetCreatedAt(*t)
	}
	return mcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (mcuo *MangaChapterUpdateOne) SetUpdatedAt(t time.Time) *MangaChapterUpdateOne {
	mcuo.mutation.SetUpdatedAt(t)
	return mcuo
}

// SetMangaID sets the "manga_id" field.
func (mcuo *MangaChapterUpdateOne) SetMangaID(px pxid.ID) *MangaChapterUpdateOne {
	mcuo.mutation.SetMangaID(px)
	return mcuo
}

// SetNillableMangaID sets the "manga_id" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableMangaID(px *pxid.ID) *MangaChapterUpdateOne {
	if px != nil {
		mcuo.SetMangaID(*px)
	}
	return mcuo
}

// SetTitle sets the "title" field.
func (mcuo *MangaChapterUpdateOne) SetTitle(s string) *MangaChapterUpdateOne {
	mcuo.mutation.SetTitle(s)
	return mcuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableTitle(s *string) *MangaChapterUpdateOne {
	if s != nil {
		mcuo.SetTitle(*s)
	}
	return mcuo
}

// SetImgURL sets the "img_url" field.
func (mcuo *MangaChapterUpdateOne) SetImgURL(s string) *MangaChapterUpdateOne {
	mcuo.mutation.SetImgURL(s)
	return mcuo
}

// SetNillableImgURL sets the "img_url" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableImgURL(s *string) *MangaChapterUpdateOne {
	if s != nil {
		mcuo.SetImgURL(*s)
	}
	return mcuo
}

// SetNumber sets the "number" field.
func (mcuo *MangaChapterUpdateOne) SetNumber(u uint) *MangaChapterUpdateOne {
	mcuo.mutation.ResetNumber()
	mcuo.mutation.SetNumber(u)
	return mcuo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableNumber(u *uint) *MangaChapterUpdateOne {
	if u != nil {
		mcuo.SetNumber(*u)
	}
	return mcuo
}

// AddNumber adds u to the "number" field.
func (mcuo *MangaChapterUpdateOne) AddNumber(u int) *MangaChapterUpdateOne {
	mcuo.mutation.AddNumber(u)
	return mcuo
}

// SetProviderName sets the "provider_name" field.
func (mcuo *MangaChapterUpdateOne) SetProviderName(s string) *MangaChapterUpdateOne {
	mcuo.mutation.SetProviderName(s)
	return mcuo
}

// SetNillableProviderName sets the "provider_name" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableProviderName(s *string) *MangaChapterUpdateOne {
	if s != nil {
		mcuo.SetProviderName(*s)
	}
	return mcuo
}

// SetChapterUpdatedAt sets the "chapter_updated_at" field.
func (mcuo *MangaChapterUpdateOne) SetChapterUpdatedAt(t time.Time) *MangaChapterUpdateOne {
	mcuo.mutation.SetChapterUpdatedAt(t)
	return mcuo
}

// SetNillableChapterUpdatedAt sets the "chapter_updated_at" field if the given value is not nil.
func (mcuo *MangaChapterUpdateOne) SetNillableChapterUpdatedAt(t *time.Time) *MangaChapterUpdateOne {
	if t != nil {
		mcuo.SetChapterUpdatedAt(*t)
	}
	return mcuo
}

// SetManga sets the "manga" edge to the Manga entity.
func (mcuo *MangaChapterUpdateOne) SetManga(m *Manga) *MangaChapterUpdateOne {
	return mcuo.SetMangaID(m.ID)
}

// Mutation returns the MangaChapterMutation object of the builder.
func (mcuo *MangaChapterUpdateOne) Mutation() *MangaChapterMutation {
	return mcuo.mutation
}

// ClearManga clears the "manga" edge to the Manga entity.
func (mcuo *MangaChapterUpdateOne) ClearManga() *MangaChapterUpdateOne {
	mcuo.mutation.ClearManga()
	return mcuo
}

// Where appends a list predicates to the MangaChapterUpdate builder.
func (mcuo *MangaChapterUpdateOne) Where(ps ...predicate.MangaChapter) *MangaChapterUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MangaChapterUpdateOne) Select(field string, fields ...string) *MangaChapterUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MangaChapter entity.
func (mcuo *MangaChapterUpdateOne) Save(ctx context.Context) (*MangaChapter, error) {
	mcuo.defaults()
	return withHooks(ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MangaChapterUpdateOne) SaveX(ctx context.Context) *MangaChapter {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MangaChapterUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MangaChapterUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcuo *MangaChapterUpdateOne) defaults() {
	if _, ok := mcuo.mutation.UpdatedAt(); !ok {
		v := mangachapter.UpdateDefaultUpdatedAt()
		mcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcuo *MangaChapterUpdateOne) check() error {
	if v, ok := mcuo.mutation.MangaID(); ok {
		if err := mangachapter.MangaIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "manga_id", err: fmt.Errorf(`ent: validator failed for field "MangaChapter.manga_id": %w`, err)}
		}
	}
	if mcuo.mutation.MangaCleared() && len(mcuo.mutation.MangaIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MangaChapter.manga"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mcuo *MangaChapterUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MangaChapterUpdateOne {
	mcuo.modifiers = append(mcuo.modifiers, modifiers...)
	return mcuo
}

func (mcuo *MangaChapterUpdateOne) sqlSave(ctx context.Context) (_node *MangaChapter, err error) {
	if err := mcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(mangachapter.Table, mangachapter.Columns, sqlgraph.NewFieldSpec(mangachapter.FieldID, field.TypeString))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MangaChapter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mangachapter.FieldID)
		for _, f := range fields {
			if !mangachapter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mangachapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcuo.mutation.CreatedAt(); ok {
		_spec.SetField(mangachapter.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(mangachapter.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mcuo.mutation.Title(); ok {
		_spec.SetField(mangachapter.FieldTitle, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.ImgURL(); ok {
		_spec.SetField(mangachapter.FieldImgURL, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.Number(); ok {
		_spec.SetField(mangachapter.FieldNumber, field.TypeUint, value)
	}
	if value, ok := mcuo.mutation.AddedNumber(); ok {
		_spec.AddField(mangachapter.FieldNumber, field.TypeUint, value)
	}
	if value, ok := mcuo.mutation.ProviderName(); ok {
		_spec.SetField(mangachapter.FieldProviderName, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.ChapterUpdatedAt(); ok {
		_spec.SetField(mangachapter.FieldChapterUpdatedAt, field.TypeTime, value)
	}
	if mcuo.mutation.MangaCleared() {
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
		edge.Schema = mcuo.schemaConfig.MangaChapter
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.MangaIDs(); len(nodes) > 0 {
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
		edge.Schema = mcuo.schemaConfig.MangaChapter
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = mcuo.schemaConfig.MangaChapter
	ctx = internal.NewSchemaConfigContext(ctx, mcuo.schemaConfig)
	_spec.AddModifiers(mcuo.modifiers...)
	_node = &MangaChapter{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mangachapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}