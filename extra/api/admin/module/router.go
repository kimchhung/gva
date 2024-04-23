package admin

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/kimchhung/gva/extra/api/admin/module/admin"
	"github.com/kimchhung/gva/extra/api/admin/module/auth"
	"github.com/kimchhung/gva/extra/api/admin/module/route"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/utils"
	"go.uber.org/fx"
)

var _ interface{ rctrl.ModuleRouter } = (*Router)(nil)

type Router struct {
	controllers []rctrl.Controller
}

func NewRouter(controllers []rctrl.Controller) *Router {
	return &Router{controllers}
}

func (r *Router) Register(app fiber.Router, cfg *config.Config) {
	basePath := "/admin"
	if cfg.API.Admin.BasePath != "" {
		basePath = cfg.API.Admin.BasePath
	}

	api := app.Group(basePath)
	api.Use(swagger.New(swagger.Config{
		Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
		BasePath: basePath,
		FilePath: "./api/admin/docs/admin_swagger.json",
		Path:     "swagger",
		Title:    "admin API Docs",
		CacheAge: 0,
	}))

	for _, controller := range r.controllers {
		rctrl.Register(api, controller)
	}
}

var NewadminModules = fx.Module("admin-module",
	admin.NewAdminModule,
	auth.NewAuthModule,
	route.NewRouteModule,

	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => rctrl.ModuleRouter
			fx.As(new(rctrl.ModuleRouter)),

			// take group params from container => []rctrl.Controller -> NewRouter
			fx.ParamTags(`group:"admin-controller"`),

			// register to container as member of module group
			fx.ResultTags(`group:"modules"`),
		),
	),
)
