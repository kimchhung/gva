package main

import (
	admin "backend/api/admin/module"
	"backend/env"

	"backend/app"
	"backend/app/router"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)

	_ "github.com/swaggo/swag"
)

var (
	cfg = env.NewConfig()
)

// @title						GVA admin API
// @version						1.0
// @description					GO VUE ADMIN Boilerplate
// @host						localhost:4000
// @BasePath					/admin/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description			Type "Bearer" followed by a space and JWT token.
func Run() {
	// * Run only web api
	// force app to enable only admin module
	cfg.API.Admin.Enable = true
	cfg.API.Bot.Enable = false

	if cfg.API.Admin.Port != "" {
		cfg.App.Port = cfg.API.Admin.Port
	}

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
