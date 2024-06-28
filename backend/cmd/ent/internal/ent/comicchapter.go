// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/comic"
	"github.com/gva/internal/ent/comicchapter"
)

// ComicChapter is the model entity for the ComicChapter schema.
type ComicChapter struct {
	config `json:"-" rql:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id" rql:"filter,sort"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Chapter holds the value of the "chapter" field.
	Chapter uint `json:"chapter,omitempty"`
	// Title holds the value of the "title" field.
	Title *string `json:"title,omitempty"`
	// Volumn holds the value of the "volumn" field.
	Volumn *string `json:"volumn,omitempty"`
	// Lang holds the value of the "lang" field.
	Lang string `json:"lang,omitempty"`
	// UpCount holds the value of the "up_count" field.
	UpCount uint `json:"upCount"`
	// DownCount holds the value of the "down_count" field.
	DownCount uint `json:"downCount"`
	// IsLastChapter holds the value of the "is_last_chapter" field.
	IsLastChapter bool `json:"isLastChapter"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ComicChapterQuery when eager-loading is set.
	Edges          ComicChapterEdges `json:"edges" rql:"-"`
	comic_chapters *pulid.ID
	selectValues   sql.SelectValues
}

// ComicChapterEdges holds the relations/edges for other nodes in the graph.
type ComicChapterEdges struct {
	// Imgs holds the value of the imgs edge.
	Imgs []*ComicImg `json:"imgs,omitempty"`
	// Comic holds the value of the comic edge.
	Comic *Comic `json:"comic,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedImgs map[string][]*ComicImg
}

// ImgsOrErr returns the Imgs value or an error if the edge
// was not loaded in eager-loading.
func (e ComicChapterEdges) ImgsOrErr() ([]*ComicImg, error) {
	if e.loadedTypes[0] {
		return e.Imgs, nil
	}
	return nil, &NotLoadedError{edge: "imgs"}
}

// ComicOrErr returns the Comic value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComicChapterEdges) ComicOrErr() (*Comic, error) {
	if e.Comic != nil {
		return e.Comic, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: comic.Label}
	}
	return nil, &NotLoadedError{edge: "comic"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ComicChapter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comicchapter.FieldID:
			values[i] = new(pulid.ID)
		case comicchapter.FieldIsLastChapter:
			values[i] = new(sql.NullBool)
		case comicchapter.FieldChapter, comicchapter.FieldUpCount, comicchapter.FieldDownCount:
			values[i] = new(sql.NullInt64)
		case comicchapter.FieldTitle, comicchapter.FieldVolumn, comicchapter.FieldLang:
			values[i] = new(sql.NullString)
		case comicchapter.FieldCreatedAt, comicchapter.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case comicchapter.ForeignKeys[0]: // comic_chapters
			values[i] = &sql.NullScanner{S: new(pulid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ComicChapter fields.
func (cc *ComicChapter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comicchapter.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cc.ID = *value
			}
		case comicchapter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cc.CreatedAt = value.Time
			}
		case comicchapter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cc.UpdatedAt = value.Time
			}
		case comicchapter.FieldChapter:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chapter", values[i])
			} else if value.Valid {
				cc.Chapter = uint(value.Int64)
			}
		case comicchapter.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				cc.Title = new(string)
				*cc.Title = value.String
			}
		case comicchapter.FieldVolumn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field volumn", values[i])
			} else if value.Valid {
				cc.Volumn = new(string)
				*cc.Volumn = value.String
			}
		case comicchapter.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				cc.Lang = value.String
			}
		case comicchapter.FieldUpCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field up_count", values[i])
			} else if value.Valid {
				cc.UpCount = uint(value.Int64)
			}
		case comicchapter.FieldDownCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field down_count", values[i])
			} else if value.Valid {
				cc.DownCount = uint(value.Int64)
			}
		case comicchapter.FieldIsLastChapter:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_last_chapter", values[i])
			} else if value.Valid {
				cc.IsLastChapter = value.Bool
			}
		case comicchapter.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field comic_chapters", values[i])
			} else if value.Valid {
				cc.comic_chapters = new(pulid.ID)
				*cc.comic_chapters = *value.S.(*pulid.ID)
			}
		default:
			cc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ComicChapter.
// This includes values selected through modifiers, order, etc.
func (cc *ComicChapter) Value(name string) (ent.Value, error) {
	return cc.selectValues.Get(name)
}

// QueryImgs queries the "imgs" edge of the ComicChapter entity.
func (cc *ComicChapter) QueryImgs() *ComicImgQuery {
	return NewComicChapterClient(cc.config).QueryImgs(cc)
}

// QueryComic queries the "comic" edge of the ComicChapter entity.
func (cc *ComicChapter) QueryComic() *ComicQuery {
	return NewComicChapterClient(cc.config).QueryComic(cc)
}

// Update returns a builder for updating this ComicChapter.
// Note that you need to call ComicChapter.Unwrap() before calling this method if this ComicChapter
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *ComicChapter) Update() *ComicChapterUpdateOne {
	return NewComicChapterClient(cc.config).UpdateOne(cc)
}

// Unwrap unwraps the ComicChapter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *ComicChapter) Unwrap() *ComicChapter {
	_tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ComicChapter is not a transactional entity")
	}
	cc.config.driver = _tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *ComicChapter) String() string {
	var builder strings.Builder
	builder.WriteString("ComicChapter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("chapter=")
	builder.WriteString(fmt.Sprintf("%v", cc.Chapter))
	builder.WriteString(", ")
	if v := cc.Title; v != nil {
		builder.WriteString("title=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := cc.Volumn; v != nil {
		builder.WriteString("volumn=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(cc.Lang)
	builder.WriteString(", ")
	builder.WriteString("up_count=")
	builder.WriteString(fmt.Sprintf("%v", cc.UpCount))
	builder.WriteString(", ")
	builder.WriteString("down_count=")
	builder.WriteString(fmt.Sprintf("%v", cc.DownCount))
	builder.WriteString(", ")
	builder.WriteString("is_last_chapter=")
	builder.WriteString(fmt.Sprintf("%v", cc.IsLastChapter))
	builder.WriteByte(')')
	return builder.String()
}

// NamedImgs returns the Imgs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (cc *ComicChapter) NamedImgs(name string) ([]*ComicImg, error) {
	if cc.Edges.namedImgs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := cc.Edges.namedImgs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (cc *ComicChapter) appendNamedImgs(name string, edges ...*ComicImg) {
	if cc.Edges.namedImgs == nil {
		cc.Edges.namedImgs = make(map[string][]*ComicImg)
	}
	if len(edges) == 0 {
		cc.Edges.namedImgs[name] = []*ComicImg{}
	} else {
		cc.Edges.namedImgs[name] = append(cc.Edges.namedImgs[name], edges...)
	}
}

// ComicChapters is a parsable slice of ComicChapter.
type ComicChapters []*ComicChapter
