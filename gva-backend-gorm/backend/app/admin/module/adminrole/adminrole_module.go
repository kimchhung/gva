package adminrole

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var AdminRoleModule = fx.Module("AdminRoleModule",
	// Register Service
	fx.Provide(NewAdminRoleService),

	// Regiser Controller
	adminrouter.Controller.Add(NewAdminRoleController),
)
