package web

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/kimchhung/gva/extra/app/module/web/demo"
	"github.com/kimchhung/gva/extra/config"
	_ "github.com/kimchhung/gva/extra/docs"
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

func (r *Router) Name() string {
	return "web"
}

// @title GVA Web API
// @version 2.0
// @description GO VUE ADMIN Boilerplate
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func (r *Router) Register(app fiber.Router, cfg *config.Config) {
	base := "/api"
	api := app.Group(base)

	api.Use(swagger.New(swagger.Config{
		Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
		BasePath: base,
		FilePath: "./docs/web_swagger.json",
		Path:     "swagger",
		Title:    "Web API Docs",
		CacheAge: 0,
	}))

	for _, controller := range r.controllers {
		controller.Routes(api)
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
			fx.ParamTags(`group:"web-controller"`),

			// register rctrl.ModuleRouter to container as member of module group
			fx.ResultTags(`group:"module"`),
		),
	),
)
