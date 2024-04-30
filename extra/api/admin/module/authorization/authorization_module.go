package authorization

import (
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/controller"
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/repository"
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewAuthorizationModule = fx.Module("AuthorizationModule",
	// Register Repository & Service
	fx.Provide(repository.NewRouteRepository),
	fx.Provide(repository.NewPermissionRepository),

	fx.Provide(service.NewPermissionService),
	fx.Provide(service.NewRouteService),

	// Regiser Controller
	fx.Provide(
		fx.Annotate(
			controller.NewAuthorizationController,
			fx.As(new(rctrl.Controller)),
			fx.ResultTags(`group:"admin-controllers"`),
		),
	),
)
