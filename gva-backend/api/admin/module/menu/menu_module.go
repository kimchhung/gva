package menu

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var MenuModule = fx.Module("MenuModule",
	// Register Service
	fx.Provide(NewMenuService),

	// Regiser Controller
	controller.ProvideAdminController(NewMenuController),
)
