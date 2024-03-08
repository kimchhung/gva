package main

import (
	"github.com/kimchhung/gva/extra/app"
	"github.com/kimchhung/gva/extra/config"
	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
)

var (
	cfg = config.NewConfig()
)

func main() {
	app.New(cfg).Run()
}
