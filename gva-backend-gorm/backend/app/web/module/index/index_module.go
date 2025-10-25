package index

import (
	webrouter "backend/app/web/router"

	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	webrouter.Controller.Add(NewIndexController),
)
