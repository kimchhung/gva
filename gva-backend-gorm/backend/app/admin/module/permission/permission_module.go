package permission

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var PermissionModule = fx.Module("PermissionModule",
	// Register Service
	fx.Provide(NewPermissionService),

	// Regiser Controller
	adminrouter.Controller.Add(NewPermissionController),
)
