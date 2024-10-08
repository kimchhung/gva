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
	// GenresColumns holds the columns for the "genres" table.
	GenresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "name", Type: field.TypeString},
		{Name: "name_id", Type: field.TypeString},
	}
	// GenresTable holds the schema information for the "genres" table.
	GenresTable = &schema.Table{
		Name:       "genres",
		Columns:    GenresColumns,
		PrimaryKey: []*schema.Column{GenresColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "genre_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{GenresColumns[4]},
			},
			{
				Name:    "genre_name_id_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{GenresColumns[6], GenresColumns[4]},
			},
		},
	}
	// MangasColumns holds the columns for the "mangas" table.
	MangasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_enable", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeInt, Default: "0"},
		{Name: "name_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
		{Name: "prodiver", Type: field.TypeString},
		{Name: "thumbnail_url", Type: field.TypeString},
		{Name: "authors", Type: field.TypeJSON},
	}
	// MangasTable holds the schema information for the "mangas" table.
	MangasTable = &schema.Table{
		Name:       "mangas",
		Columns:    MangasColumns,
		PrimaryKey: []*schema.Column{MangasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "manga_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{MangasColumns[4]},
			},
			{
				Name:    "manga_name_name_id_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{MangasColumns[6], MangasColumns[5], MangasColumns[4]},
			},
		},
	}
	// MangaChaptersColumns holds the columns for the "manga_chapters" table.
	MangaChaptersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "img_url", Type: field.TypeString},
		{Name: "number", Type: field.TypeUint},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "chapter_updated_at", Type: field.TypeTime},
		{Name: "manga_id", Type: field.TypeString},
	}
	// MangaChaptersTable holds the schema information for the "manga_chapters" table.
	MangaChaptersTable = &schema.Table{
		Name:       "manga_chapters",
		Columns:    MangaChaptersColumns,
		PrimaryKey: []*schema.Column{MangaChaptersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "manga_chapters_mangas_chapters",
				Columns:    []*schema.Column{MangaChaptersColumns[8]},
				RefColumns: []*schema.Column{MangasColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "mangachapter_provider_name_manga_id_number",
				Unique:  true,
				Columns: []*schema.Column{MangaChaptersColumns[6], MangaChaptersColumns[8], MangaChaptersColumns[5]},
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
	// GenreMangasColumns holds the columns for the "genre_mangas" table.
	GenreMangasColumns = []*schema.Column{
		{Name: "genre_id", Type: field.TypeString},
		{Name: "manga_id", Type: field.TypeString},
	}
	// GenreMangasTable holds the schema information for the "genre_mangas" table.
	GenreMangasTable = &schema.Table{
		Name:       "genre_mangas",
		Columns:    GenreMangasColumns,
		PrimaryKey: []*schema.Column{GenreMangasColumns[0], GenreMangasColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "genre_mangas_genre_id",
				Columns:    []*schema.Column{GenreMangasColumns[0]},
				RefColumns: []*schema.Column{GenresColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "genre_mangas_manga_id",
				Columns:    []*schema.Column{GenreMangasColumns[1]},
				RefColumns: []*schema.Column{MangasColumns[0]},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		DepartmentsTable,
		GenresTable,
		MangasTable,
		MangaChaptersTable,
		PermissionsTable,
		RolesTable,
		AdminRolesTable,
		GenreMangasTable,
		RolePermissionsTable,
	}
)

func init() {
	AdminsTable.ForeignKeys[0].RefTable = DepartmentsTable
	DepartmentsTable.ForeignKeys[0].RefTable = DepartmentsTable
	MangaChaptersTable.ForeignKeys[0].RefTable = MangasTable
	AdminRolesTable.ForeignKeys[0].RefTable = AdminsTable
	AdminRolesTable.ForeignKeys[1].RefTable = RolesTable
	GenreMangasTable.ForeignKeys[0].RefTable = GenresTable
	GenreMangasTable.ForeignKeys[1].RefTable = MangasTable
	RolePermissionsTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
}
