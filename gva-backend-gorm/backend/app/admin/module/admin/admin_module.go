package admin

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var AdminModule = fx.Module("AdminModule",
	// Register Service
	fx.Provide(NewAdminService),

	// Regiser Controller
	adminrouter.Controller.Add(NewAdminController),
)
