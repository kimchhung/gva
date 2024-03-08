package auth

import (
	"github.com/kimchhung/gva/extra/api/dashboard/module/auth/controller"
	"github.com/kimchhung/gva/extra/api/dashboard/module/auth/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewAuthModule = fx.Module("AuthModule",
	fx.Provide(service.NewAuthService),

	// Regiser Controller
	fx.Provide(fx.Annotate(controller.NewAuthController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"dashboard-controller"`),
	)),
)
