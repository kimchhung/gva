package module

import (
	"github.com/gva/api/bot/module/comic"
	"github.com/gva/api/bot/module/index"
	"github.com/gva/app/common/controller"
	"github.com/gva/internal/ctr"

	"go.uber.org/fx"
)

var NewbotModules = fx.Module("bot-module",
	index.IndexModule,
	comic.ComicModuleModule,

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
