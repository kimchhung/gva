package main

import (
	"github.com/kimchhung/gva/backend-echo/app"
	"github.com/kimchhung/gva/backend-echo/app/router"
	"github.com/kimchhung/gva/backend/api/web/docs"
	web "github.com/kimchhung/gva/backend/api/web/module"

	"github.com/kimchhung/gva/backend-echo/config"

	_ "github.com/kimchhung/gva/backend-echo/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg = config.NewConfig()
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

	docs.SwaggerInfoweb.BasePath = "web"
	if cfg.API.Web.BasePath != "" {
		docs.SwaggerInfoweb.BasePath = cfg.API.Web.BasePath
	}

	// force app to enable only web module
	cfg.API.Web.Enable = true
	cfg.API.Admin.Enable = false

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
