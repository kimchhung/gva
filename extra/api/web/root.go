package web

import (
	"github.com/kimchhung/gva/extra/api/web/docs"
	web "github.com/kimchhung/gva/extra/api/web/module"
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/app/router"

	"github.com/kimchhung/gva/extra/config"

	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg = config.NewConfig()
)

// @title GVA Web API
// @version 1.0
// @description GO VUE ADMIN Boilerplate
// @host localhost:8080

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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
