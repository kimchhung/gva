package todo

import (
	"github.com/kimchhung/gva/extra/app/module/todo/controller"
	"github.com/kimchhung/gva/extra/app/module/todo/repository"
	"github.com/kimchhung/gva/extra/app/module/todo/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewTodoModule = fx.Module("TodoModule",
	// Register Repository & Service
	fx.Provide(repository.NewTodoRepository),
	fx.Provide(service.NewTodoService),

	// Regiser Controller
	fx.Annotate(
		controller.NewTodoController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"controllers"`),
	),
)
