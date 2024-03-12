package main

import (
	"github.com/kimchhung/gva/extra/api/admin/docs"
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/app/module"
	"github.com/kimchhung/gva/extra/config"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)

	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg = config.NewConfig()
)

// @title GVA admin API
// @version 1.0
// @description GO VUE ADMIN Boilerplate
// @host localhost:8080
// @BasePath /admin
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Param locale header string true "en"
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// * Run only web api
	docs.SwaggerInfoadmin.BasePath = "admin"
	if cfg.API.Admin.BasePath != "" {
		docs.SwaggerInfoadmin.BasePath = cfg.API.Admin.BasePath
	}

	// force app to enable only admin module
	cfg.API.Admin.Enable = true
	cfg.API.Web.Enable = false

	if cfg.API.Admin.Port != "" {
		cfg.App.Port = cfg.API.Admin.Port
	}

	app.New(
		/* global config */
		cfg,

		/* Web, admin |> Module <| */
		module.NewModules(cfg),
	).Run()
}
