package permission

const (
	TodoGroup TPermissionGroup = "todo"
)

var (
	TodoSuper  = newKey(TodoGroup, ActionSuper)
	TodoView   = newKey(TodoGroup, ActionView)
	TodoAdd    = newKey(TodoGroup, ActionAdd)
	TodoEdit   = newKey(TodoGroup, ActionEdit)
	TodoDelete = newKey(TodoGroup, ActionDelete)

	TodoSeeder = NewSeeder(TodoGroup,
		TodoSuper,
		TodoView,
		TodoAdd,
		TodoEdit,
		TodoDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, TodoSeeder)
}
