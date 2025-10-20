package operationlog

import (
	"backend/app/common/controller"

	"go.uber.org/fx"
)

// Register bulkly
var OperationLogModule = fx.Module("OperationLogModule",
	// Register Service
	fx.Provide(NewOperationLogService),

	// Regiser Controller
	controller.Admin.AddController(NewOperationLogController),
)
