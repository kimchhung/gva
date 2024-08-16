package main

import (
	admin "github.com/gva/api/admin/module"
	bot "github.com/gva/api/bot/module"
	web "github.com/gva/api/web/module"
	"github.com/gva/app"
	"github.com/gva/app/router"
	"github.com/gva/env"
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

	/* Enable Web Module */
	if cfg.API.Web.Enable {
		modules = append(modules, web.NewWebModules)
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
