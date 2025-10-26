package middleware

import (
	"backend/core/utils/request"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) ResponseHook() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (returnErr error) {
			request.SetHook(c.Request().Context())
			return next(c)
		}
	}
}
