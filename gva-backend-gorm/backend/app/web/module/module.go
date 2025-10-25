package module

import (
	"backend/app/web/module/index"
	webrouter "backend/app/web/router"
	"backend/core/router"

	"go.uber.org/fx"
)

var NewWebModules = fx.Module("web-module",
	index.IndexModule,

	// register router
	router.Add(webrouter.NewRouter),
)
