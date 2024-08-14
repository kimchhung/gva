// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "whitelist_ips", Type: field.TypeJSON},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "department_id", Type: field.TypeString, Nullable: true},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "admins_departments_members",
				Columns:    []*schema.Column{AdminsColumns[9]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "admin_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{AdminsColumns[4]},
			},
			{
				Name:    "admin_username_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{AdminsColumns[5], AdminsColumns[4]},
			},
		},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "name_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "pid", Type: field.TypeString, Nullable: true},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "departments_departments_children",
				Columns:    []*schema.Column{DepartmentsColumns[7]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "department_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{DepartmentsColumns[3]},
			},
			{
				Name:    "department_name_id_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{DepartmentsColumns[5], DepartmentsColumns[3]},
			},
		},
	}
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "path", Type: field.TypeString},
		{Name: "component", Type: field.TypeString},
		{Name: "redirect", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"cata_log", "menu", "button", "external_link"}, Default: "cata_log"},
		{Name: "meta", Type: field.TypeJSON},
		{Name: "pid", Type: field.TypeString, Nullable: true},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menus_menus_children",
				Columns:    []*schema.Column{MenusColumns[12]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "menu_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{MenusColumns[4]},
			},
		},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "group", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "scope", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"dynamic", "static"}, Default: "dynamic"},
		{Name: "order", Type: field.TypeInt, Nullable: true, Default: 0},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
	}
	// RegionsColumns holds the columns for the "regions" table.
	RegionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "name_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"continent", "country", "city", "street", "any"}},
		{Name: "pid", Type: field.TypeString, Nullable: true},
	}
	// RegionsTable holds the schema information for the "regions" table.
	RegionsTable = &schema.Table{
		Name:       "regions",
		Columns:    RegionsColumns,
		PrimaryKey: []*schema.Column{RegionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "regions_regions_children",
				Columns:    []*schema.Column{RegionsColumns[8]},
				RefColumns: []*schema.Column{RegionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "region_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{RegionsColumns[3]},
			},
			{
				Name:    "region_name_id_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{RegionsColumns[5], RegionsColumns[3]},
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt},
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
				Columns: []*schema.Column{RolesColumns[4]},
			},
		},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "name", Type: field.TypeString},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "todo_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{TodosColumns[3]},
			},
		},
	}
	// AdminRolesColumns holds the columns for the "admin_roles" table.
	AdminRolesColumns = []*schema.Column{
		{Name: "admin_id", Type: field.TypeString},
		{Name: "role_id", Type: field.TypeString},
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
		{Name: "role_id", Type: field.TypeString},
		{Name: "permission_id", Type: field.TypeString},
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
	// RoleRoutesColumns holds the columns for the "role_routes" table.
	RoleRoutesColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeString},
		{Name: "menu_id", Type: field.TypeString},
	}
	// RoleRoutesTable holds the schema information for the "role_routes" table.
	RoleRoutesTable = &schema.Table{
		Name:       "role_routes",
		Columns:    RoleRoutesColumns,
		PrimaryKey: []*schema.Column{RoleRoutesColumns[0], RoleRoutesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_routes_role_id",
				Columns:    []*schema.Column{RoleRoutesColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_routes_menu_id",
				Columns:    []*schema.Column{RoleRoutesColumns[1]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		DepartmentsTable,
		MenusTable,
		PermissionsTable,
		RegionsTable,
		RolesTable,
		TodosTable,
		AdminRolesTable,
		RolePermissionsTable,
		RoleRoutesTable,
	}
)

func init() {
	AdminsTable.ForeignKeys[0].RefTable = DepartmentsTable
	DepartmentsTable.ForeignKeys[0].RefTable = DepartmentsTable
	MenusTable.ForeignKeys[0].RefTable = MenusTable
	RegionsTable.ForeignKeys[0].RefTable = RegionsTable
	AdminRolesTable.ForeignKeys[0].RefTable = AdminsTable
	AdminRolesTable.ForeignKeys[1].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	RoleRoutesTable.ForeignKeys[0].RefTable = RolesTable
	RoleRoutesTable.ForeignKeys[1].RefTable = MenusTable
}