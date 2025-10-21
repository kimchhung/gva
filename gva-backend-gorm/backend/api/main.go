package main

import (
	adminmodule "backend/api/admin/module"
	botmodule "backend/api/bot/module"
	webmodule "backend/api/web/module"
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
		modules = append(modules, adminmodule.NewAdminModules)
	}

	/* Enable Bot Module */
	if cfg.API.Bot.Enable {
		modules = append(modules, botmodule.NewbotModules)
	}

	if cfg.API.Web.Enable {
		modules = append(modules, webmodule.NewWebModules)
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
