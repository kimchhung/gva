package module

import (
	index "github.com/kimchhung/gva/extra/api/web/module/index"
	"github.com/kimchhung/gva/extra/app/constant"
	"github.com/kimchhung/gva/extra/internal/echoc"
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
