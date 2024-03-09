package permission

import (
	"github.com/kimchhung/gva/extra/api/admin/module/permission/controller"
	"github.com/kimchhung/gva/extra/api/admin/module/permission/repository"
	"github.com/kimchhung/gva/extra/api/admin/module/permission/service"
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
		fx.ResultTags(`group:"admin-controller"`),
	)),
)
