package controller

import (
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Controller string

const (
	Bot   Controller = `group:"bot-controllers"`
	Admin Controller = `group:"admin-controllers"`
	Web   Controller = `group:"web-controllers"`

	TagRouters = `group:"routers"`
)

func (tag Controller) AddController(c any) fx.Option {
	return fx.Provide(
		fx.Annotate(c,
			fx.As(new(ctr.CTR)),
			fx.ResultTags(string(tag)),
		),
	)
}

func (tag Controller) AddRouter(r any) fx.Option {
	return fx.Provide(
		fx.Annotate(r,
			// convert type *Router => ctr.ModuleRouter
			fx.As(new(ctr.ModuleRouter)),

			// take group params from container => []ctr.CTR -> NewRouter
			fx.ParamTags(string(tag)),

			// register to container as member of module group
			fx.ResultTags(TagRouters),
		),
	)
}

// register to container  type echoc.Controller and tag as admin module
func ProvideController(contructor any, tag Controller) fx.Option {
	return fx.Provide(
		fx.Annotate(
			contructor,
			fx.As(new(ctr.CTR)),
			fx.ResultTags(string(tag)),
		),
	)
}

func RegisterEcho(api *echo.Group, icontrollers []ctr.CTR) error {
	controllers, err := ctr.Reflect(icontrollers...)
	if err != nil {
		return err
	}

	for _, c := range controllers {
		group := api
		if !c.IsEmptyPrefix() {
			group = api.Group(c.GetPrefix(), c.Middlewares...)
		}

		for _, r := range c.Routes {
			if r.Callback != nil {
				r.Callback(group)
				continue
			}

			handler := func(c echo.Context) (err error) {
				for _, fn := range r.ScopeHandler() {
					if err = fn(c); err != nil {
						return err
					}
				}
				return
			}

			route := group.Add(r.GetMethod(), r.GetPath(), handler, r.Middlewares...)
			route.Name = r.GetName()
		}
	}

	return nil
}
