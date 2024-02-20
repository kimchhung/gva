package permission

import (
	"github.com/kimchhung/gva/app/module/permission/controller"
	"github.com/kimchhung/gva/app/module/permission/repository"
	"github.com/kimchhung/gva/app/module/permission/service"
	"github.com/kimchhung/gva/internal/rctrl"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	rctrl.Router
} = &PermissionRouter{}

type PermissionRouter struct {
	app        fiber.Router
	controller controller.IPermissionController
}

// Router methods
func NewPermissionRouter(fiber *fiber.App, controller controller.IPermissionController) *PermissionRouter {
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
	fx.Provide(fx.Annotate(
		controller.NewPermissionController,
		fx.As(new(controller.IPermissionController)),
	)),

	// Register Router
	fx.Provide(fx.Annotate(
		NewPermissionRouter,
		fx.As(new(rctrl.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
