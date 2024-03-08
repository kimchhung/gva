package module

import (
	"slices"

	"github.com/gofiber/fiber/v2"

	"github.com/kimchhung/gva/extra/app/module/dashboard"
	"github.com/kimchhung/gva/extra/app/module/web"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"go.uber.org/fx"
)

var _ interface{ rctrl.ModuleRouter } = (*Router)(nil)

type Router struct {
	modules []rctrl.ModuleRouter
}

func NewRouter(modules []rctrl.ModuleRouter) *Router {
	r := &Router{
		modules: modules,
	}

	return r
}

func (*Router) Name() string {
	return "base"
}

func (r *Router) Register(app fiber.Router, cfg *config.Config) {
	for _, r := range r.modules {
		if slices.Contains(cfg.Module.Enables, r.Name()) {
			r.Register(app, cfg)
		}
	}
}

var NewModules = fx.Module("app",

	/* Web Module */
	web.NewWebModules,

	/* Dashboard Module */
	dashboard.NewDashboardModules,

	/* Register module router to fiber base on config */
	fx.Provide(
		fx.Annotate(NewRouter,
			fx.ParamTags(`group:"module"`),
		),
	),
)
