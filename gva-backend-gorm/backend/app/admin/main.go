package main

import (
	admin "backend/app/admin/module"
	"backend/app/share/app"
	"backend/core/env"

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
	cfg.Admin.Enable = true
	cfg.Bot.Enable = false

	if cfg.Admin.Port != "" {
		cfg.App.Port = cfg.Admin.Port
	}

	/* Admin |> module <| */

	app.New(
		/* global config */
		cfg,
		admin.NewAdminModules,
	).Run()
}

func main() {
	Run()
}
