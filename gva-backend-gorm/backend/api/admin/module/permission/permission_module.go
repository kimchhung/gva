package permission

import (
	"backend/app/common/controller"

	"go.uber.org/fx"
)

// Register bulkly
var PermissionModule = fx.Module("PermissionModule",
	// Register Service
	fx.Provide(NewPermissionService),

	// Regiser Controller
	controller.Admin.AddController(NewPermissionController),
)
