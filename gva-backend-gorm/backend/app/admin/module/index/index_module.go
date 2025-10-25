package index

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	adminrouter.Controller.Add(NewIndexController),
)
