package module

import (
	"backend/api/bot/module/index"
	"backend/app/common/controller"

	"go.uber.org/fx"
)

var NewbotModules = fx.Module("bot-module",
	index.IndexModule,
	controller.Bot.AddRouter(NewRouter),
)
