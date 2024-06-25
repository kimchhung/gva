package module_template

import "fmt"

var special = "`"
var Module = fmt.Sprintf(
	`package {{.EntitySnake}}

import (
	"github.com/gva/api/admin/module/{{.EntitySnake}}/controller"
	"github.com/gva/api/admin/module/{{.EntitySnake}}/repository"
	"github.com/gva/api/admin/module/{{.EntitySnake}}/service"
	"github.com/gva/internal/rctrl"
	"github.com/gva/app/constant"

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
			fx.ResultTags(constant.APIAdminControllerGroup),
		),
	),
)
`, special, special,
)
