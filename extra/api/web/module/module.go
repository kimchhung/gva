package module

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/swag/example/basic/docs"

	"github.com/kimchhung/gva/extra/api/web/module/demo"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/utils"
	"go.uber.org/fx"
)

var _ interface{ rctrl.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []rctrl.Controller
}

func NewRouter(controllers ...rctrl.Controller) *Router {
	return &Router{
		controllers,
	}
}

func (r *Router) Register(app fiber.Router, cfg *config.Config) {
	basePath := "/web"
	if cfg.API.Web.BasePath != "" {
		basePath = cfg.API.Web.BasePath
		docs.SwaggerInfo.BasePath = basePath
	}

	api := app.Group(basePath)
	api.Use(swagger.New(swagger.Config{
		Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
		BasePath: basePath,
		FilePath: "./api/web/docs/web_swagger.json",
		Path:     "swagger",
		Title:    "Web API Docs",
		CacheAge: 0,
	}))

	for _, controller := range r.controllers {
		rctrl.Register(app, controller)
	}
}

var NewWebModules = fx.Module("web-module",
	demo.NewDemoModule,
	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => rctrl.ModuleRouter
			fx.As(new(rctrl.ModuleRouter)),

			// take group params from container => []rctrl.Controller -> NewRouter
			fx.ParamTags(`group:"web-controllers"`),

			// register rctrl.ModuleRouter to container as member of module group
			fx.ResultTags(`group:"modules"`),
		),
	),
)
