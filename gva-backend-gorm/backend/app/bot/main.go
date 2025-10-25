package main

import (
	bot "backend/app/bot/module"
	"backend/app/share/app"

	"backend/env"

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
	cfg.API.Admin.Enable = false
	cfg.API.Bot.Enable = true

	// overwrite app port
	if cfg.API.Bot.Port != "" {
		cfg.App.Port = cfg.API.Bot.Port
	}

	/* Web |> module <| */
	app.New(
		/* global config */
		cfg,
		bot.NewbotModules,
	).Run()
}

func main() {
	Run()
}
