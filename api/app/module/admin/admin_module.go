package admin

import (
	"github.com/kimchhung/gva/app/module/admin/controller"
	"github.com/kimchhung/gva/app/module/admin/repository"
	"github.com/kimchhung/gva/app/module/admin/service"
	"github.com/kimchhung/gva/internal/control_route"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	control_route.Router
} = &AdminRouter{}

type AdminRouter struct {
	app        fiber.Router
	controller *controller.AdminController
}

// Router methods
func NewAdminRouter(fiber *fiber.App, controller *controller.AdminController) *AdminRouter {
	return &AdminRouter{
		app:        fiber,
		controller: controller,
	}
}

func (r *AdminRouter) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var NewAdminModule = fx.Module("AdminModule",
	// Register Repository & Service
	fx.Provide(repository.NewAdminRepository),
	fx.Provide(service.NewAdminService),

	// Regiser Controller
	fx.Provide(controller.NewAdminController),

	// Register Router
	fx.Provide(NewAdminRouter),
	fx.Provide(fx.Annotate(
		NewAdminRouter,
		fx.As(new(control_route.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
