package module

import (
	"github.com/labstack/echo/v4"

	docs "github.com/gva/api/admin/docs"
	"github.com/gva/app/common/controller"
	"github.com/gva/env"
	"github.com/gva/internal/ctr"
	"github.com/gva/utils"
	"github.com/gva/utils/swagger"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var _ interface{ ctr.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []ctr.CTR
}

func NewRouter(controllers []ctr.CTR) *Router {
	return &Router{controllers}
}

func (r *Router) Register(args ...any) {
	app := args[0].(*echo.Echo)
	cfg := args[1].(*env.Config)

	//default value if not exist in env config
	utils.SetIfEmpty(&cfg.API.Admin.BasePath, "/admin/v1")
	docs.SwaggerInfoadmin.BasePath = cfg.API.Admin.BasePath

	swagger.Register(app, cfg.API.Admin.BasePath, cfg.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfoadmin.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := app.Group(cfg.API.Admin.BasePath)
	if err := controller.RegisterEcho(api, r.controllers); err != nil {
		panic(err)
	}
}
