package router

import (
	"github.com/gva/app/constant"
	"github.com/gva/internal/echoc"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var _ interface{ echoc.ModuleRouter } = (*Router)(nil)

type Router struct {
	modules []echoc.ModuleRouter
}

func NewRouter(modules []echoc.ModuleRouter) *Router {
	r := &Router{
		modules: modules,
	}

	return r
}

func (r *Router) Register(app *echo.Echo, args ...any) {
	for _, r := range r.modules {
		r.Register(app, args...)
	}
}

func WithRouter(modules ...fx.Option) []fx.Option {
	app := append(modules, fx.Provide(
		// register as *Router
		fx.Annotate(NewRouter,
			// take group params from container => []echoc.ModuleRouter -> NewRouter
			fx.ParamTags(constant.TagModule),
		),
	))

	return app
}
