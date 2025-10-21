package module

import (
	"backend/api/web/docs"
	"backend/app/common/controller"
	"backend/env"
	"backend/internal/ctr"
	"backend/internal/ctxutil"
	"backend/utils/swagger"
	"context"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var _ interface{ ctr.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []ctr.CTR
}

func NewRouter(controllers []ctr.CTR) *Router {
	return &Router{controllers}
}

func (r *Router) Register(ctx context.Context) {
	app := ctxutil.MustValue[*echo.Echo](ctx)
	cfg := ctxutil.MustValue[*env.Config](ctx)

	//default value if not exist in env config
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
