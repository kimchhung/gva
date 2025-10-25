package coremiddleware

import (
	corecontext "backend/core/context"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) RequestContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := corecontext.WithRequestContext(req.Context())
			c.SetRequest(req.WithContext(ctx))

			return next(c)
		}
	}
}
