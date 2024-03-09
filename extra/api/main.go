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

func main() {
	app.New(
		/* global config */
		cfg,

		/* Web, admin |> Module <| */
		module.NewModules(cfg),
	).Run()
}
