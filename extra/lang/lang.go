package lang

import (
	"context"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var (
	UT *ut.UniversalTranslator
)

const (
	LocaleEN = "en"
	LocaleZH = "zh"
)

func init() {
	UT = ut.New(en.New(), zh.New())
	err := UT.Import(ut.FormatJSON, "translation")
	if err != nil {
		log.Error().Err(err).Msg("UT.Import(ut.FormatJSON")
	}
}

func GetTranslator(locale string) ut.Translator {
	trans, found := UT.GetTranslator(locale)
	if !found {
		log.Warn().Msgf("translator not found for locale %s, using default", "en")
	}

	return trans
}

type (
	langKey struct{}
	Option  func(*Config)
	Config  struct {
		locale   string
		key      *string
		fallback *string
		params   []string
	}
)

func WithLocale(locale string) Option {
	return func(c *Config) {
		c.locale = locale
	}
}

func WithContext(ctx context.Context) Option {
	return func(c *Config) {
		if locale, ok := ctx.Value(langKey{}).(string); ok {
			c.locale = locale
		}
	}
}

func WithFiberCtx(ctx *fiber.Ctx) Option {
	return func(c *Config) {
		if locale, ok := ctx.UserContext().Value(langKey{}).(string); ok {
			c.locale = locale
		}
	}
}

// default locale
func Middleware(headerName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if headerName == "" {
			headerName = "locale"
		}

		ctx := context.WithValue(c.UserContext(), langKey{}, c.Get(headerName))
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func Get(opt Option, opts ...Option) ut.Translator {
	config := &Config{locale: LocaleEN}
	opt(config)

	for _, op := range opts {
		op(config)
	}

	return GetTranslator(config.locale)
}

type TOption func(*Config)

func Key(key string) TOption {
	return func(c *Config) {
		c.key = &key
	}
}

// default use key
func Fallback(text string) TOption {
	return func(c *Config) {
		c.fallback = &text
	}
}

// fallback
func FallbackWithoutSeperator(seperator string) Option {
	return func(c *Config) {
		if c.key != nil {
			text := strings.ReplaceAll(*c.key, seperator, " ")
			c.fallback = &text
		}
	}
}

func Params(params ...string) Option {
	return func(c *Config) {
		c.params = params
	}
}

// translate key with params
func T(opt TOption, opts ...Option) string {
	config := &Config{locale: LocaleEN}
	opt(config)

	for _, op := range opts {
		op(config)
	}

	if config.key == nil {
		return ""
	}

	if config.locale == "" {
		return *config.key
	}

	t := GetTranslator(config.locale)
	trans, err := t.T(config.key, config.params...)
	if err != nil {
		log.Warn().Msgf("translator not found for locale %s, key %s", config.locale, *config.key)

		if config.fallback != nil {
			return *config.fallback
		}

		return *config.key
	}

	return trans
}
