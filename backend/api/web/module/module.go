package module

import (
	"github.com/kimchhung/gva/backend-echo/app/constant"
	"github.com/kimchhung/gva/backend-echo/internal/echoc"
	index "github.com/kimchhung/gva/backend/api/web/module/index"
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
			fx.ParamTags(constant.TagWebController),

			// register echoc.ModuleRouter to container as member of module group
			fx.ResultTags(constant.TagModule),
		),
	),
)
