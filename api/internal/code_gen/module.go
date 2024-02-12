package code_gen

var moduleTemplate = `package {{.EntitySnake}}

import (
	"gva/app/module/{{.EntitySnake}}/controller"
	"gva/app/module/{{.EntitySnake}}/repo"
	"gva/app/module/{{.EntitySnake}}/service"
	"gva/internal/control_route"

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
	fx.Provide(repo.New{{.Entity}}Repository),
	fx.Provide(service.New{{.Entity}}Service),

	// Regiser Controller
	fx.Provide(controller.New{{.Entity}}Controller),

	// Register Router
	fx.Provide(New{{.Entity}}Router),
	fx.Invoke(func(r *{{.Entity}}Router) {
		r.Register()
	}),
)

`
