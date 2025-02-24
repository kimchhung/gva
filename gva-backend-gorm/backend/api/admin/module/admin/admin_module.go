package admin

import (
	"backend/app/common/controller"

	"go.uber.org/fx"
)

// Register bulkly
var AdminModule = fx.Module("AdminModule",
	// Register Service
	fx.Provide(NewAdminService),

	// Regiser Controller
	controller.ProvideAdminController(NewAdminController),
)
