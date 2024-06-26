// Code generated by ent, DO NOT EDIT.

package comicimg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldUpdatedAt, v))
}

// B2key applies equality check predicate on the "b2key" field. It's identical to B2keyEQ.
func B2key(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldB2key, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldHeight, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldName, v))
}

// OptimizedSize applies equality check predicate on the "optimized_size" field. It's identical to OptimizedSizeEQ.
func OptimizedSize(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldOptimizedSize, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldSize, v))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldWidth, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldUpdatedAt, v))
}

// B2keyEQ applies the EQ predicate on the "b2key" field.
func B2keyEQ(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldB2key, v))
}

// B2keyNEQ applies the NEQ predicate on the "b2key" field.
func B2keyNEQ(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldB2key, v))
}

// B2keyIn applies the In predicate on the "b2key" field.
func B2keyIn(vs ...string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldB2key, vs...))
}

// B2keyNotIn applies the NotIn predicate on the "b2key" field.
func B2keyNotIn(vs ...string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldB2key, vs...))
}

// B2keyGT applies the GT predicate on the "b2key" field.
func B2keyGT(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldB2key, v))
}

// B2keyGTE applies the GTE predicate on the "b2key" field.
func B2keyGTE(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldB2key, v))
}

// B2keyLT applies the LT predicate on the "b2key" field.
func B2keyLT(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldB2key, v))
}

// B2keyLTE applies the LTE predicate on the "b2key" field.
func B2keyLTE(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldB2key, v))
}

// B2keyContains applies the Contains predicate on the "b2key" field.
func B2keyContains(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldContains(FieldB2key, v))
}

// B2keyHasPrefix applies the HasPrefix predicate on the "b2key" field.
func B2keyHasPrefix(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldHasPrefix(FieldB2key, v))
}

// B2keyHasSuffix applies the HasSuffix predicate on the "b2key" field.
func B2keyHasSuffix(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldHasSuffix(FieldB2key, v))
}

// B2keyEqualFold applies the EqualFold predicate on the "b2key" field.
func B2keyEqualFold(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEqualFold(FieldB2key, v))
}

// B2keyContainsFold applies the ContainsFold predicate on the "b2key" field.
func B2keyContainsFold(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldContainsFold(FieldB2key, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldHeight, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldContainsFold(FieldName, v))
}

// OptimizedSizeEQ applies the EQ predicate on the "optimized_size" field.
func OptimizedSizeEQ(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldOptimizedSize, v))
}

// OptimizedSizeNEQ applies the NEQ predicate on the "optimized_size" field.
func OptimizedSizeNEQ(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldOptimizedSize, v))
}

// OptimizedSizeIn applies the In predicate on the "optimized_size" field.
func OptimizedSizeIn(vs ...int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldOptimizedSize, vs...))
}

// OptimizedSizeNotIn applies the NotIn predicate on the "optimized_size" field.
func OptimizedSizeNotIn(vs ...int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldOptimizedSize, vs...))
}

// OptimizedSizeGT applies the GT predicate on the "optimized_size" field.
func OptimizedSizeGT(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldOptimizedSize, v))
}

// OptimizedSizeGTE applies the GTE predicate on the "optimized_size" field.
func OptimizedSizeGTE(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldOptimizedSize, v))
}

// OptimizedSizeLT applies the LT predicate on the "optimized_size" field.
func OptimizedSizeLT(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldOptimizedSize, v))
}

// OptimizedSizeLTE applies the LTE predicate on the "optimized_size" field.
func OptimizedSizeLTE(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldOptimizedSize, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldSize, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int) predicate.ComicImg {
	return predicate.ComicImg(sql.FieldLTE(FieldWidth, v))
}

// HasChapter applies the HasEdge predicate on the "chapter" edge.
func HasChapter() predicate.ComicImg {
	return predicate.ComicImg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChapterTable, ChapterColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ComicChapter
		step.Edge.Schema = schemaConfig.ComicImg
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChapterWith applies the HasEdge predicate on the "chapter" edge with a given conditions (other predicates).
func HasChapterWith(preds ...predicate.ComicChapter) predicate.ComicImg {
	return predicate.ComicImg(func(s *sql.Selector) {
		step := newChapterStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ComicChapter
		step.Edge.Schema = schemaConfig.ComicImg
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ComicImg) predicate.ComicImg {
	return predicate.ComicImg(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ComicImg) predicate.ComicImg {
	return predicate.ComicImg(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ComicImg) predicate.ComicImg {
	return predicate.ComicImg(sql.NotPredicates(p))
}
