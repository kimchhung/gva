// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "whitelist_ips", Type: field.TypeJSON},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "admin_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{AdminsColumns[3]},
			},
			{
				Name:    "admin_username_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{AdminsColumns[4], AdminsColumns[3]},
			},
		},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "name", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "is_changeable", Type: field.TypeBool},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{RolesColumns[3]},
			},
		},
	}
	// AdminRolesColumns holds the columns for the "admin_roles" table.
	AdminRolesColumns = []*schema.Column{
		{Name: "admin_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// AdminRolesTable holds the schema information for the "admin_roles" table.
	AdminRolesTable = &schema.Table{
		Name:       "admin_roles",
		Columns:    AdminRolesColumns,
		PrimaryKey: []*schema.Column{AdminRolesColumns[0], AdminRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "admin_roles_admin_id",
				Columns:    []*schema.Column{AdminRolesColumns[0]},
				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "admin_roles_role_id",
				Columns:    []*schema.Column{AdminRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RolePermissionsColumns holds the columns for the "role_permissions" table.
	RolePermissionsColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "permission_id", Type: field.TypeInt},
	}
	// RolePermissionsTable holds the schema information for the "role_permissions" table.
	RolePermissionsTable = &schema.Table{
		Name:       "role_permissions",
		Columns:    RolePermissionsColumns,
		PrimaryKey: []*schema.Column{RolePermissionsColumns[0], RolePermissionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permissions_role_id",
				Columns:    []*schema.Column{RolePermissionsColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_permissions_permission_id",
				Columns:    []*schema.Column{RolePermissionsColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		PermissionsTable,
		RolesTable,
		AdminRolesTable,
		RolePermissionsTable,
	}
)

func init() {
	AdminRolesTable.ForeignKeys[0].RefTable = AdminsTable
	AdminRolesTable.ForeignKeys[1].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
}
