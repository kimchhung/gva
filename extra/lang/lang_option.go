package lang

import (
	"context"
)

func Locale(locale LocaleType) LocaleOption {
	return func(c *Config) {
		c.locale = locale
	}
}

// get locale from context
func ForContext(ctx context.Context) LocaleOption {
	return func(c *Config) {
		if locale, ok := ctx.Value(langKey{}).(string); ok {
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
