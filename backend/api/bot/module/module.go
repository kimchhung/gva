package module

import (
	"github.com/gva/api/bot/module/comic"
	"github.com/gva/api/bot/module/index"
	"github.com/gva/app/constant"
	"github.com/gva/internal/echoc"

	"go.uber.org/fx"
)

var NewbotModules = fx.Module("bot-module",
	index.IndexModule,
	comic.ComicModuleModule,

	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => echoc.ModuleRouter
			fx.As(new(echoc.ModuleRouter)),

			// take group params from container => []echoc.Controller -> NewRouter
			fx.ParamTags(constant.BotAdminController),

			// register to container as member of module group
			fx.ResultTags(constant.TagModule),
		),
	),
)
