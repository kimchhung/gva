package module

import (
	"backend/api/web/module/index"

	"backend/app/common/controller"

	"go.uber.org/fx"
)

var NewWebModules = fx.Module("web-module",
	index.IndexModule,

	controller.Web.AddRouter(NewRouter),
)
