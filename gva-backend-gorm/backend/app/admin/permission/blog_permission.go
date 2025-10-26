package permission

const (
	BlogGroup TPermissionGroup = "blog"
)

var (
	BlogSuper  = newKey(BlogGroup, ActionSuper)
	BlogView   = newKey(BlogGroup, ActionView)
	BlogAdd    = newKey(BlogGroup, ActionAdd)
	BlogEdit   = newKey(BlogGroup, ActionEdit)
	BlogDelete = newKey(BlogGroup, ActionDelete)

	BlogSeeder = NewSeeder(BlogGroup,
		BlogSuper,
		BlogView,
		BlogAdd,
		BlogEdit,
		BlogDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, BlogSeeder)
}
