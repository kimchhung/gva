package main

import (
	crawler "github.com/kimchhung/gva/backend/api/crawler/module"
	"github.com/kimchhung/gva/backend/api/web/docs"
	"github.com/kimchhung/gva/backend/app"
	"github.com/kimchhung/gva/backend/app/router"
	"github.com/kimchhung/gva/backend/env"

	_ "github.com/kimchhung/gva/backend/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg = env.NewConfig()
)

// @title					    GVA Crawler API
// @version				     	1.0
// @description				    GO VUE ADMIN Boilerplate
// @host						localhost:8080
// @BasePath					/crawler/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @Param   accept-language  header     string     false  "some description"
func Run() {
	// * Run only web api

	docs.SwaggerInfoweb.BasePath = "crawler"
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
	modules := router.WithRouter(crawler.NewCrawlerModules)

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
