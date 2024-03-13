package main

import (
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

	app.New(
		/* global config */
		cfg,
		modules,
	).Run()
}
