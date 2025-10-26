package middleware

import (
	admincontext "backend/app/admin/context"
	apperror "backend/app/share/error"
	"backend/app/share/model"
	"backend/app/share/service"
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (m *Middleware) jwtAdminValidator(ctx context.Context, out *model.Admin) service.ClaimValidator {
	return func(claims jwt.MapClaims) error {
		id, ok := claims["id"].(float64)
		if !ok || id <= 0 {
			return apperror.ErrUnauthorized
		}

		admin, err := m.admin_r.GetById(ctx, uint(id))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.ErrNotFound
			}

			return err
		}

		if admin.Status == 0 {
			return apperror.ErrValidationError
		}

		err = m.admin_r.GetRolesByID(admin.ID, admin)
		if err != nil {
			return err
		}

		*out = *admin
		return nil
	}
}

func (m *Middleware) JwtGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token := strings.TrimSpace(strings.Replace(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ", "", 1))
			if token == "" {
				return apperror.ErrUnauthorized
			}

			ctx := c.Request().Context()
			admin := new(model.Admin)

			if _, err := m.jwt_s.ValidateToken(token, m.jwtAdminValidator(ctx, admin)); err != nil {
				m.log.Debug("s.ValidateToken", zap.Error(err))
				return apperror.ErrUnauthorized
			}

			adminctx := admincontext.NewAdminContext(admincontext.WithAdmin(admin))
			admincontext.SetAdminContext(ctx, adminctx)
			return next(c)
		}
	}
}
