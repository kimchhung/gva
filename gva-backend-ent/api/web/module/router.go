package module

import (
	"github.com/gva/api/web/docs"
	"github.com/gva/app/common/controller"
	"github.com/gva/env"
	"github.com/gva/internal/utils/swagger"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/gva/internal/ctr"

	"github.com/gva/internal/utils"
)

var _ interface{ ctr.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []ctr.CTR
}

func NewRouter(controllers ...ctr.CTR) *Router {
	return &Router{controllers}
}

func (r *Router) Register(args ...any) {
	app := args[0].(*echo.Echo)
	cfg := args[1].(*env.Config)

	//default value if not exist in env config
	utils.SetIfEmpty(&cfg.API.Web.BasePath, "/web/v1")
	docs.SwaggerInfoweb.BasePath = cfg.API.Web.BasePath

	swagger.Register(app, cfg.API.Web.BasePath, cfg.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfoweb.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := app.Group(cfg.API.Web.BasePath)
	if err := controller.RegisterEcho(api, r.controllers); err != nil {
		panic(err)
	}
}
