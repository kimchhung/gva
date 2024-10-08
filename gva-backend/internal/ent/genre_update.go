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
	"github.com/gva/internal/ent/genre"
	"github.com/gva/internal/ent/manga"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// GenreUpdate is the builder for updating Genre entities.
type GenreUpdate struct {
	config
	hooks     []Hook
	mutation  *GenreMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GenreUpdate builder.
func (gu *GenreUpdate) Where(ps ...predicate.Genre) *GenreUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GenreUpdate) SetCreatedAt(t time.Time) *GenreUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GenreUpdate) SetNillableCreatedAt(t *time.Time) *GenreUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GenreUpdate) SetUpdatedAt(t time.Time) *GenreUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetIsEnable sets the "is_enable" field.
func (gu *GenreUpdate) SetIsEnable(b bool) *GenreUpdate {
	gu.mutation.SetIsEnable(b)
	return gu
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (gu *GenreUpdate) SetNillableIsEnable(b *bool) *GenreUpdate {
	if b != nil {
		gu.SetIsEnable(*b)
	}
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GenreUpdate) SetDeletedAt(i int) *GenreUpdate {
	gu.mutation.ResetDeletedAt()
	gu.mutation.SetDeletedAt(i)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GenreUpdate) SetNillableDeletedAt(i *int) *GenreUpdate {
	if i != nil {
		gu.SetDeletedAt(*i)
	}
	return gu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (gu *GenreUpdate) AddDeletedAt(i int) *GenreUpdate {
	gu.mutation.AddDeletedAt(i)
	return gu
}

// SetName sets the "name" field.
func (gu *GenreUpdate) SetName(s string) *GenreUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gu *GenreUpdate) SetNillableName(s *string) *GenreUpdate {
	if s != nil {
		gu.SetName(*s)
	}
	return gu
}

// SetNameID sets the "name_id" field.
func (gu *GenreUpdate) SetNameID(s string) *GenreUpdate {
	gu.mutation.SetNameID(s)
	return gu
}

// SetNillableNameID sets the "name_id" field if the given value is not nil.
func (gu *GenreUpdate) SetNillableNameID(s *string) *GenreUpdate {
	if s != nil {
		gu.SetNameID(*s)
	}
	return gu
}

// AddMangaIDs adds the "mangas" edge to the Manga entity by IDs.
func (gu *GenreUpdate) AddMangaIDs(ids ...pxid.ID) *GenreUpdate {
	gu.mutation.AddMangaIDs(ids...)
	return gu
}

// AddMangas adds the "mangas" edges to the Manga entity.
func (gu *GenreUpdate) AddMangas(m ...*Manga) *GenreUpdate {
	ids := make([]pxid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMangaIDs(ids...)
}

// Mutation returns the GenreMutation object of the builder.
func (gu *GenreUpdate) Mutation() *GenreMutation {
	return gu.mutation
}

// ClearMangas clears all "mangas" edges to the Manga entity.
func (gu *GenreUpdate) ClearMangas() *GenreUpdate {
	gu.mutation.ClearMangas()
	return gu
}

// RemoveMangaIDs removes the "mangas" edge to Manga entities by IDs.
func (gu *GenreUpdate) RemoveMangaIDs(ids ...pxid.ID) *GenreUpdate {
	gu.mutation.RemoveMangaIDs(ids...)
	return gu
}

// RemoveMangas removes "mangas" edges to Manga entities.
func (gu *GenreUpdate) RemoveMangas(m ...*Manga) *GenreUpdate {
	ids := make([]pxid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMangaIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GenreUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GenreUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GenreUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GenreUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GenreUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := genre.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gu *GenreUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GenreUpdate {
	gu.modifiers = append(gu.modifiers, modifiers...)
	return gu
}

func (gu *GenreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(genre.Table, genre.Columns, sqlgraph.NewFieldSpec(genre.FieldID, field.TypeString))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.SetField(genre.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(genre.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.IsEnable(); ok {
		_spec.SetField(genre.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.SetField(genre.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(genre.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(genre.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.NameID(); ok {
		_spec.SetField(genre.FieldNameID, field.TypeString, value)
	}
	if gu.mutation.MangasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GenreMangas
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedMangasIDs(); len(nodes) > 0 && !gu.mutation.MangasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GenreMangas
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MangasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GenreMangas
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = gu.schemaConfig.Genre
	ctx = internal.NewSchemaConfigContext(ctx, gu.schemaConfig)
	_spec.AddModifiers(gu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{genre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GenreUpdateOne is the builder for updating a single Genre entity.
type GenreUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GenreMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (guo *GenreUpdateOne) SetCreatedAt(t time.Time) *GenreUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GenreUpdateOne) SetNillableCreatedAt(t *time.Time) *GenreUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GenreUpdateOne) SetUpdatedAt(t time.Time) *GenreUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetIsEnable sets the "is_enable" field.
func (guo *GenreUpdateOne) SetIsEnable(b bool) *GenreUpdateOne {
	guo.mutation.SetIsEnable(b)
	return guo
}

// SetNillableIsEnable sets the "is_enable" field if the given value is not nil.
func (guo *GenreUpdateOne) SetNillableIsEnable(b *bool) *GenreUpdateOne {
	if b != nil {
		guo.SetIsEnable(*b)
	}
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GenreUpdateOne) SetDeletedAt(i int) *GenreUpdateOne {
	guo.mutation.ResetDeletedAt()
	guo.mutation.SetDeletedAt(i)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GenreUpdateOne) SetNillableDeletedAt(i *int) *GenreUpdateOne {
	if i != nil {
		guo.SetDeletedAt(*i)
	}
	return guo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (guo *GenreUpdateOne) AddDeletedAt(i int) *GenreUpdateOne {
	guo.mutation.AddDeletedAt(i)
	return guo
}

// SetName sets the "name" field.
func (guo *GenreUpdateOne) SetName(s string) *GenreUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guo *GenreUpdateOne) SetNillableName(s *string) *GenreUpdateOne {
	if s != nil {
		guo.SetName(*s)
	}
	return guo
}

// SetNameID sets the "name_id" field.
func (guo *GenreUpdateOne) SetNameID(s string) *GenreUpdateOne {
	guo.mutation.SetNameID(s)
	return guo
}

// SetNillableNameID sets the "name_id" field if the given value is not nil.
func (guo *GenreUpdateOne) SetNillableNameID(s *string) *GenreUpdateOne {
	if s != nil {
		guo.SetNameID(*s)
	}
	return guo
}

// AddMangaIDs adds the "mangas" edge to the Manga entity by IDs.
func (guo *GenreUpdateOne) AddMangaIDs(ids ...pxid.ID) *GenreUpdateOne {
	guo.mutation.AddMangaIDs(ids...)
	return guo
}

// AddMangas adds the "mangas" edges to the Manga entity.
func (guo *GenreUpdateOne) AddMangas(m ...*Manga) *GenreUpdateOne {
	ids := make([]pxid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMangaIDs(ids...)
}

// Mutation returns the GenreMutation object of the builder.
func (guo *GenreUpdateOne) Mutation() *GenreMutation {
	return guo.mutation
}

// ClearMangas clears all "mangas" edges to the Manga entity.
func (guo *GenreUpdateOne) ClearMangas() *GenreUpdateOne {
	guo.mutation.ClearMangas()
	return guo
}

// RemoveMangaIDs removes the "mangas" edge to Manga entities by IDs.
func (guo *GenreUpdateOne) RemoveMangaIDs(ids ...pxid.ID) *GenreUpdateOne {
	guo.mutation.RemoveMangaIDs(ids...)
	return guo
}

// RemoveMangas removes "mangas" edges to Manga entities.
func (guo *GenreUpdateOne) RemoveMangas(m ...*Manga) *GenreUpdateOne {
	ids := make([]pxid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMangaIDs(ids...)
}

// Where appends a list predicates to the GenreUpdate builder.
func (guo *GenreUpdateOne) Where(ps ...predicate.Genre) *GenreUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GenreUpdateOne) Select(field string, fields ...string) *GenreUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Genre entity.
func (guo *GenreUpdateOne) Save(ctx context.Context) (*Genre, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GenreUpdateOne) SaveX(ctx context.Context) *Genre {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GenreUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GenreUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GenreUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := genre.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (guo *GenreUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GenreUpdateOne {
	guo.modifiers = append(guo.modifiers, modifiers...)
	return guo
}

func (guo *GenreUpdateOne) sqlSave(ctx context.Context) (_node *Genre, err error) {
	_spec := sqlgraph.NewUpdateSpec(genre.Table, genre.Columns, sqlgraph.NewFieldSpec(genre.FieldID, field.TypeString))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Genre.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, genre.FieldID)
		for _, f := range fields {
			if !genre.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != genre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.SetField(genre.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(genre.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.IsEnable(); ok {
		_spec.SetField(genre.FieldIsEnable, field.TypeBool, value)
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.SetField(genre.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(genre.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(genre.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.NameID(); ok {
		_spec.SetField(genre.FieldNameID, field.TypeString, value)
	}
	if guo.mutation.MangasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GenreMangas
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedMangasIDs(); len(nodes) > 0 && !guo.mutation.MangasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GenreMangas
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MangasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(manga.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GenreMangas
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = guo.schemaConfig.Genre
	ctx = internal.NewSchemaConfigContext(ctx, guo.schemaConfig)
	_spec.AddModifiers(guo.modifiers...)
	_node = &Genre{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{genre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
