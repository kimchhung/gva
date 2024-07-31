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
	"github.com/gva/app/database/schema/types"
	"github.com/gva/internal/ent/menu"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
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
	// Pid holds the value of the "pid" field.
	Pid *pxid.ID `json:"pid,omitempty" rql:"filter,sort"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty" rql:"filter,sort"`
	// Component holds the value of the "component" field.
	Component string `json:"component,omitempty" rql:"filter,sort"`
	// Redirect holds the value of the "redirect" field.
	Redirect *string `json:"redirect,omitempty" rql:"filter,sort"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" rql:"filter,sort"`
	// Order holds the value of the "order" field.
	Order int `json:"order" rql:"filter,sort"`
	// Type holds the value of the "type" field.
	Type menu.Type `json:"type,omitempty" rql:"filter,sort"`
	// Meta holds the value of the "meta" field.
	Meta types.MenuMeta `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges" rql:"-"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Menu `json:"children,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedChildren map[string][]*Menu
	namedRoles    map[string][]*Role
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ChildrenOrErr() ([]*Menu, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[2] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldPid:
			values[i] = &sql.NullScanner{S: new(pxid.ID)}
		case menu.FieldMeta:
			values[i] = new([]byte)
		case menu.FieldID:
			values[i] = new(pxid.ID)
		case menu.FieldIsEnable:
			values[i] = new(sql.NullBool)
		case menu.FieldDeletedAt, menu.FieldOrder:
			values[i] = new(sql.NullInt64)
		case menu.FieldPath, menu.FieldComponent, menu.FieldRedirect, menu.FieldName, menu.FieldType:
			values[i] = new(sql.NullString)
		case menu.FieldCreatedAt, menu.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			if value, ok := values[i].(*pxid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case menu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case menu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case menu.FieldIsEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enable", values[i])
			} else if value.Valid {
				m.IsEnable = value.Bool
			}
		case menu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = int(value.Int64)
			}
		case menu.FieldPid:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				m.Pid = new(pxid.ID)
				*m.Pid = *value.S.(*pxid.ID)
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = value.String
			}
		case menu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				m.Component = value.String
			}
		case menu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				m.Redirect = new(string)
				*m.Redirect = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				m.Order = int(value.Int64)
			}
		case menu.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = menu.Type(value.String)
			}
		case menu.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Menu entity.
func (m *Menu) QueryParent() *MenuQuery {
	return NewMenuClient(m.config).QueryParent(m)
}

// QueryChildren queries the "children" edge of the Menu entity.
func (m *Menu) QueryChildren() *MenuQuery {
	return NewMenuClient(m.config).QueryChildren(m)
}

// QueryRoles queries the "roles" edge of the Menu entity.
func (m *Menu) QueryRoles() *RoleQuery {
	return NewMenuClient(m.config).QueryRoles(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
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
	if v := m.Pid; v != nil {
		builder.WriteString("pid=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(m.Path)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(m.Component)
	builder.WriteString(", ")
	if v := m.Redirect; v != nil {
		builder.WriteString("redirect=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", m.Order))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", m.Type))
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(fmt.Sprintf("%v", m.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Menu) NamedChildren(name string) ([]*Menu, error) {
	if m.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Menu) appendNamedChildren(name string, edges ...*Menu) {
	if m.Edges.namedChildren == nil {
		m.Edges.namedChildren = make(map[string][]*Menu)
	}
	if len(edges) == 0 {
		m.Edges.namedChildren[name] = []*Menu{}
	} else {
		m.Edges.namedChildren[name] = append(m.Edges.namedChildren[name], edges...)
	}
}

// NamedRoles returns the Roles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Menu) NamedRoles(name string) ([]*Role, error) {
	if m.Edges.namedRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Menu) appendNamedRoles(name string, edges ...*Role) {
	if m.Edges.namedRoles == nil {
		m.Edges.namedRoles = make(map[string][]*Role)
	}
	if len(edges) == 0 {
		m.Edges.namedRoles[name] = []*Role{}
	} else {
		m.Edges.namedRoles[name] = append(m.Edges.namedRoles[name], edges...)
	}
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
