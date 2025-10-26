package botrouter

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	docs "backend/app/bot/docs"
	"backend/core/env"
	"backend/core/router"
	coretype "backend/core/type"
	"backend/core/utils/swagger"
	"backend/internal/ctr"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	_ interface{ coretype.AppRouter } = (*Router)(nil)
)

const (
	Controller router.Group = "bot-controller"
)

type RouterParam struct {
	fx.In
	App         *echo.Echo
	Env         *env.Config
	Controllers []ctr.CTR `group:"bot-controller"`
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
	docs.SwaggerInfobot.BasePath = r.env.Bot.BasePath
	docs.SwaggerInfobot.Host = r.env.App.Host

	swagger.Register(r.app, r.env.Bot.BasePath, r.env.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfobot.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := r.app.Group(r.env.Bot.BasePath)
	if err := router.RegisterEcho(api, r.controllers); err != nil {
		panic(err)
	}
}
