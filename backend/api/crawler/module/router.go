package module

import (
	"github.com/labstack/echo/v4"

	docs "github.com/gva/api/crawler/docs"
	"github.com/gva/env"
	"github.com/gva/internal/echoc"
	"github.com/gva/utils"
	"github.com/gva/utils/swagger"

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
	cfg := args[0].(*env.Config)

	//default value if not exist in env config
	utils.SetIfEmpty(&cfg.API.Crawler.BasePath, "/crawler/v1")
	docs.SwaggerInfocrawler.BasePath = cfg.API.Crawler.BasePath

	swagger.Register(app, cfg.API.Crawler.BasePath, cfg.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfocrawler.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := app.Group(cfg.API.Crawler.BasePath)
	for _, controller := range r.controllers {
		echoc.Register(api, controller)
	}
}
