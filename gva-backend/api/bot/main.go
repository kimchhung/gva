package main

import (
	"github.com/gva/api/bot/docs"
	bot "github.com/gva/api/bot/module"
	"github.com/gva/app"
	"github.com/gva/app/router"
	"github.com/gva/env"

	_ "github.com/gva/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg     = env.NewConfig()
	version = "0.001"
)

// @title						GVA bot API
// @version					    1.0
// @description					GO VUE ADMIN Boilerplate
// @host						localhost:5500
// @BasePath					/bot/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @Param  						accept-language  header     string     false  "some description"
func Run() {
	// * Run only web api

	// force app to enable only web module
	cfg.API.Web.Enable = false
	cfg.API.Admin.Enable = false
	cfg.API.Bot.Enable = true

	// overwrite app port
	if cfg.API.Bot.Address != "" {
		cfg.App.Address = cfg.API.Bot.Address
	}

	docs.SwaggerInfobot.Host = cfg.Middleware.Swagger.Host
	docs.SwaggerInfobot.BasePath = cfg.API.Bot.BasePath
	docs.SwaggerInfobot.Version = version

	/* Web |> module <| */
	modules := router.WithRouter(bot.NewbotModules)

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
