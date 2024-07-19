package module_template

var special = "`"
var Module = `package {{.EntityAllLower}}

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var {{.EntityPascal}}Module = fx.Module("{{.EntityPascal}}Module",
	// Register Service
	fx.Provide(New{{.EntityPascal}}Service),

	// Regiser Controller
	controller.ProvideAdminController(New{{.EntityPascal}}Controller),
)
`
