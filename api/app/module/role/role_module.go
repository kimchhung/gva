package role

import (
	"gva/app/module/role/controller"
	"gva/app/module/role/repository"
	"gva/app/module/role/service"
	"gva/internal/control_route"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	control_route.Router
} = &RoleRouter{}

type RoleRouter struct {
	app        fiber.Router
	controller *controller.RoleController
}

// Router methods
func NewRoleRouter(fiber *fiber.App, controller *controller.RoleController) *RoleRouter {
	return &RoleRouter{
		app:        fiber,
		controller: controller,
	}
}

func (r *RoleRouter) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var NewRoleModule = fx.Module("RoleModule",
	// Register Repository & Service
	fx.Provide(repository.NewRoleRepository),
	fx.Provide(service.NewRoleService),

	// Regiser Controller
	fx.Provide(controller.NewRoleController),

	// Register Router
	fx.Provide(NewRoleRouter),
	fx.Provide(fx.Annotate(
		NewRoleRouter,
		fx.As(new(control_route.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
