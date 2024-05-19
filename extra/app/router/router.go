package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/app/constant"
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

func (r *Router) Register(app fiber.Router, args ...any) {
	for _, r := range r.modules {
		r.Register(app, args...)
	}
}

func WithRouter(modules ...fx.Option) []fx.Option {
	app := append(modules, fx.Provide(
		// register as *Router
		fx.Annotate(NewRouter,
			// take group params from container => []rctrl.ModuleRouter -> NewRouter
			fx.ParamTags(constant.TagModule),
		),
	))

	return app
}
