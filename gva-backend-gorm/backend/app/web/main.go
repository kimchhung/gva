package main

import (
	"backend/app/share/app"
	"backend/core/env"

	web "backend/app/web/module"

	_ "github.com/swaggo/swag"
)

var (
	cfg = env.NewConfig()
)

// @title					    GVA bot API
// @version				     	1.0
// @description				    GO VUE ADMIN Boilerplate
// @host						localhost:4000
// @BasePath					/bot/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @Param  		 accept-language  header     string     false  "some description"
func Run() {
	// * Run only web api

	// force app to enable only web module
	cfg.Admin.Enable = false
	cfg.Bot.Enable = false
	cfg.Web.Enable = true

	// overwrite app port
	if cfg.Bot.Port != "" {
		cfg.App.Port = cfg.Bot.Port
	}

	/* Web |> module <| */
	app.New(
		/* global config */
		cfg,
		web.NewWebModules,
	).Run()
}

func main() {
	Run()
}
