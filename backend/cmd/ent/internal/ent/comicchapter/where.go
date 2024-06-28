// Code generated by ent, DO NOT EDIT.

package comicchapter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldUpdatedAt, v))
}

// Chapter applies equality check predicate on the "chapter" field. It's identical to ChapterEQ.
func Chapter(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldChapter, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldTitle, v))
}

// Volumn applies equality check predicate on the "volumn" field. It's identical to VolumnEQ.
func Volumn(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldVolumn, v))
}

// Lang applies equality check predicate on the "lang" field. It's identical to LangEQ.
func Lang(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldLang, v))
}

// UpCount applies equality check predicate on the "up_count" field. It's identical to UpCountEQ.
func UpCount(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldUpCount, v))
}

// DownCount applies equality check predicate on the "down_count" field. It's identical to DownCountEQ.
func DownCount(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldDownCount, v))
}

// IsLastChapter applies equality check predicate on the "is_last_chapter" field. It's identical to IsLastChapterEQ.
func IsLastChapter(v bool) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldIsLastChapter, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldUpdatedAt, v))
}

// ChapterEQ applies the EQ predicate on the "chapter" field.
func ChapterEQ(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldChapter, v))
}

// ChapterNEQ applies the NEQ predicate on the "chapter" field.
func ChapterNEQ(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldChapter, v))
}

// ChapterIn applies the In predicate on the "chapter" field.
func ChapterIn(vs ...uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldChapter, vs...))
}

// ChapterNotIn applies the NotIn predicate on the "chapter" field.
func ChapterNotIn(vs ...uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldChapter, vs...))
}

// ChapterGT applies the GT predicate on the "chapter" field.
func ChapterGT(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldChapter, v))
}

// ChapterGTE applies the GTE predicate on the "chapter" field.
func ChapterGTE(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldChapter, v))
}

// ChapterLT applies the LT predicate on the "chapter" field.
func ChapterLT(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldChapter, v))
}

// ChapterLTE applies the LTE predicate on the "chapter" field.
func ChapterLTE(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldChapter, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldContainsFold(FieldTitle, v))
}

// VolumnEQ applies the EQ predicate on the "volumn" field.
func VolumnEQ(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldVolumn, v))
}

// VolumnNEQ applies the NEQ predicate on the "volumn" field.
func VolumnNEQ(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldVolumn, v))
}

// VolumnIn applies the In predicate on the "volumn" field.
func VolumnIn(vs ...string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldVolumn, vs...))
}

// VolumnNotIn applies the NotIn predicate on the "volumn" field.
func VolumnNotIn(vs ...string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldVolumn, vs...))
}

// VolumnGT applies the GT predicate on the "volumn" field.
func VolumnGT(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldVolumn, v))
}

// VolumnGTE applies the GTE predicate on the "volumn" field.
func VolumnGTE(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldVolumn, v))
}

// VolumnLT applies the LT predicate on the "volumn" field.
func VolumnLT(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldVolumn, v))
}

// VolumnLTE applies the LTE predicate on the "volumn" field.
func VolumnLTE(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldVolumn, v))
}

// VolumnContains applies the Contains predicate on the "volumn" field.
func VolumnContains(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldContains(FieldVolumn, v))
}

// VolumnHasPrefix applies the HasPrefix predicate on the "volumn" field.
func VolumnHasPrefix(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldHasPrefix(FieldVolumn, v))
}

// VolumnHasSuffix applies the HasSuffix predicate on the "volumn" field.
func VolumnHasSuffix(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldHasSuffix(FieldVolumn, v))
}

// VolumnIsNil applies the IsNil predicate on the "volumn" field.
func VolumnIsNil() predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIsNull(FieldVolumn))
}

// VolumnNotNil applies the NotNil predicate on the "volumn" field.
func VolumnNotNil() predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotNull(FieldVolumn))
}

// VolumnEqualFold applies the EqualFold predicate on the "volumn" field.
func VolumnEqualFold(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEqualFold(FieldVolumn, v))
}

// VolumnContainsFold applies the ContainsFold predicate on the "volumn" field.
func VolumnContainsFold(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldContainsFold(FieldVolumn, v))
}

// LangEQ applies the EQ predicate on the "lang" field.
func LangEQ(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldLang, v))
}

// LangNEQ applies the NEQ predicate on the "lang" field.
func LangNEQ(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldLang, v))
}

// LangIn applies the In predicate on the "lang" field.
func LangIn(vs ...string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldLang, vs...))
}

// LangNotIn applies the NotIn predicate on the "lang" field.
func LangNotIn(vs ...string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldLang, vs...))
}

// LangGT applies the GT predicate on the "lang" field.
func LangGT(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldLang, v))
}

// LangGTE applies the GTE predicate on the "lang" field.
func LangGTE(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldLang, v))
}

