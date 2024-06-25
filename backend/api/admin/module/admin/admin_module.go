package admin

import (
	"github.com/gva/app/constant"

	"go.uber.org/fx"
)

// Register bulkly
var AdminModule = fx.Module("AdminModule",

	fx.Provide(NewAdminService),

	// Regiser Controller
	constant.ProvideAdminController(NewAdminController),
)
