package module

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/kimchhung/gva/extra/api/admin/docs"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/utils"
)

var _ interface{ rctrl.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []rctrl.Controller
}

func NewRouter(controllers []rctrl.Controller) *Router {
	return &Router{controllers}
}

func (r *Router) Register(app fiber.Router, args ...any) {
	cfg := args[0].(*config.Config)

	//default value if not exist in env config
	utils.SetIfEmpty(&cfg.API.Admin.BasePath, "/admin")
	docs.SwaggerInfoadmin.BasePath = cfg.API.Admin.BasePath

	api := app.Group(cfg.API.Admin.BasePath,
		swagger.New(swagger.Config{
			Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
			BasePath: cfg.API.Admin.BasePath,
			FilePath: "./api/admin/docs/admin_swagger.json",
			Path:     "swagger",
			Title:    "admin API Docs",
			CacheAge: -1,
		}),
	)

	for _, controller := range r.controllers {
		rctrl.Register(api, controller)
	}
}
