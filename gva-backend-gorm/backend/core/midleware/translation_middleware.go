package coremiddleware

import (
	"backend/core/lang"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (m *Middleware) Translation() echo.MiddlewareFunc {
	if m.cfg.Middleware.Translation.Enable {
		if err := m.translator.Initialize(); err != nil {
			m.log.Error("failed to initialize translator", zap.Error(err))
		}

		m.translator.SetAsDefaultTranslator()
	}

	if m.translator.IsInitialized() {
		m.validator.RegisterValidatorTranslation()
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			preferredLanguage := lang.LocaleType(c.Request().Header.Get("Accept-Language"))
			switch preferredLanguage {
			case "km", "km-KH":
				preferredLanguage = lang.LocaleKM
			case "zh", "zh-CN":
				preferredLanguage = lang.LocaleZH
			default:
				preferredLanguage = lang.LocaleEN
			}

			ctx := lang.WithContext(c.Request().Context(), preferredLanguage)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
