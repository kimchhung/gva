package webrouter

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	docs "backend/app/web/docs"
	"backend/core/router"
	coretype "backend/core/type"
	"backend/core/utils"
	"backend/core/utils/swagger"
	"backend/env"
	"backend/internal/ctr"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	_ interface{ coretype.AppRouter } = (*Router)(nil)
)

const (
	Controller router.Group = "web-controller"
)

type RouterParam struct {
	fx.In
	App         *echo.Echo
	Env         *env.Config
	Controllers []ctr.CTR `group:"web-controller"`
}

type Router struct {
	app         *echo.Echo
	env         *env.Config
	controllers []ctr.CTR
}

func NewRouter(p RouterParam) *Router {
	return &Router{
		controllers: p.Controllers,
		app:         p.App,
		env:         p.Env,
	}
}

func (r *Router) Register(ctx context.Context) {
	//default value if not exist in env config
	utils.SetIfEmpty(&r.env.API.Web.BasePath, "/web/v1")
	docs.SwaggerInfoweb.BasePath = r.env.API.Web.BasePath
	docs.SwaggerInfoweb.Host = r.env.App.Host

	swagger.Register(r.app, r.env.API.Web.BasePath, r.env.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfoweb.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := r.app.Group(r.env.API.Web.BasePath)
	if err := router.RegisterEcho(api, r.controllers); err != nil {
		panic(err)
	}
}
