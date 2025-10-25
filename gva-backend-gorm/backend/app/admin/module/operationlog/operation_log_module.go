package operationlog

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var OperationLogModule = fx.Module("OperationLogModule",
	// Register Service
	fx.Provide(NewOperationLogService),

	// Regiser Controller
	adminrouter.Controller.Add(NewOperationLogController),
)
