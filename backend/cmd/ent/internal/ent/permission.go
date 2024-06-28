// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/permission"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-" rql:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id" rql:"filter,sort"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges        PermissionEdges `json:"edges" rql:"-"`
	selectValues sql.SelectValues
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedRoles map[string][]*Role
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			values[i] = new(pulid.ID)
		case permission.FieldOrder:
			values[i] = new(sql.NullInt64)
		case permission.FieldGroup, permission.FieldName, permission.FieldKey:
			values[i] = new(sql.NullString)
		case permission.FieldCreatedAt, permission.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case permission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case permission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case permission.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				pe.Group = value.String
			}
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				pe.Key = value.String
			}
		case permission.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				pe.Order = int(value.Int64)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Permission entity.
func (pe *Permission) QueryRoles() *RoleQuery {
	return NewPermissionClient(pe.config).QueryRoles(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(pe.Group)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(pe.Key)
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", pe.Order))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRoles returns the Roles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Permission) NamedRoles(name string) ([]*Role, error) {
	if pe.Edges.namedRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Permission) appendNamedRoles(name string, edges ...*Role) {
	if pe.Edges.namedRoles == nil {
		pe.Edges.namedRoles = make(map[string][]*Role)
	}
	if len(edges) == 0 {
		pe.Edges.namedRoles[name] = []*Role{}
	} else {
		pe.Edges.namedRoles[name] = append(pe.Edges.namedRoles[name], edges...)
	}
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
