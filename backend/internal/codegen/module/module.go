package module_template

var special = "`"
var Module = `package {{.EntityAllLower}}

import (
	"github.com/gva/internal/echoc"
	"github.com/gva/app/common/controller"

	"go.uber.org/fx"
)

// Register bulkly
var New{{.EntityPascal}}Module = fx.Module("{{.EntityPascal}}Module",
// Register Repository & Service
fx.Provide(repository.New{{.EntityPascal}}Repository),
fx.Provide(service.New{{.EntityPascal}}Service),

// Regiser Controller
fx.Provide(
	fx.Annotate(
		controller.New{{.EntityPascal}}Controller,
		fx.As(new(echoc.Controller)),
		fx.ResultTags(controller.APIAdminControllerGroup),
	),
),
)
`
