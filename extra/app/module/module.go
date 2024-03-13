package module

import (
	"github.com/gofiber/fiber/v2"

	admin "github.com/kimchhung/gva/extra/api/admin/module"
	web "github.com/kimchhung/gva/extra/api/web/module"
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

func New(cfg *config.Config) fx.Option {
	modules := []fx.Option{
		/* Register module router to fiber base on config */
		fx.Provide(
			// register as *Router
			fx.Annotate(NewRouter,
				// take group params from container => []rctrl.ModuleRouter -> NewRouter
				fx.ParamTags(`group:"modules"`),
			),
		),
	}

	/* Enable admin Module */
	if cfg.API.Admin.Enable {
		modules = append(modules, admin.NewadminModules)
	}

	/* Enable Web Module */
	if cfg.API.Web.Enable {
		modules = append(modules, web.NewWebModules)
	}

	return fx.Module("app", modules...)
}
