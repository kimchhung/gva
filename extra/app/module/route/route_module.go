package route

import (
	"github.com/kimchhung/gva/extra/app/module/route/controller"
	"github.com/kimchhung/gva/extra/app/module/route/repository"
	"github.com/kimchhung/gva/extra/app/module/route/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewRouteModule = fx.Module("RouteModule",
	// Register Repository & Service
	fx.Provide(repository.NewRouteRepository),
	fx.Provide(service.NewRouteService),

	// Regiser Controller
	fx.Provide(fx.Annotate(controller.NewRouteController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"controllers"`),
	)),
)
