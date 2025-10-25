package router

import (
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
)

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
