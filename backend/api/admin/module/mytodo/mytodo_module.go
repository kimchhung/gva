package mytodo

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var MyTodoModule = fx.Module("MyTodoModule",
	// Register Service
	fx.Provide(NewMyTodoService),

	// Regiser Controller
	controller.ProvideAdminController(NewMyTodoController),
)
