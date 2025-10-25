package index

import (
	botrouter "backend/app/bot/router"

	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	botrouter.Controller.Add(NewIndexController),
)
