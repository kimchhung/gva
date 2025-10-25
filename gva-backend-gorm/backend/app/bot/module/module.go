package module

import (
	"backend/app/bot/module/index"
	botrouter "backend/app/bot/router"
	"backend/core/router"

	"go.uber.org/fx"
)

var NewbotModules = fx.Module("bot-module",
	index.IndexModule,
	router.Add(botrouter.NewRouter),
)
