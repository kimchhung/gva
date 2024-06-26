// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "whitelist_ips", Type: field.TypeJSON},
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
				Columns: []*schema.Column{AdminsColumns[4]},
			},
			{
				Name:    "admin_username_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{AdminsColumns[5], AdminsColumns[4]},
			},
		},
	}
	// ComicsColumns holds the columns for the "comics" table.
	ComicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chapter", Type: field.TypeUint},
		{Name: "title", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "covers", Type: field.TypeJSON},
		{Name: "status", Type: field.TypeString},
		{Name: "is_translate_completed", Type: field.TypeBool, Default: false},
		{Name: "up_count", Type: field.TypeUint, Default: 0},
		{Name: "last_chapter_id", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "final_chapter_id", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
	}
	// ComicsTable holds the schema information for the "comics" table.
	ComicsTable = &schema.Table{
		Name:       "comics",
		Columns:    ComicsColumns,
		PrimaryKey: []*schema.Column{ComicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comics_comic_chapters_last_chapter",
				Columns:    []*schema.Column{ComicsColumns[10]},
				RefColumns: []*schema.Column{ComicChaptersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comics_comic_chapters_final_chapter",
				Columns:    []*schema.Column{ComicsColumns[11]},
				RefColumns: []*schema.Column{ComicChaptersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ComicChaptersColumns holds the columns for the "comic_chapters" table.
	ComicChaptersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chapter", Type: field.TypeUint},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "volumn", Type: field.TypeString, Nullable: true},
		{Name: "lang", Type: field.TypeString},
		{Name: "up_count", Type: field.TypeUint, Default: 0},
		{Name: "down_count", Type: field.TypeUint, Default: 0},
		{Name: "is_last_chapter", Type: field.TypeBool, Default: false},
		{Name: "comic_chapters", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
	}
	// ComicChaptersTable holds the schema information for the "comic_chapters" table.
	ComicChaptersTable = &schema.Table{
		Name:       "comic_chapters",
		Columns:    ComicChaptersColumns,
		PrimaryKey: []*schema.Column{ComicChaptersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comic_chapters_comics_chapters",
				Columns:    []*schema.Column{ComicChaptersColumns[10]},
				RefColumns: []*schema.Column{ComicsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ComicImgsColumns holds the columns for the "comic_imgs" table.
	ComicImgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "b2key", Type: field.TypeString, Unique: true},
		{Name: "height", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "optimized_size", Type: field.TypeInt64},
		{Name: "size", Type: field.TypeInt64},
		{Name: "width", Type: field.TypeInt},
		{Name: "comic_chapter_imgs", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
	}
	// ComicImgsTable holds the schema information for the "comic_imgs" table.
	ComicImgsTable = &schema.Table{
		Name:       "comic_imgs",
		Columns:    ComicImgsColumns,
		PrimaryKey: []*schema.Column{ComicImgsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comic_imgs_comic_chapters_imgs",
				Columns:    []*schema.Column{ComicImgsColumns[9]},
				RefColumns: []*schema.Column{ComicChaptersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GenresColumns holds the columns for the "genres" table.
	GenresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"comic"}},
	}
	// GenresTable holds the schema information for the "genres" table.
	GenresTable = &schema.Table{
		Name:       "genres",
		Columns:    GenresColumns,
		PrimaryKey: []*schema.Column{GenresColumns[0]},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
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
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
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
	// RoutesColumns holds the columns for the "routes" table.
	RoutesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
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
		{Name: "parent_id", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
	}
	// RoutesTable holds the schema information for the "routes" table.
	RoutesTable = &schema.Table{
		Name:       "routes",
		Columns:    RoutesColumns,
		PrimaryKey: []*schema.Column{RoutesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routes_routes_children",
				Columns:    []*schema.Column{RoutesColumns[12]},
				RefColumns: []*schema.Column{RoutesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "route_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{RoutesColumns[4]},
			},
			{
				Name:    "idx_uniq_route",
				Unique:  true,
				Columns: []*schema.Column{RoutesColumns[12], RoutesColumns[5], RoutesColumns[10], RoutesColumns[4]},
			},
		},
	}
	// AdminRolesColumns holds the columns for the "admin_roles" table.
	AdminRolesColumns = []*schema.Column{
		{Name: "admin_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "role_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
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
		{Name: "role_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "permission_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
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
		{Name: "role_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
		{Name: "route_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(21)"}},
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
				Symbol:     "role_routes_route_id",
				Columns:    []*schema.Column{RoleRoutesColumns[1]},
				RefColumns: []*schema.Column{RoutesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		ComicsTable,
		ComicChaptersTable,
		ComicImgsTable,
		GenresTable,
		PermissionsTable,
		RolesTable,
		RoutesTable,
		AdminRolesTable,
		RolePermissionsTable,
		RoleRoutesTable,
	}
)

func init() {
	ComicsTable.ForeignKeys[0].RefTable = ComicChaptersTable
	ComicsTable.ForeignKeys[1].RefTable = ComicChaptersTable
	ComicChaptersTable.ForeignKeys[0].RefTable = ComicsTable
	ComicImgsTable.ForeignKeys[0].RefTable = ComicChaptersTable
	RoutesTable.ForeignKeys[0].RefTable = RoutesTable
	AdminRolesTable.ForeignKeys[0].RefTable = AdminsTable
	AdminRolesTable.ForeignKeys[1].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	RoleRoutesTable.ForeignKeys[0].RefTable = RolesTable
	RoleRoutesTable.ForeignKeys[1].RefTable = RoutesTable
}
