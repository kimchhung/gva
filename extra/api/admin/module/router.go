package admin

import (
	"strings"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"

	"github.com/kimchhung/gva/extra/api/admin/module/admin"
	"github.com/kimchhung/gva/extra/api/admin/module/auth"
	"github.com/kimchhung/gva/extra/api/admin/module/authorization"

	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
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

func (r *Router) Register(app fiber.Router, cfg *config.Config, args ...any) {
	utils.SetIfEmpty(cfg.API.Admin.BasePath, "/admin")
	api := app.Group(cfg.API.Admin.BasePath)

	api.Get("/now", func(c *fiber.Ctx) error {
		return request.Response(c, response.Data(time.Now().UTC().Format(time.RFC3339)))
	})

	api.Use(
		r.useSwagger(cfg),
		r.useJwtGuard(cfg, args[0].(*database.Database)),
	)

	for _, controller := range r.controllers {
		rctrl.Register(api, controller)
	}
}

func (r *Router) useSwagger(cfg *config.Config) fiber.Handler {
	return swagger.New(swagger.Config{
		Next:     utils.IsEnabled(cfg.Middleware.Swagger.Enable),
		BasePath: cfg.API.Admin.BasePath,
		FilePath: "./api/admin/docs/admin_swagger.json",
		Path:     "swagger",
		Title:    "admin API Docs",
		CacheAge: -1,
	})
}

func (r *Router) useJwtGuard(cfg *config.Config, db *database.Database) fiber.Handler {
	jwt_ := services.NewJwtService(cfg, db)

	return skip.New(jwt_.ProtectAdmin(), func(c *fiber.Ctx) bool {
		switch {
		case strings.Contains(string(c.Request().URI().Path()), "/auth/"):
			return true
		}

		return false
	})
}

var NewAdminModules = fx.Module("admin-module",
	admin.NewAdminModule,
	auth.NewAuthModule,
	authorization.NewAuthorizationModule,

	// manage admin, menu, role, route

	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => rctrl.ModuleRouter
			fx.As(new(rctrl.ModuleRouter)),

			// take group params from container => []rctrl.Controller -> NewRouter
			fx.ParamTags(`group:"admin-controllers"`),

			// register to container as member of module group
			fx.ResultTags(`group:"modules"`),
		),
	),
)
