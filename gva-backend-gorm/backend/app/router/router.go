package router

import (
	"backend/app/common/controller"
	"backend/internal/ctr"
	"context"

	"go.uber.org/fx"
)

var _ interface{ ctr.ModuleRouter } = (*Router)(nil)

type Router struct {
	modules []ctr.ModuleRouter
}

func NewRouter(modules []ctr.ModuleRouter) *Router {
	r := &Router{
		modules: modules,
	}

	return r
}

func (r *Router) Register(ctx context.Context) {
	for _, r := range r.modules {
		r.Register(ctx)
	}
}

func WithRouter(modules ...fx.Option) []fx.Option {
	app := append(modules, fx.Provide(
		// register as *Router
		fx.Annotate(NewRouter,
			// take group params from container => []ctr.ModuleRouter -> NewRouter
			fx.ParamTags(controller.TagRouters),
		),
	))

	return app
}
