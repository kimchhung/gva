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

	Blog       string = "blogs"
	Todo       string = "todos"
	TronEnergy string = "tron_energys" // #inject:tableName (do not remove or change position this comment, it is used by the code generator)
)
