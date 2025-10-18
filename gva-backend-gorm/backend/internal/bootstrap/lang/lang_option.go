package lang

import (
	"context"

	"github.com/labstack/echo/v4"
)

func Locale(locale LocaleType) LocaleOption {
	return func(c *Config) {
		c.locale = locale
	}
}

// get prefered language from request context
func ForContext(ctx context.Context) LocaleOption {
	return func(c *Config) {
		locale := FromContext(ctx)
		c.locale = locale
	}
}

// get prefered language from request context
func ForEcho(ctx echo.Context) LocaleOption {
	return func(c *Config) {
		if locale, ok := ctx.Request().Context().Value(langKey{}).(string); ok {
			c.locale = LocaleType(locale)
		}
	}
}

type TOption func(*Config)

// default use key
func FallbackFunc(fn func(key string) string) TranslateOption {
	return func(c *Config) {
		c.fallbackFunc = fn
	}
}

func Params(params ...string) TranslateOption {
	return func(c *Config) {
		c.params = params
	}
}
