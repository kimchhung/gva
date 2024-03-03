package role

import (
	"github.com/kimchhung/gva/extra/app/module/role/controller"
	"github.com/kimchhung/gva/extra/app/module/role/repository"
	"github.com/kimchhung/gva/extra/app/module/role/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	rctrl.Router
} = &RoleRouter{}

type RoleRouter struct {
	app        fiber.Router
	controller controller.IRoleController
}

// Router methods
func NewRoleRouter(fiber *fiber.App, controller controller.IRoleController) *RoleRouter {
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
	fx.Provide(fx.Annotate(
		controller.NewRoleController,
		fx.As(new(controller.IRoleController))),
	),

	// Register Router
	fx.Provide(fx.Annotate(
		NewRoleRouter,
		fx.As(new(rctrl.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
