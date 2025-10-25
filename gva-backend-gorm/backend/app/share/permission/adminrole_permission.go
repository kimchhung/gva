package permission

const (
	AdminRoleGroup TPermissionGroup = "admin_role"
)

var (
	AdminRoleSuper  = newKey(AdminRoleGroup, ActionSuper)
	AdminRoleView   = newKey(AdminRoleGroup, ActionView)
	AdminRoleAdd    = newKey(AdminRoleGroup, ActionAdd)
	AdminRoleEdit   = newKey(AdminRoleGroup, ActionEdit)
	AdminRoleDelete = newKey(AdminRoleGroup, ActionDelete)

	AdminRoleSeeder = NewSeeder(AdminRoleGroup,
		AdminRoleSuper,
		AdminRoleView,
		AdminRoleAdd,
		AdminRoleEdit,
		AdminRoleDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, AdminRoleSeeder)
}
