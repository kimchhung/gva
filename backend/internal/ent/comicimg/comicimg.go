// Code generated by ent, DO NOT EDIT.

package comicimg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comicimg type in the database.
	Label = "comic_img"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldB2key holds the string denoting the b2key field in the database.
	FieldB2key = "b2key"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOptimizedSize holds the string denoting the optimized_size field in the database.
	FieldOptimizedSize = "optimized_size"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// EdgeChapter holds the string denoting the chapter edge name in mutations.
	EdgeChapter = "chapter"
	// Table holds the table name of the comicimg in the database.
	Table = "comic_imgs"
	// ChapterTable is the table that holds the chapter relation/edge.
	ChapterTable = "comic_imgs"
	// ChapterInverseTable is the table name for the ComicChapter entity.
	// It exists in this package in order to avoid circular dependency with the "comicchapter" package.
	ChapterInverseTable = "comic_chapters"
	// ChapterColumn is the table column denoting the chapter relation/edge.
	ChapterColumn = "comic_chapter_imgs"
)

// Columns holds all SQL columns for comicimg fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldB2key,
	FieldHeight,
	FieldName,
	FieldOptimizedSize,
	FieldSize,
	FieldWidth,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comic_imgs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"comic_chapter_imgs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the ComicImg queries.
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

// ByB2key orders the results by the b2key field.
func ByB2key(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldB2key, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOptimizedSize orders the results by the optimized_size field.
func ByOptimizedSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOptimizedSize, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByWidth orders the results by the width field.
func ByWidth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWidth, opts...).ToFunc()
}

// ByChapterField orders the results by chapter field.
func ByChapterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChapterStep(), sql.OrderByField(field, opts...))
	}
}
func newChapterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChapterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ChapterTable, ChapterColumn),
	)
}