package main

import (
	admin "github.com/kimchhung/gva/backend/api/admin/module"
	web "github.com/kimchhung/gva/backend/api/web/module"
	"github.com/kimchhung/gva/backend/app"
	"github.com/kimchhung/gva/backend/app/router"
	"github.com/kimchhung/gva/backend/env"

	_ "github.com/kimchhung/gva/backend/internal/ent/runtime"
)

var (
	cfg = env.NewConfig()
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
