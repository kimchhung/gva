package main

import (
	web "github.com/gva/api/web/module"
	"github.com/gva/app"
	"github.com/gva/app/router"
	"github.com/gva/env"

	_ "github.com/gva/internal/ent/runtime"

	_ "github.com/swaggo/swag"
)

var (
	cfg = env.NewConfig()
)

// @title					    GVA Web API
// @version				     	1.0
// @description				    GO VUE ADMIN Boilerplate
// @host						localhost:8080
// @BasePath					/web/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @Param   accept-language  header     string     false  "some description"
func Run() {
	// * Run only web api

	// force app to enable only web module
	cfg.API.Web.Enable = true
	cfg.API.Admin.Enable = false
	cfg.API.Bot.Enable = false

	// overwrite app port
	if cfg.API.Web.Port != "" {
		cfg.App.Port = cfg.API.Web.Port
	}

	/* Web |> module <| */
	modules := router.WithRouter(web.NewWebModules)

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
