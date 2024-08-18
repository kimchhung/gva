package main

import (
	"github.com/gva/api/admin/docs"
	admin "github.com/gva/api/admin/module"
	"github.com/gva/env"

	"github.com/gva/app"
	"github.com/gva/app/router"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)

	_ "github.com/gva/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg = env.NewConfig()
	version = "0.001"
)

// @title						GVA admin API
// @version						1.0
// @description					GO VUE ADMIN Boilerplate
// @host						localhost:4000
// @BasePath					/admin/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description					Type "Bearer" followed by a space and JWT token.
func Run() {
	// * Run only web api
	// force app to enable only admin module
	cfg.API.Web.Enable = false
	cfg.API.Admin.Enable = true
	cfg.API.Bot.Enable = false

	if cfg.API.Admin.Address != "" {
		cfg.App.Address = cfg.API.Admin.Address
	}

	docs.SwaggerInfoadmin.Host = cfg.Middleware.Swagger.Host
	docs.SwaggerInfoadmin.BasePath = cfg.API.Admin.BasePath
	docs.SwaggerInfoadmin.Version = version

	/* Admin |> module <| */
	modules := router.WithRouter(
		admin.NewAdminModules,
	)

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
