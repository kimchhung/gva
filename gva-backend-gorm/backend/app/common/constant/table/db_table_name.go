package table

const (
	// Role and Permission
	Role      string = "roles"
	AdminRole string = "admin_roles"

	AdminUser string = "admin_users"
	Admin     string = "admins"

	// System
	Configuration string = "configurations"

	//
	Permission   string = "permissions"
	OperationLog string = "operation_logs"

	Blog string = "blogs" // #inject:tableName (do not remove or change position this comment, it is used by the code generator)
)
