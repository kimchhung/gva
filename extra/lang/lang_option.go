package lang

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func Locale(locale string) LocaleOption {
	return func(c *Config) {
		c.locale = locale
	}
}

// get locale from context
func Ctx(ctx context.Context) LocaleOption {
	return func(c *Config) {
		if locale, ok := ctx.Value(langKey{}).(string); ok {
			c.locale = locale
		}
	}
}

// get locale from FiberCtx
func FiberCtx(ctx *fiber.Ctx) LocaleOption {
	return func(c *Config) {
		if locale, ok := ctx.UserContext().Value(langKey{}).(string); ok {
			c.locale = locale
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
