package route

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var RouteModule = fx.Module("RouteModule",
	// Register Service
	fx.Provide(NewRouteService),

	// Regiser Controller
	controller.ProvideAdminController(NewRouteController),
)
