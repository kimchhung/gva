// Code generated by ent, DO NOT EDIT.

package manga

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pxid.ID) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsEnable applies equality check predicate on the "is_enable" field. It's identical to IsEnableEQ.
func IsEnable(v bool) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldIsEnable, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldDeletedAt, v))
}

// NameID applies equality check predicate on the "name_id" field. It's identical to NameIDEQ.
func NameID(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldNameID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldName, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldDesc, v))
}

// Prodiver applies equality check predicate on the "prodiver" field. It's identical to ProdiverEQ.
func Prodiver(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldProdiver, v))
}

// ThumbnailURL applies equality check predicate on the "thumbnail_url" field. It's identical to ThumbnailURLEQ.
func ThumbnailURL(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldThumbnailURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsEnableEQ applies the EQ predicate on the "is_enable" field.
func IsEnableEQ(v bool) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldIsEnable, v))
}

// IsEnableNEQ applies the NEQ predicate on the "is_enable" field.
func IsEnableNEQ(v bool) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldIsEnable, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldDeletedAt, v))
}

// NameIDEQ applies the EQ predicate on the "name_id" field.
func NameIDEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldNameID, v))
}

// NameIDNEQ applies the NEQ predicate on the "name_id" field.
func NameIDNEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldNameID, v))
}

// NameIDIn applies the In predicate on the "name_id" field.
func NameIDIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldNameID, vs...))
}

// NameIDNotIn applies the NotIn predicate on the "name_id" field.
func NameIDNotIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldNameID, vs...))
}

// NameIDGT applies the GT predicate on the "name_id" field.
func NameIDGT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldNameID, v))
}

// NameIDGTE applies the GTE predicate on the "name_id" field.
func NameIDGTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldNameID, v))
}

// NameIDLT applies the LT predicate on the "name_id" field.
func NameIDLT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldNameID, v))
}

// NameIDLTE applies the LTE predicate on the "name_id" field.
func NameIDLTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldNameID, v))
}

// NameIDContains applies the Contains predicate on the "name_id" field.
func NameIDContains(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContains(FieldNameID, v))
}

// NameIDHasPrefix applies the HasPrefix predicate on the "name_id" field.
func NameIDHasPrefix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasPrefix(FieldNameID, v))
}

// NameIDHasSuffix applies the HasSuffix predicate on the "name_id" field.
func NameIDHasSuffix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasSuffix(FieldNameID, v))
}

// NameIDEqualFold applies the EqualFold predicate on the "name_id" field.
func NameIDEqualFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEqualFold(FieldNameID, v))
}

// NameIDContainsFold applies the ContainsFold predicate on the "name_id" field.
func NameIDContainsFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContainsFold(FieldNameID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContainsFold(FieldName, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContainsFold(FieldDesc, v))
}

// ProdiverEQ applies the EQ predicate on the "prodiver" field.
func ProdiverEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldProdiver, v))
}

// ProdiverNEQ applies the NEQ predicate on the "prodiver" field.
func ProdiverNEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldProdiver, v))
}

// ProdiverIn applies the In predicate on the "prodiver" field.
func ProdiverIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldProdiver, vs...))
}

// ProdiverNotIn applies the NotIn predicate on the "prodiver" field.
func ProdiverNotIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldProdiver, vs...))
}

// ProdiverGT applies the GT predicate on the "prodiver" field.
func ProdiverGT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldProdiver, v))
}

// ProdiverGTE applies the GTE predicate on the "prodiver" field.
func ProdiverGTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldProdiver, v))
}

// ProdiverLT applies the LT predicate on the "prodiver" field.
func ProdiverLT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldProdiver, v))
}

// ProdiverLTE applies the LTE predicate on the "prodiver" field.
func ProdiverLTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldProdiver, v))
}

// ProdiverContains applies the Contains predicate on the "prodiver" field.
func ProdiverContains(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContains(FieldProdiver, v))
}

// ProdiverHasPrefix applies the HasPrefix predicate on the "prodiver" field.
func ProdiverHasPrefix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasPrefix(FieldProdiver, v))
}

// ProdiverHasSuffix applies the HasSuffix predicate on the "prodiver" field.
func ProdiverHasSuffix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasSuffix(FieldProdiver, v))
}

