package main

import (
	admin "github.com/kimchhung/gva/extra/api/admin/module"
	web "github.com/kimchhung/gva/extra/api/web/module"
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/app/module"
	"github.com/kimchhung/gva/extra/config"
	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
)

var (
	cfg = config.NewConfig()
)

// Run both web and admin api
func main() {
	/* Web, admin |> Module <| */
	modules := module.New(cfg)

	/* Enable admin Module */
	if cfg.API.Admin.Enable {
		modules = append(modules, admin.NewadminModules)
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
