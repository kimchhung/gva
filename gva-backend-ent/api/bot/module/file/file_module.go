package file

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var FileModuleModule = fx.Module("FileModule",
	fx.Provide(NewFileService),
	// Regiser Controller
	controller.ProvideBotController(NewFileController),
)
