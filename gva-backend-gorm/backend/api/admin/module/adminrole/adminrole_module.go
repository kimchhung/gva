package adminrole

import (
	"backend/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var AdminRoleModule = fx.Module("AdminRoleModule",
	// Register Service
	fx.Provide(NewAdminRoleService),

	// Regiser Controller
	controller.ProvideAdminController(NewAdminRoleController),
)
