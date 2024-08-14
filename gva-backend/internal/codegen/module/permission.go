package module_template

var Permission = `package permission

const (
	{{.EntityPascal}}Group PermissionGroup = "{{.EntitySnake}}"
)

var (
	{{.EntityPascal}}Super  = newKey({{.EntityPascal}}Group, ActionSuper)
	{{.EntityPascal}}View   = newKey({{.EntityPascal}}Group, ActionView)
	{{.EntityPascal}}Add    = newKey({{.EntityPascal}}Group, ActionAdd)
	{{.EntityPascal}}Edit   = newKey({{.EntityPascal}}Group, ActionEdit)
	{{.EntityPascal}}Delete = newKey({{.EntityPascal}}Group, ActionDelete)

	{{.EntityPascal}}Seeder = NewSeeder({{.EntityPascal}}Group,
		{{.EntityPascal}}Super,
		{{.EntityPascal}}View,
		{{.EntityPascal}}Add,
		{{.EntityPascal}}Edit,
		{{.EntityPascal}}Delete,
	)
)

func init() {
	allSeeders = append(allSeeders, {{.EntityPascal}}Seeder)
}
`
