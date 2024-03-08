package permission

import (
	"github.com/kimchhung/gva/extra/api/dashboard/module/permission/controller"
	"github.com/kimchhung/gva/extra/api/dashboard/module/permission/repository"
	"github.com/kimchhung/gva/extra/api/dashboard/module/permission/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewPermissionModule = fx.Module("PermissionModule",
	// Register Repository & Service
	fx.Provide(repository.NewPermissionRepository),
	fx.Provide(service.NewPermissionService),

	// Regiser Controller
	fx.Provide(fx.Annotate(
		controller.NewPermissionController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"dashboard-controller"`),
	)),
)
