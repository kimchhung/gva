package admin

import (
	"github.com/kimchhung/gva/app/module/admin/controller"
	"github.com/kimchhung/gva/app/module/admin/repository"
	"github.com/kimchhung/gva/app/module/admin/service"
	"github.com/kimchhung/gva/internal/rctrl"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	rctrl.Router
} = &AdminRouter{}

type AdminRouter struct {
	app        fiber.Router
	controller controller.IAdminController
}

// Router methods
func NewAdminRouter(fiber *fiber.App, controller controller.IAdminController) *AdminRouter {
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
	fx.Provide(fx.Annotate(
		controller.NewAdminController,
		fx.As(new(controller.IAdminController)),
	)),

	// Register Router
	fx.Provide(fx.Annotate(
		NewAdminRouter,
		fx.As(new(rctrl.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
