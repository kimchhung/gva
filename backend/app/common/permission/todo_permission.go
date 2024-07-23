package permission

const (
	todoGroup PermissionGroup = "todo"
)

var (
	todoSuper  = newKey(todoGroup, ActionSuper)
	todoView   = newKey(todoGroup, ActionView)
	todoAdd    = newKey(todoGroup, ActionAdd)
	todoEdit   = newKey(todoGroup, ActionEdit)
	todoDelete = newKey(todoGroup, ActionDelete)

	todoSeeder = NewSeeder(todoGroup,
		todoSuper,
		todoView,
		todoAdd,
		todoEdit,
		todoDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, todoSeeder)
}
