package module

import (
	"github.com/gva/api/web/graph"
	mdatetime "github.com/gva/api/web/module/datetime"
	"github.com/gva/app/common/controller"
	"github.com/gva/internal/ctr"

	"go.uber.org/fx"
)

var NewWebModules = fx.Module("web-module",
	mdatetime.DatetimeModule,
	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => ctr.ModuleRouter
			fx.As(new(ctr.ModuleRouter)),

			// take group params from container => []ctr.CTR -> NewRouter
			fx.ParamTags(controller.TagWebController),

			// register ctr.ModuleRouter to container as member of module group
			fx.ResultTags(controller.TagModule),
		),
	),

	// lifecycle on start hook depend on order
	graph.Module,
)