// ProdiverEqualFold applies the EqualFold predicate on the "prodiver" field.
func ProdiverEqualFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEqualFold(FieldProdiver, v))
}

// ProdiverContainsFold applies the ContainsFold predicate on the "prodiver" field.
func ProdiverContainsFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContainsFold(FieldProdiver, v))
}

// ThumbnailURLEQ applies the EQ predicate on the "thumbnail_url" field.
func ThumbnailURLEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEQ(FieldThumbnailURL, v))
}

// ThumbnailURLNEQ applies the NEQ predicate on the "thumbnail_url" field.
func ThumbnailURLNEQ(v string) predicate.Manga {
	return predicate.Manga(sql.FieldNEQ(FieldThumbnailURL, v))
}

// ThumbnailURLIn applies the In predicate on the "thumbnail_url" field.
func ThumbnailURLIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLNotIn applies the NotIn predicate on the "thumbnail_url" field.
func ThumbnailURLNotIn(vs ...string) predicate.Manga {
	return predicate.Manga(sql.FieldNotIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLGT applies the GT predicate on the "thumbnail_url" field.
func ThumbnailURLGT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGT(FieldThumbnailURL, v))
}

// ThumbnailURLGTE applies the GTE predicate on the "thumbnail_url" field.
func ThumbnailURLGTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldGTE(FieldThumbnailURL, v))
}

// ThumbnailURLLT applies the LT predicate on the "thumbnail_url" field.
func ThumbnailURLLT(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLT(FieldThumbnailURL, v))
}

// ThumbnailURLLTE applies the LTE predicate on the "thumbnail_url" field.
func ThumbnailURLLTE(v string) predicate.Manga {
	return predicate.Manga(sql.FieldLTE(FieldThumbnailURL, v))
}

// ThumbnailURLContains applies the Contains predicate on the "thumbnail_url" field.
func ThumbnailURLContains(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContains(FieldThumbnailURL, v))
}

// ThumbnailURLHasPrefix applies the HasPrefix predicate on the "thumbnail_url" field.
func ThumbnailURLHasPrefix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasPrefix(FieldThumbnailURL, v))
}

// ThumbnailURLHasSuffix applies the HasSuffix predicate on the "thumbnail_url" field.
func ThumbnailURLHasSuffix(v string) predicate.Manga {
	return predicate.Manga(sql.FieldHasSuffix(FieldThumbnailURL, v))
}

// ThumbnailURLEqualFold applies the EqualFold predicate on the "thumbnail_url" field.
func ThumbnailURLEqualFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldEqualFold(FieldThumbnailURL, v))
}

// ThumbnailURLContainsFold applies the ContainsFold predicate on the "thumbnail_url" field.
func ThumbnailURLContainsFold(v string) predicate.Manga {
	return predicate.Manga(sql.FieldContainsFold(FieldThumbnailURL, v))
}

// HasChapters applies the HasEdge predicate on the "chapters" edge.
func HasChapters() predicate.Manga {
	return predicate.Manga(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChaptersTable, ChaptersColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.MangaChapter
		step.Edge.Schema = schemaConfig.MangaChapter
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChaptersWith applies the HasEdge predicate on the "chapters" edge with a given conditions (other predicates).
func HasChaptersWith(preds ...predicate.MangaChapter) predicate.Manga {
	return predicate.Manga(func(s *sql.Selector) {
		step := newChaptersStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.MangaChapter
		step.Edge.Schema = schemaConfig.MangaChapter
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGenres applies the HasEdge predicate on the "genres" edge.
func HasGenres() predicate.Manga {
	return predicate.Manga(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GenresTable, GenresPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Genre
		step.Edge.Schema = schemaConfig.GenreMangas
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenresWith applies the HasEdge predicate on the "genres" edge with a given conditions (other predicates).
func HasGenresWith(preds ...predicate.Genre) predicate.Manga {
	return predicate.Manga(func(s *sql.Selector) {
		step := newGenresStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Genre
		step.Edge.Schema = schemaConfig.GenreMangas
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Manga) predicate.Manga {
	return predicate.Manga(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Manga) predicate.Manga {
	return predicate.Manga(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Manga) predicate.Manga {
	return predicate.Manga(sql.NotPredicates(p))
}
