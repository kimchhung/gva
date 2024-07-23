package permission

const (
	DepartmentGroup PermissionGroup = "department"
)

var (
	DepartmentSuper  = newKey(DepartmentGroup, ActionSuper)
	DepartmentView   = newKey(DepartmentGroup, ActionView)
	DepartmentAdd    = newKey(DepartmentGroup, ActionAdd)
	DepartmentEdit   = newKey(DepartmentGroup, ActionEdit)
	DepartmentDelete = newKey(DepartmentGroup, ActionDelete)

	DepartmentSeeder = NewSeeder(DepartmentGroup,
		DepartmentSuper,
		DepartmentView,
		DepartmentAdd,
		DepartmentEdit,
		DepartmentDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, DepartmentSeeder)
}
