package module_template

import "fmt"

var special = "`"
var Module = fmt.Sprintf(
	`package {{.EntitySnake}}

import (
	"github.com/kimchhung/gva/extra/app/module/dashboard/{{.EntitySnake}}/controller"
	"github.com/kimchhung/gva/extra/app/module/dashboard/{{.EntitySnake}}/repository"
	"github.com/kimchhung/gva/extra/app/module/dashboard/{{.EntitySnake}}/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

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
			controller.NewTodoController,
			fx.As(new(rctrl.Controller)),
			fx.ResultTags(%sgroup:"controllers"%s),
		),
	),
)
`, special, special,
)
