package main

import (
	adminmodule "backend/app/admin/module"
	botmodule "backend/app/bot/module"
	"backend/app/share/app"
	"backend/core/env"

	webmodule "backend/app/web/module"

	"go.uber.org/fx"
)

var (
	cfg = env.NewConfig()
)

// Run both web and admin api
func Run() {

	/* Web, admin |> Module <| */
	modules := []fx.Option{}

	/* Enable admin Module */
	if cfg.Admin.Enable {
		modules = append(modules, adminmodule.NewAdminModules)
	}

	/* Enable Bot Module */
	if cfg.Bot.Enable {
		modules = append(modules, botmodule.NewbotModules)
	}

	if cfg.Web.Enable {
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
