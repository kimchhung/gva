package todo

import (
	"github.com/kimchhung/gva/app/module/todo/controller"
	"github.com/kimchhung/gva/app/module/todo/repository"
	"github.com/kimchhung/gva/app/module/todo/service"
	"github.com/kimchhung/gva/internal/rctrl"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	rctrl.Router
} = &TodoRouter{}

type TodoRouter struct {
	app        fiber.Router
	controller controller.ITodoController
}

// Router methods
func NewTodoRouter(fiber *fiber.App, controller controller.ITodoController) *TodoRouter {
	return &TodoRouter{
		app:        fiber,
		controller: controller,
	}
}

func (r *TodoRouter) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var NewTodoModule = fx.Module("TodoModule",
	// Register Repository & Service
	fx.Provide(repository.NewTodoRepository),
	fx.Provide(service.NewTodoService),

	// Regiser Controller
	fx.Provide(fx.Annotate(
		controller.NewTodoController,
		fx.As(new(controller.ITodoController))),
	),

	// Register Router
	fx.Provide(fx.Annotate(
		NewTodoRouter,
		fx.As(new(rctrl.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
