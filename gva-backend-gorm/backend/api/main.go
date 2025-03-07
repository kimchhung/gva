package main

import (
	admin "backend/api/admin/module"
	bot "backend/api/bot/module"
	"backend/app"
	"backend/app/router"
	"backend/env"
)

var (
	cfg = env.NewConfig()
)

// Run both web and admin api
func Run() {

	/* Web, admin |> Module <| */
	modules := router.WithRouter()

	/* Enable admin Module */
	if cfg.API.Admin.Enable {
		modules = append(modules, admin.NewAdminModules)
	}

	/* Enable Bot Module */
	if cfg.API.Bot.Enable {
		modules = append(modules, bot.NewbotModules)
	}

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
