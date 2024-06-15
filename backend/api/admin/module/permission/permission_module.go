package permission

import (
	"github.com/kimchhung/gva/backend/app/constant"

	"go.uber.org/fx"
)

// Register bulkly
var PermissionModule = fx.Module("PermissionModule",
	// Register Service

	fx.Provide(NewPermissionService),

	// Regiser Controller
	constant.ProvideAdminController(NewPermissionController),
)
