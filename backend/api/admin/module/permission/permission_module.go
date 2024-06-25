package permission

import (
	"github.com/gva/app/constant"

	"go.uber.org/fx"
)

// Register bulkly
var PermissionModule = fx.Module("PermissionModule",
	// Register Service

	fx.Provide(NewPermissionService),

	// Regiser Controller
	constant.ProvideAdminController(NewPermissionController),
)
