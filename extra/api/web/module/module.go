package module

import (
	index "github.com/kimchhung/gva/extra/api/web/module/index"
	"github.com/kimchhung/gva/extra/app/constant"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"go.uber.org/fx"
)

var NewWebModules = fx.Module("web-module",
	index.IndexModule,
	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => rctrl.ModuleRouter
			fx.As(new(rctrl.ModuleRouter)),

			// take group params from container => []rctrl.Controller -> NewRouter
			fx.ParamTags(constant.TagWebController),

			// register rctrl.ModuleRouter to container as member of module group
			fx.ResultTags(constant.TagModule),
		),
	),
)
