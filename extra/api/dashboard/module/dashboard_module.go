package dashboard

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/kimchhung/gva/extra/api/dashboard/module/admin"
	"github.com/kimchhung/gva/extra/api/dashboard/module/auth"
	"github.com/kimchhung/gva/extra/api/dashboard/module/permission"
	"github.com/kimchhung/gva/extra/api/dashboard/module/role"
	"github.com/kimchhung/gva/extra/api/dashboard/module/route"
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
	return &Router{
		controllers,
	}
}

func (r *Router) Register(app fiber.Router, cfg *config.Config) {
	basePath := "/dashboard"
	if cfg.API.Dashboard.BasePath != "" {
		basePath = cfg.API.Dashboard.BasePath
	}

	api := app.Group(basePath)
	api.Use(swagger.New(swagger.Config{
		Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
		BasePath: basePath,
		FilePath: "./docs/dashboard/dashboard_swagger.json",
		Path:     "swagger",
		Title:    "Dashboard API Docs",
		CacheAge: 0,
	}))

	for _, controller := range r.controllers {
		controller.Routes(api)
	}
}

var NewDashboardModules = fx.Module("dashboard-module",
	admin.NewAdminModule,
	auth.NewAuthModule,
	role.NewRoleModule,
	permission.NewPermissionModule,
	route.NewRouteModule,

	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => rctrl.ModuleRouter
			fx.As(new(rctrl.ModuleRouter)),

			// take group params from container => []rctrl.Controller -> NewRouter
			fx.ParamTags(`group:"dashboard-controller"`),

			// register to container as member of module group
			fx.ResultTags(`group:"module"`),
		),
	),
)
