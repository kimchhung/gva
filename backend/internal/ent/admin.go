// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kimchhung/gva/backend/internal/ent/admin"
)

// Admin is the model entity for the Admin schema.
type Admin struct {
	config `json:"-" rql:"-"`
	// ID of the ent.
	ID int `json:"id" rql:"filter,sort"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" rql:"filter,sort"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// IsEnable holds the value of the "is_enable" field.
	IsEnable bool `json:"isEnable"  rql:"filter,sort"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt int `json:"-"`
	// Username holds the value of the "username" field.
	Username string `json:"username" rql:"column=username,filter,sort"`
	// Password holds the value of the "password" field.
	Password string `json:"-" rql:"-"`
	// WhitelistIps holds the value of the "whitelist_ips" field.
	WhitelistIps []string `json:"whitelistIps"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"displayName,omitempty" rql:"filter,sort"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdminQuery when eager-loading is set.
	Edges        AdminEdges `json:"edges" rql:"-"`
	selectValues sql.SelectValues
}

// AdminEdges holds the relations/edges for other nodes in the graph.
type AdminEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e AdminEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Admin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case admin.FieldWhitelistIps:
			values[i] = new([]byte)
		case admin.FieldIsEnable:
			values[i] = new(sql.NullBool)
		case admin.FieldID, admin.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case admin.FieldUsername, admin.FieldPassword, admin.FieldDisplayName:
			values[i] = new(sql.NullString)
		case admin.FieldCreatedAt, admin.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Admin fields.
func (a *Admin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case admin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case admin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case admin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case admin.FieldIsEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enable", values[i])
			} else if value.Valid {
				a.IsEnable = value.Bool
			}
		case admin.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = int(value.Int64)
			}
		case admin.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case admin.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case admin.FieldWhitelistIps:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field whitelist_ips", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.WhitelistIps); err != nil {
					return fmt.Errorf("unmarshal field whitelist_ips: %w", err)
				}
			}
		case admin.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				a.DisplayName = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Admin.
// This includes values selected through modifiers, order, etc.
func (a *Admin) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Admin entity.
func (a *Admin) QueryRoles() *RoleQuery {
	return NewAdminClient(a.config).QueryRoles(a)
}

// Update returns a builder for updating this Admin.
// Note that you need to call Admin.Unwrap() before calling this method if this Admin
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Admin) Update() *AdminUpdateOne {
	return NewAdminClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Admin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Admin) Unwrap() *Admin {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Admin is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Admin) String() string {
	var builder strings.Builder
	builder.WriteString("Admin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_enable=")
	builder.WriteString(fmt.Sprintf("%v", a.IsEnable))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("whitelist_ips=")
	builder.WriteString(fmt.Sprintf("%v", a.WhitelistIps))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(a.DisplayName)
	builder.WriteByte(')')
	return builder.String()
}

// Admins is a parsable slice of Admin.
type Admins []*Admin