// LangLT applies the LT predicate on the "lang" field.
func LangLT(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldLang, v))
}

// LangLTE applies the LTE predicate on the "lang" field.
func LangLTE(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldLang, v))
}

// LangContains applies the Contains predicate on the "lang" field.
func LangContains(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldContains(FieldLang, v))
}

// LangHasPrefix applies the HasPrefix predicate on the "lang" field.
func LangHasPrefix(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldHasPrefix(FieldLang, v))
}

// LangHasSuffix applies the HasSuffix predicate on the "lang" field.
func LangHasSuffix(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldHasSuffix(FieldLang, v))
}

// LangEqualFold applies the EqualFold predicate on the "lang" field.
func LangEqualFold(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEqualFold(FieldLang, v))
}

// LangContainsFold applies the ContainsFold predicate on the "lang" field.
func LangContainsFold(v string) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldContainsFold(FieldLang, v))
}

// UpCountEQ applies the EQ predicate on the "up_count" field.
func UpCountEQ(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldUpCount, v))
}

// UpCountNEQ applies the NEQ predicate on the "up_count" field.
func UpCountNEQ(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldUpCount, v))
}

// UpCountIn applies the In predicate on the "up_count" field.
func UpCountIn(vs ...uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldUpCount, vs...))
}

// UpCountNotIn applies the NotIn predicate on the "up_count" field.
func UpCountNotIn(vs ...uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldUpCount, vs...))
}

// UpCountGT applies the GT predicate on the "up_count" field.
func UpCountGT(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldUpCount, v))
}

// UpCountGTE applies the GTE predicate on the "up_count" field.
func UpCountGTE(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldUpCount, v))
}

// UpCountLT applies the LT predicate on the "up_count" field.
func UpCountLT(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldUpCount, v))
}

// UpCountLTE applies the LTE predicate on the "up_count" field.
func UpCountLTE(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldUpCount, v))
}

// DownCountEQ applies the EQ predicate on the "down_count" field.
func DownCountEQ(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldDownCount, v))
}

// DownCountNEQ applies the NEQ predicate on the "down_count" field.
func DownCountNEQ(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldDownCount, v))
}

// DownCountIn applies the In predicate on the "down_count" field.
func DownCountIn(vs ...uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldIn(FieldDownCount, vs...))
}

// DownCountNotIn applies the NotIn predicate on the "down_count" field.
func DownCountNotIn(vs ...uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNotIn(FieldDownCount, vs...))
}

// DownCountGT applies the GT predicate on the "down_count" field.
func DownCountGT(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGT(FieldDownCount, v))
}

// DownCountGTE applies the GTE predicate on the "down_count" field.
func DownCountGTE(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldGTE(FieldDownCount, v))
}

// DownCountLT applies the LT predicate on the "down_count" field.
func DownCountLT(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLT(FieldDownCount, v))
}

// DownCountLTE applies the LTE predicate on the "down_count" field.
func DownCountLTE(v uint) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldLTE(FieldDownCount, v))
}

// IsLastChapterEQ applies the EQ predicate on the "is_last_chapter" field.
func IsLastChapterEQ(v bool) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldEQ(FieldIsLastChapter, v))
}

// IsLastChapterNEQ applies the NEQ predicate on the "is_last_chapter" field.
func IsLastChapterNEQ(v bool) predicate.ComicChapter {
	return predicate.ComicChapter(sql.FieldNEQ(FieldIsLastChapter, v))
}

// HasImgs applies the HasEdge predicate on the "imgs" edge.
func HasImgs() predicate.ComicChapter {
	return predicate.ComicChapter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImgsTable, ImgsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ComicImg
		step.Edge.Schema = schemaConfig.ComicImg
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImgsWith applies the HasEdge predicate on the "imgs" edge with a given conditions (other predicates).
func HasImgsWith(preds ...predicate.ComicImg) predicate.ComicChapter {
	return predicate.ComicChapter(func(s *sql.Selector) {
		step := newImgsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ComicImg
		step.Edge.Schema = schemaConfig.ComicImg
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComic applies the HasEdge predicate on the "comic" edge.
func HasComic() predicate.ComicChapter {
	return predicate.ComicChapter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComicTable, ComicColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Comic
		step.Edge.Schema = schemaConfig.ComicChapter
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasComicWith applies the HasEdge predicate on the "comic" edge with a given conditions (other predicates).
func HasComicWith(preds ...predicate.Comic) predicate.ComicChapter {
	return predicate.ComicChapter(func(s *sql.Selector) {
		step := newComicStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Comic
		step.Edge.Schema = schemaConfig.ComicChapter
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ComicChapter) predicate.ComicChapter {
	return predicate.ComicChapter(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ComicChapter) predicate.ComicChapter {
	return predicate.ComicChapter(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ComicChapter) predicate.ComicChapter {
	return predicate.ComicChapter(sql.NotPredicates(p))
}
