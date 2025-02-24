package module

import (
	"backend/api/bot/module/index"
	"backend/app/common/controller"
	"backend/internal/ctr"

	"go.uber.org/fx"
)

var NewbotModules = fx.Module("bot-module",
	index.IndexModule,

	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => ctr.ModuleRouter
			fx.As(new(ctr.ModuleRouter)),

			// take group params from container => []ctr.CTR -> NewRouter
			fx.ParamTags(controller.BotAdminController),

			// register to container as member of module group
			fx.ResultTags(controller.TagModule),
		),
	),
)
