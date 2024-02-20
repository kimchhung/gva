package module_template

import "fmt"

var special = "`"
var Module = fmt.Sprintf(
	`package {{.EntitySnake}}

import (
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/controller"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/repository"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/service"
	"github.com/kimchhung/gva/internal/rctrl"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	rctrl.Router
} = &{{.EntityPascal}}Router{}

type {{.EntityPascal}}Router struct {
	app        fiber.Router
	controller controller.I{{.EntityPascal}}Controller
}

// Router methods
func New{{.EntityPascal}}Router(fiber *fiber.App, controller controller.I{{.EntityPascal}}Controller) *{{.EntityPascal}}Router {
	return &{{.EntityPascal}}Router{
		app:        fiber,
		controller: controller,
	}
}

func (r *{{.EntityPascal}}Router) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var New{{.EntityPascal}}Module = fx.Module("{{.EntityPascal}}Module",
	// Register Repository & Service
	fx.Provide(repository.New{{.EntityPascal}}Repository),
	fx.Provide(service.New{{.EntityPascal}}Service),

	// Regiser Controller
	fx.Provide(fx.Annotate(
		controller.New{{.EntityPascal}}Controller,
		fx.As(new(controller.I{{.EntityPascal}}Controller))),
	),

	// Register Router
	fx.Provide(fx.Annotate(
		New{{.EntityPascal}}Router,
		fx.As(new(rctrl.Router)),
		fx.ResultTags(%sgroup:"routers"%s),
	)),
)
`, special, special,
)
