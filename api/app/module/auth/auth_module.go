package auth

import (
	"github.com/kimchhung/gva/app/module/auth/controller"
	"github.com/kimchhung/gva/app/module/auth/service"
	"github.com/kimchhung/gva/internal/rctrl"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	rctrl.Router
} = &AuthRouter{}

type AuthRouter struct {
	app        fiber.Router
	controller controller.IAuthController
}

// Router methods
func NewAuthRouter(fiber *fiber.App, controller controller.IAuthController) *AuthRouter {
	return &AuthRouter{
		app:        fiber,
		controller: controller,
	}
}

func (r *AuthRouter) Register() {
	r.controller.Routes(r.app)
}

// Register bulkly
var NewAuthModule = fx.Module("AuthModule",
	fx.Provide(service.NewAuthService),

	// Regiser Controller
	fx.Provide(fx.Annotate(
		controller.NewAuthController,
		fx.As(new(controller.IAuthController))),
	),

	// Register Router
	fx.Provide(fx.Annotate(
		NewAuthRouter,
		fx.As(new(rctrl.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
