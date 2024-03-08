package dashboard

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/kimchhung/gva/extra/app/module/dashboard/admin"
	"github.com/kimchhung/gva/extra/app/module/dashboard/auth"
	"github.com/kimchhung/gva/extra/app/module/dashboard/permission"
	"github.com/kimchhung/gva/extra/app/module/dashboard/role"
	"github.com/kimchhung/gva/extra/app/module/dashboard/route"
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
	base := "/dashboard"
	api := app.Group(base)

	api.Use(swagger.New(swagger.Config{
		Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
		BasePath: "/" + base,
		FilePath: "./docs/dashboard_swagger.json",
		Path:     "swagger",
		Title:    "Web API Docs",
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
