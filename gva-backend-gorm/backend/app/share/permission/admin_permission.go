package permission

const (
	AdminGroup TPermissionGroup = "admin"
)

var (
	AdminSuper  = newKey(AdminGroup, ActionSuper)
	AdminView   = newKey(AdminGroup, ActionView)
	AdminAdd    = newKey(AdminGroup, ActionAdd)
	AdminEdit   = newKey(AdminGroup, ActionEdit)
	AdminDelete = newKey(AdminGroup, ActionDelete)

	AdminSeeder = NewSeeder(AdminGroup,
		AdminSuper,
		AdminView,
		AdminAdd,
		AdminEdit,
		AdminDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, AdminSeeder)
}
