package permission

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var PermissionModule = fx.Module("PermissionModule",
	// Register Service
	fx.Provide(NewPermissionService),

	// Regiser Controller
	controller.ProvideAdminController(NewPermissionController),
)
