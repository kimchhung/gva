// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/genre"
)

// Genre is the model entity for the Genre schema.
type Genre struct {
	config `json:"-" rql:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id" rql:"filter,sort"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type         genre.Type `json:"type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Genre) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case genre.FieldID:
			values[i] = new(pulid.ID)
		case genre.FieldName, genre.FieldType:
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
			if value, ok := values[i].(*pulid.ID); !ok {
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
		case genre.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ge.Name = value.String
			}
		case genre.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ge.Type = genre.Type(value.String)
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
	builder.WriteString("name=")
	builder.WriteString(ge.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ge.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Genres is a parsable slice of Genre.
type Genres []*Genre
