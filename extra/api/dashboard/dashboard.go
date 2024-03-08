package main

import (
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/docs/dashboard"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)

	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
)

var (
	cfg = config.NewConfig()
)

// @title GVA Dashboard API
// @version 1.0
// @description GO VUE ADMIN Boilerplate
// @host localhost:8080
// @BasePath /dashboard
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	dashboard.SwaggerInfodashboard.BasePath = "dashboard"
	if cfg.API.Dashboard.BasePath != "" {
		dashboard.SwaggerInfodashboard.BasePath = cfg.API.Dashboard.BasePath
	}

	// force app to enable dashboard module
	cfg.API.Dashboard.Enable = true
	cfg.API.Web.Enable = false

	if cfg.API.Dashboard.Port != "" {
		cfg.App.Port = cfg.API.Dashboard.Port
	}

	app.New(cfg).Run()
}
