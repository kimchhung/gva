package main

import (
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/config"
	web "github.com/kimchhung/gva/extra/docs/web"

	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
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
func main() {

	web.SwaggerInfoweb.BasePath = "web"
	if cfg.API.Web.BasePath != "" {
		web.SwaggerInfoweb.BasePath = cfg.API.Web.BasePath
	}

	// force app to enable web module
	cfg.API.Web.Enable = true
	cfg.API.Dashboard.Enable = false

	// overwrite app port
	if cfg.API.Web.Port != "" {
		cfg.App.Port = cfg.API.Web.Port
	}

	app.New(cfg).Run()
}
