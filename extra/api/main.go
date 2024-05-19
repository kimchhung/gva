package main

import (
	admin "github.com/kimchhung/gva/extra/api/admin/module"
	web "github.com/kimchhung/gva/extra/api/web/module"
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/app/router"

	"github.com/kimchhung/gva/extra/config"
	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
)

var (
	cfg = config.NewConfig()
)

// Run both web and admin api
func Run() {
	/* Web, admin |> Module <| */
	modules := router.WithRouter()

	/* Enable admin Module */
	if cfg.API.Admin.Enable {
		modules = append(modules, admin.APIAdminModules)
	}

	/* Enable Web Module */
	if cfg.API.Web.Enable {
		modules = append(modules, web.NewWebModules)
	}

	app.New(
		/* global config */
		cfg,
		modules...,
	).Run()
}

func main() {
	Run()
}
