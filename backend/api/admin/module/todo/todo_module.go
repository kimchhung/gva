package todo

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var TodoModule = fx.Module("TodoModule",
	// Register Service
	fx.Provide(NewTodoService),

	// Regiser Controller
	controller.ProvideAdminController(NewTodoController),
)
