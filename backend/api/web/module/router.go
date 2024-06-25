package module

import (
	"github.com/gva/api/web/docs"
	"github.com/gva/env"
	"github.com/gva/utils/swagger"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/gva/internal/echoc"
	"github.com/gva/utils"
)

var _ interface{ echoc.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []echoc.Controller
}

func NewRouter(controllers ...echoc.Controller) *Router {
	return &Router{controllers}
}

func (r *Router) Register(app *echo.Echo, args ...any) {
	cfg := args[0].(*env.Config)

	//default value if not exist in env config
	utils.SetIfEmpty(&cfg.API.Web.BasePath, "/web/v1")
	docs.SwaggerInfoweb.BasePath = cfg.API.Web.BasePath

	swagger.Register(app, cfg.API.Web.BasePath, cfg.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfoweb.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := app.Group(cfg.API.Web.BasePath)
	for _, controller := range r.controllers {
		echoc.Register(api, controller)
	}
}
