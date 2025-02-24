package permission

const (
	MenuGroup PermissionGroup = "menu"
)

var (
	MenuSuper  = newKey(MenuGroup, ActionSuper)
	MenuView   = newKey(MenuGroup, ActionView)
	MenuAdd    = newKey(MenuGroup, ActionAdd)
	MenuEdit   = newKey(MenuGroup, ActionEdit)
	MenuDelete = newKey(MenuGroup, ActionDelete)

	MenuSeeder = NewSeeder(MenuGroup,
		MenuSuper,
		MenuView,
		MenuAdd,
		MenuEdit,
		MenuDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, MenuSeeder)
}
