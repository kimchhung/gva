package permission

import (
	"gva/app/module/permission/controller"
	"gva/app/module/permission/repository"
	"gva/app/module/permission/service"
	"gva/internal/control_route"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	control_route.Router
} = &PermissionRouter{}

type PermissionRouter struct {
	app        fiber.Router
	controller *controller.PermissionController
}

// Router methods
func NewPermissionRouter(fiber *fiber.App, controller *controller.PermissionController) *PermissionRouter {
	return &PermissionRouter{
		app:        fiber,
		controller: controller,
	}
}

func (r *PermissionRouter) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var NewPermissionModule = fx.Module("PermissionModule",
	// Register Repository & Service
	fx.Provide(repository.NewPermissionRepository),
	fx.Provide(service.NewPermissionService),

	// Regiser Controller
	fx.Provide(controller.NewPermissionController),

	// Register Router
	fx.Provide(NewPermissionRouter),
	fx.Provide(fx.Annotate(
		NewPermissionRouter,
		fx.As(new(control_route.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
