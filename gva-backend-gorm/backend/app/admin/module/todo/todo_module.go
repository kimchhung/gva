package todo

import (
	adminrouter "backend/app/admin/router"
	"go.uber.org/fx"
)

// Register bulkly
var TodoModule = fx.Module("TodoModule",
	// Register Service
	fx.Provide(NewTodoService),

	// Regiser Controller
	adminrouter.Controller.Add(NewTodoController),
)
