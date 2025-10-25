package module_template

var special = "`"
var Module = `package {{.EntityAllLower}}

import (
	adminrouter "backend/app/admin/router"
	"go.uber.org/fx"
)

// Register bulkly
var {{.EntityPascal}}Module = fx.Module("{{.EntityPascal}}Module",
	// Register Service
	fx.Provide(New{{.EntityPascal}}Service),

	// Regiser Controller
	adminrouter.Controller.Add(New{{.EntityPascal}}Controller),
)
`
