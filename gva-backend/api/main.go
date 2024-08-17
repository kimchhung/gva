package main

import (
	admindocs "github.com/gva/api/admin/docs"
	admin "github.com/gva/api/admin/module"
	botdocs "github.com/gva/api/bot/docs"
	bot "github.com/gva/api/bot/module"
	webdocs "github.com/gva/api/web/docs"
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

	admindocs.SwaggerInfoadmin.Host = cfg.Middleware.Swagger.Host
	admindocs.SwaggerInfoadmin.BasePath = cfg.API.Admin.BasePath

	botdocs.SwaggerInfobot.Host = cfg.Middleware.Swagger.Host
	botdocs.SwaggerInfobot.BasePath = cfg.API.Bot.BasePath

	webdocs.SwaggerInfoweb.Host = cfg.Middleware.Swagger.Host
	webdocs.SwaggerInfoweb.BasePath = cfg.API.Web.BasePath

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
