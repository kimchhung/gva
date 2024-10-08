// Code generated by ent, DO NOT EDIT.

package mangachapter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/pxid"
)

const (
	// Label holds the string label denoting the mangachapter type in the database.
	Label = "manga_chapter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMangaID holds the string denoting the manga_id field in the database.
	FieldMangaID = "manga_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldImgURL holds the string denoting the img_url field in the database.
	FieldImgURL = "img_url"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldProviderName holds the string denoting the provider_name field in the database.
	FieldProviderName = "provider_name"
	// FieldChapterUpdatedAt holds the string denoting the chapter_updated_at field in the database.
	FieldChapterUpdatedAt = "chapter_updated_at"
	// EdgeManga holds the string denoting the manga edge name in mutations.
	EdgeManga = "manga"
	// Table holds the table name of the mangachapter in the database.
	Table = "manga_chapters"
	// MangaTable is the table that holds the manga relation/edge.
	MangaTable = "manga_chapters"
	// MangaInverseTable is the table name for the Manga entity.
	// It exists in this package in order to avoid circular dependency with the "manga" package.
	MangaInverseTable = "mangas"
	// MangaColumn is the table column denoting the manga relation/edge.
	MangaColumn = "manga_id"
)

// Columns holds all SQL columns for mangachapter fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMangaID,
	FieldTitle,
	FieldImgURL,
	FieldNumber,
	FieldProviderName,
	FieldChapterUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// MangaIDValidator is a validator for the "manga_id" field. It is called by the builders before save.
	MangaIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pxid.ID
)

// OrderOption defines the ordering options for the MangaChapter queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByMangaID orders the results by the manga_id field.
func ByMangaID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMangaID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByImgURL orders the results by the img_url field.
func ByImgURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImgURL, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByProviderName orders the results by the provider_name field.
func ByProviderName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderName, opts...).ToFunc()
}

// ByChapterUpdatedAt orders the results by the chapter_updated_at field.
func ByChapterUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChapterUpdatedAt, opts...).ToFunc()
}

// ByMangaField orders the results by manga field.
func ByMangaField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMangaStep(), sql.OrderByField(field, opts...))
	}
}
func newMangaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MangaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MangaTable, MangaColumn),
	)
}
