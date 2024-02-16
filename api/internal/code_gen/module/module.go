package module_template

import "fmt"

var special = "`"
var Module = fmt.Sprintf(
	`package {{.EntitySnake}}

import (
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/controller"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/repository"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/service"
	"github.com/kimchhung/gva/internal/control_route"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	control_route.Router
} = &{{.Entity}}Router{}

type {{.Entity}}Router struct {
	app        fiber.Router
	controller *controller.{{.Entity}}Controller
}

// Router methods
func New{{.Entity}}Router(fiber *fiber.App, controller *controller.{{.Entity}}Controller) *{{.Entity}}Router {
	return &{{.Entity}}Router{
		app:        fiber,
		controller: controller,
	}
}

func (r *{{.Entity}}Router) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var New{{.Entity}}Module = fx.Module("{{.Entity}}Module",
	// Register Repository & Service
	fx.Provide(repository.New{{.Entity}}Repository),
	fx.Provide(service.New{{.Entity}}Service),

	// Regiser Controller
	fx.Provide(controller.New{{.Entity}}Controller),

	// Register Router
	fx.Provide(New{{.Entity}}Router),
	fx.Provide(fx.Annotate(
		New{{.Entity}}Router,
		fx.As(new(control_route.Router)),
		fx.ResultTags(%sgroup:"routers"%s),
	)),
)
`, special, special,
)
