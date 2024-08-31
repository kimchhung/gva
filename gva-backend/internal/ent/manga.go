// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/manga"
)

// Manga is the model entity for the Manga schema.
type Manga struct {
	config `json:"-" rql:"-"`
	// ID of the ent.
	ID pxid.ID `json:"id" rql:"filter,sort"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// IsEnable holds the value of the "is_enable" field.
	IsEnable bool `json:"isEnable"  rql:"filter,sort"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt int `json:"-"`
	// NameID holds the value of the "name_id" field.
	NameID string `json:"name_id" rql:"column=name_id,filter,sort"`
	// Name holds the value of the "name" field.
	Name string `json:"name" rql:"column=name,filter,sort"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc"`
	// Prodiver holds the value of the "prodiver" field.
	Prodiver string `json:"provider"`
	// ThumbnailURL holds the value of the "thumbnail_url" field.
	ThumbnailURL string `json:"thumbnailUrl"`
	// Authors holds the value of the "authors" field.
	Authors []string `json:"authors"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MangaQuery when eager-loading is set.
	Edges        MangaEdges `json:"edges" rql:"-"`
	selectValues sql.SelectValues
}

// MangaEdges holds the relations/edges for other nodes in the graph.
type MangaEdges struct {
	// Chapters holds the value of the chapters edge.
	Chapters []*MangaChapter `json:"chapters,omitempty"`
	// Genres holds the value of the genres edge.
	Genres []*Genre `json:"genres,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedChapters map[string][]*MangaChapter
	namedGenres   map[string][]*Genre
}

// ChaptersOrErr returns the Chapters value or an error if the edge
// was not loaded in eager-loading.
func (e MangaEdges) ChaptersOrErr() ([]*MangaChapter, error) {
	if e.loadedTypes[0] {
		return e.Chapters, nil
	}
	return nil, &NotLoadedError{edge: "chapters"}
}

// GenresOrErr returns the Genres value or an error if the edge
// was not loaded in eager-loading.
func (e MangaEdges) GenresOrErr() ([]*Genre, error) {
	if e.loadedTypes[1] {
		return e.Genres, nil
	}
	return nil, &NotLoadedError{edge: "genres"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Manga) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case manga.FieldAuthors:
			values[i] = new([]byte)
		case manga.FieldID:
			values[i] = new(pxid.ID)
		case manga.FieldIsEnable:
			values[i] = new(sql.NullBool)
		case manga.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case manga.FieldNameID, manga.FieldName, manga.FieldDesc, manga.FieldProdiver, manga.FieldThumbnailURL:
			values[i] = new(sql.NullString)
		case manga.FieldCreatedAt, manga.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Manga fields.
func (m *Manga) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case manga.FieldID:
			if value, ok := values[i].(*pxid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case manga.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case manga.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case manga.FieldIsEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enable", values[i])
			} else if value.Valid {
				m.IsEnable = value.Bool
			}
		case manga.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = int(value.Int64)
			}
		case manga.FieldNameID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_id", values[i])
			} else if value.Valid {
				m.NameID = value.String
			}
		case manga.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case manga.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				m.Desc = value.String
			}
		case manga.FieldProdiver:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prodiver", values[i])
			} else if value.Valid {
				m.Prodiver = value.String
			}
		case manga.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				m.ThumbnailURL = value.String
			}
		case manga.FieldAuthors:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field authors", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Authors); err != nil {
					return fmt.Errorf("unmarshal field authors: %w", err)
				}
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Manga.
// This includes values selected through modifiers, order, etc.
func (m *Manga) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryChapters queries the "chapters" edge of the Manga entity.
func (m *Manga) QueryChapters() *MangaChapterQuery {
	return NewMangaClient(m.config).QueryChapters(m)
}

// QueryGenres queries the "genres" edge of the Manga entity.
func (m *Manga) QueryGenres() *GenreQuery {
	return NewMangaClient(m.config).QueryGenres(m)
}

// Update returns a builder for updating this Manga.
// Note that you need to call Manga.Unwrap() before calling this method if this Manga
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Manga) Update() *MangaUpdateOne {
	return NewMangaClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Manga entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Manga) Unwrap() *Manga {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Manga is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Manga) String() string {
	var builder strings.Builder
	builder.WriteString("Manga(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_enable=")
	builder.WriteString(fmt.Sprintf("%v", m.IsEnable))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", m.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("name_id=")
	builder.WriteString(m.NameID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(m.Desc)
	builder.WriteString(", ")
	builder.WriteString("prodiver=")
	builder.WriteString(m.Prodiver)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(m.ThumbnailURL)
	builder.WriteString(", ")
	builder.WriteString("authors=")
	builder.WriteString(fmt.Sprintf("%v", m.Authors))
	builder.WriteByte(')')
	return builder.String()
}

// NamedChapters returns the Chapters named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Manga) NamedChapters(name string) ([]*MangaChapter, error) {
	if m.Edges.namedChapters == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedChapters[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Manga) appendNamedChapters(name string, edges ...*MangaChapter) {
	if m.Edges.namedChapters == nil {
		m.Edges.namedChapters = make(map[string][]*MangaChapter)
	}
	if len(edges) == 0 {
		m.Edges.namedChapters[name] = []*MangaChapter{}
	} else {
		m.Edges.namedChapters[name] = append(m.Edges.namedChapters[name], edges...)
	}
}

// NamedGenres returns the Genres named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Manga) NamedGenres(name string) ([]*Genre, error) {
	if m.Edges.namedGenres == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedGenres[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Manga) appendNamedGenres(name string, edges ...*Genre) {
	if m.Edges.namedGenres == nil {
		m.Edges.namedGenres = make(map[string][]*Genre)
	}
	if len(edges) == 0 {
		m.Edges.namedGenres[name] = []*Genre{}
	} else {
		m.Edges.namedGenres[name] = append(m.Edges.namedGenres[name], edges...)
	}
}

// Mangas is a parsable slice of Manga.
type Mangas []*Manga
