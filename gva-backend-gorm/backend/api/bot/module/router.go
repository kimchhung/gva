package module

import (
	"github.com/labstack/echo/v4"

	docs "backend/api/bot/docs"
	"backend/app/common/controller"
	"backend/env"
	"backend/internal/ctr"
	"backend/utils"
	"backend/utils/swagger"

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
	utils.SetIfEmpty(&cfg.API.Bot.BasePath, "/bot/v1")
	docs.SwaggerInfobot.BasePath = cfg.API.Bot.BasePath

	swagger.Register(app, cfg.API.Bot.BasePath, cfg.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfobot.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := app.Group(cfg.API.Bot.BasePath)
	if err := controller.RegisterEcho(api, r.controllers); err != nil {
		panic(err)
	}
}
