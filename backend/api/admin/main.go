package main

import (
	"github.com/kimchhung/gva/backend/api/admin/docs"
	admin "github.com/kimchhung/gva/backend/api/admin/module"
	"github.com/kimchhung/gva/backend/env"

	"github.com/kimchhung/gva/backend/app"
	"github.com/kimchhung/gva/backend/app/router"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)

	_ "github.com/kimchhung/gva/backend/internal/ent/runtime"
	_ "github.com/swaggo/swag"
)

var (
	cfg = env.NewConfig()
)

// @title						GVA admin API
// @version						1.0
// @description					GO VUE ADMIN Boilerplate
// @host						localhost:8080
// @BasePath					/admin/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description					Type "Bearer" followed by a space and JWT token.
func Run() {
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

	/* Admin |> module <| */
	modules := router.WithRouter(
		admin.APIAdminModules,
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
