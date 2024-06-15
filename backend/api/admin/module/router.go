package module

import (
	"github.com/labstack/echo/v4"

	"github.com/kimchhung/gva/backend-echo/config"
	"github.com/kimchhung/gva/backend-echo/internal/echoc"
	"github.com/kimchhung/gva/backend-echo/utils"
	"github.com/kimchhung/gva/backend-echo/utils/swagger"
	docs "github.com/kimchhung/gva/backend/api/admin/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var _ interface{ echoc.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []echoc.Controller
}

func NewRouter(controllers []echoc.Controller) *Router {
	return &Router{controllers}
}

func (r *Router) Register(app *echo.Echo, args ...any) {
	cfg := args[0].(*config.Config)

	//default value if not exist in env config
	utils.SetIfEmpty(&cfg.API.Admin.BasePath, "/admin/v1")
	docs.SwaggerInfoadmin.BasePath = cfg.API.Admin.BasePath

	swagger.Register(app, cfg.API.Admin.BasePath, cfg.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfoadmin.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := app.Group(cfg.API.Admin.BasePath)
	for _, controller := range r.controllers {
		echoc.Register(api, controller)
	}
}
