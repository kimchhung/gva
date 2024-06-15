// Code generated by ent, DO NOT EDIT.

package role

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldIsEnable holds the string denoting the is_enable field in the database.
	FieldIsEnable = "is_enable"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldIsChangeable holds the string denoting the is_changeable field in the database.
	FieldIsChangeable = "is_changeable"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgeRoutes holds the string denoting the routes edge name in mutations.
	EdgeRoutes = "routes"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// AdminsTable is the table that holds the admins relation/edge. The primary key declared below.
	AdminsTable = "admin_roles"
	// AdminsInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminsInverseTable = "admins"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "role_permissions"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permissions"
	// RoutesTable is the table that holds the routes relation/edge. The primary key declared below.
	RoutesTable = "role_routes"
	// RoutesInverseTable is the table name for the Route entity.
	// It exists in this package in order to avoid circular dependency with the "route" package.
	RoutesInverseTable = "routes"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIsEnable,
	FieldDeletedAt,
	FieldName,
	FieldDescription,
	FieldOrder,
	FieldIsChangeable,
}

var (
	// AdminsPrimaryKey and AdminsColumn2 are the table columns denoting the
	// primary key for the admins relation (M2M).
	AdminsPrimaryKey = []string{"admin_id", "role_id"}
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"role_id", "permission_id"}
	// RoutesPrimaryKey and RoutesColumn2 are the table columns denoting the
	// primary key for the routes relation (M2M).
	RoutesPrimaryKey = []string{"role_id", "route_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/kimchhung/gva/backend/internal/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsEnable holds the default value on creation for the "is_enable" field.
	DefaultIsEnable bool
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt int
)

// OrderOption defines the ordering options for the Role queries.
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

// ByIsEnable orders the results by the is_enable field.
func ByIsEnable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsEnable, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByIsChangeable orders the results by the is_changeable field.
func ByIsChangeable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsChangeable, opts...).ToFunc()
}

// ByAdminsCount orders the results by admins count.
func ByAdminsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdminsStep(), opts...)
	}
}

// ByAdmins orders the results by admins terms.
func ByAdmins(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPermissionsCount orders the results by permissions count.
func ByPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionsStep(), opts...)
	}
}

// ByPermissions orders the results by permissions terms.
func ByPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoutesCount orders the results by routes count.
func ByRoutesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutesStep(), opts...)
	}
}

// ByRoutes orders the results by routes terms.
func ByRoutes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAdminsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdminsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AdminsTable, AdminsPrimaryKey...),
	)
}
func newPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
	)
}
func newRoutesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, RoutesTable, RoutesPrimaryKey...),
	)
}
