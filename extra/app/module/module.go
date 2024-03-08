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

func (r *Router) Register(app fiber.Router, cfg *config.Config) {
	for _, r := range r.modules {
		r.Register(app, cfg)
	}
}

func NewModules(cfg *config.Config) fx.Option {
	modules := []fx.Option{
		/* Register module router to fiber base on config */
		fx.Provide(
			// register as *Router
			fx.Annotate(NewRouter,
				// take group params from container => []rctrl.ModuleRouter -> NewRouter
				fx.ParamTags(`group:"module"`),
			),
		),
	}

	/* Enable Dashboard Module */
	if slices.Contains(cfg.Module.Enables, "dashboard") {
		modules = append(modules, dashboard.NewDashboardModules)
	}

	/* Enable Web Module */
	if slices.Contains(cfg.Module.Enables, "web") {
		modules = append(modules, web.NewWebModules)
	}

	return fx.Module("app", modules...)
}
