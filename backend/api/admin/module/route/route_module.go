package route

import (
	"github.com/kimchhung/gva/backend-echo/app/constant"
	"go.uber.org/fx"
)

// Register bulkly
var RouteModule = fx.Module("RouteModule",
	// Register Service
	fx.Provide(NewRouteService),

	// Regiser Controller
	constant.ProvideAdminController(NewRouteController),
)
