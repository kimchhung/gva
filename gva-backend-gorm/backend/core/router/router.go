package router

import (
	coretype "backend/core/type"
	"context"

	"go.uber.org/fx"
)

var _ interface{ coretype.AppRouter } = (*Router)(nil)

func Add(r any) fx.Option {
	return fx.Provide(
		fx.Annotate(r,
			// convert type *Router => ctr.ModuleRouter
			fx.As(new(coretype.AppRouter)),

			// register to container as member of module group
			fx.ResultTags(`group:"app-routers"`),
		),
	)
}

type RouterParam struct {
	fx.In
	AppRouters []coretype.AppRouter `group:"app-routers"`
}

type Router struct {
	appRouters []coretype.AppRouter
}

func NewRouter(p RouterParam) *Router {
	r := &Router{
		appRouters: p.AppRouters,
	}

	return r
}

func (r *Router) Register(ctx context.Context) {
	for _, appRouter := range r.appRouters {
		appRouter.Register(ctx)
	}
}
