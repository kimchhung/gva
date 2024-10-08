// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/genre"
)

// Genre is the model entity for the Genre schema.
type Genre struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name" rql:"column=name,filter,sort"`
	// NameID holds the value of the "name_id" field.
	NameID string `json:"name" rql:"column=name,filter,sort"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GenreQuery when eager-loading is set.
	Edges        GenreEdges `json:"edges" rql:"-"`
	selectValues sql.SelectValues
}

// GenreEdges holds the relations/edges for other nodes in the graph.
type GenreEdges struct {
	// Mangas holds the value of the mangas edge.
	Mangas []*Manga `json:"mangas,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedMangas map[string][]*Manga
}

// MangasOrErr returns the Mangas value or an error if the edge
// was not loaded in eager-loading.
func (e GenreEdges) MangasOrErr() ([]*Manga, error) {
	if e.loadedTypes[0] {
		return e.Mangas, nil
	}
	return nil, &NotLoadedError{edge: "mangas"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Genre) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case genre.FieldID:
			values[i] = new(pxid.ID)
		case genre.FieldIsEnable:
			values[i] = new(sql.NullBool)
		case genre.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case genre.FieldName, genre.FieldNameID:
			values[i] = new(sql.NullString)
		case genre.FieldCreatedAt, genre.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Genre fields.
func (ge *Genre) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case genre.FieldID:
			if value, ok := values[i].(*pxid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ge.ID = *value
			}
		case genre.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ge.CreatedAt = value.Time
			}
		case genre.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ge.UpdatedAt = value.Time
			}
		case genre.FieldIsEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enable", values[i])
			} else if value.Valid {
				ge.IsEnable = value.Bool
			}
		case genre.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ge.DeletedAt = int(value.Int64)
			}
		case genre.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ge.Name = value.String
			}
		case genre.FieldNameID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_id", values[i])
			} else if value.Valid {
				ge.NameID = value.String
			}
		default:
			ge.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Genre.
// This includes values selected through modifiers, order, etc.
func (ge *Genre) Value(name string) (ent.Value, error) {
	return ge.selectValues.Get(name)
}

// QueryMangas queries the "mangas" edge of the Genre entity.
func (ge *Genre) QueryMangas() *MangaQuery {
	return NewGenreClient(ge.config).QueryMangas(ge)
}

// Update returns a builder for updating this Genre.
// Note that you need to call Genre.Unwrap() before calling this method if this Genre
// was returned from a transaction, and the transaction was committed or rolled back.
func (ge *Genre) Update() *GenreUpdateOne {
	return NewGenreClient(ge.config).UpdateOne(ge)
}

// Unwrap unwraps the Genre entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ge *Genre) Unwrap() *Genre {
	_tx, ok := ge.config.driver.(*txDriver)
	if !ok {
		panic("ent: Genre is not a transactional entity")
	}
	ge.config.driver = _tx.drv
	return ge
}

// String implements the fmt.Stringer.
func (ge *Genre) String() string {
	var builder strings.Builder
	builder.WriteString("Genre(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ge.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ge.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ge.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_enable=")
	builder.WriteString(fmt.Sprintf("%v", ge.IsEnable))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ge.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ge.Name)
	builder.WriteString(", ")
	builder.WriteString("name_id=")
	builder.WriteString(ge.NameID)
	builder.WriteByte(')')
	return builder.String()
}

// NamedMangas returns the Mangas named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ge *Genre) NamedMangas(name string) ([]*Manga, error) {
	if ge.Edges.namedMangas == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ge.Edges.namedMangas[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ge *Genre) appendNamedMangas(name string, edges ...*Manga) {
	if ge.Edges.namedMangas == nil {
		ge.Edges.namedMangas = make(map[string][]*Manga)
	}
	if len(edges) == 0 {
		ge.Edges.namedMangas[name] = []*Manga{}
	} else {
		ge.Edges.namedMangas[name] = append(ge.Edges.namedMangas[name], edges...)
	}
}

// Genres is a parsable slice of Genre.
type Genres []*Genre
