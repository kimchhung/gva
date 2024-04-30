package module_template

import "fmt"

var special = "`"
var Module = fmt.Sprintf(
	`package {{.EntitySnake}}

import (
	"github.com/kimchhung/gva/extra/api/admin/module/{{.EntitySnake}}/controller"
	"github.com/kimchhung/gva/extra/api/admin/module/{{.EntitySnake}}/repository"
	"github.com/kimchhung/gva/extra/api/admin/module/{{.EntitySnake}}/service"
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
			controller.New{{.EntityPascal}}Controller,
			fx.As(new(rctrl.Controller)),
			fx.ResultTags(%sgroup:"controllers"%s),
		),
	),
)
`, special, special,
)
