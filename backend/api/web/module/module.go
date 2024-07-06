package module

import (
	"github.com/gva/api/web/graph"
	index "github.com/gva/api/web/module/index"
	"github.com/gva/app/common/controller"
	"github.com/gva/internal/echoc"

	"go.uber.org/fx"
)

var NewWebModules = fx.Module("web-module",
	index.IndexModule,
	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => echoc.ModuleRouter
			fx.As(new(echoc.ModuleRouter)),

			// take group params from container => []echoc.Controller -> NewRouter
			fx.ParamTags(controller.TagWebController),

			// register echoc.ModuleRouter to container as member of module group
			fx.ResultTags(controller.TagModule),
		),
	),

	// lifecycle on start hook depend on order
	graph.Module,
)
