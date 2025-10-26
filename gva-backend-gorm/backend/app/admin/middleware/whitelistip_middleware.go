package middleware

import (
	admincontext "backend/app/admin/context"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) IpGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			adminctx, err := admincontext.GetAdminContext(c.Request().Context())
			if err != nil {
				return next(c)
			}

			if adminctx.Admin.IpWhiteList != nil {
				currentIp := m.ip_s.GetCurrentIP(c)
				if err := m.ip_s.VerifyWhiteListIP(currentIp, adminctx.Admin.IpWhiteList); err != nil {
					return err
				}
			}

			return next(c)
		}
	}
}
