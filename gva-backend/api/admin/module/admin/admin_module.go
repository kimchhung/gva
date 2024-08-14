package admin

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var AdminModule = fx.Module("AdminModule",

	fx.Provide(NewAdminService),

	// Regiser Controller
	controller.ProvideAdminController(NewAdminController),
)
