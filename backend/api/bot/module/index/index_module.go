package index

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	controller.ProvideBotController(NewIndexController),
)
