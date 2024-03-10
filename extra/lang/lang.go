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

type LocaleType string

const (
	LocaleEN LocaleType = "en"
	LocaleZH LocaleType = "zh"
)

func init() {
	en := en.New()
	UT = ut.New(en, en, zh.New())

	if err := UT.Import(ut.FormatJSON, "./lang"); err != nil {
		log.Panic().Err(err).Msg("UT.Import(ut.FormatJSON")
	}

	if err := UT.VerifyTranslations(); err != nil {
		log.Panic().Err(err).Msg("VerifyTranslations")
	}
}

func getTranslator(locale LocaleType) ut.Translator {
	trans, found := UT.GetTranslator(string(locale))
	if !found {
		log.Error().Msgf("translator not found for locale %s, using default", "en")
	}

	return trans
}

type (
	langKey         struct{}
	TranslateOption func(*Config)
	LocaleOption    func(*Config)
	Config          struct {
		locale       LocaleType
		fallbackFunc func(key string) string
		params       []string
	}
)

// default locale
func Register(headerName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if headerName == "" {
			headerName = "locale"
		}

		locale := c.Get(headerName)
		if locale == "" {
			return c.Next()
		}

		ctx := context.WithValue(c.UserContext(), langKey{}, c.Get(headerName))
		c.SetUserContext(ctx)
		return c.Next()
	}
}

// get translators
func GetTranslator(localeOpt LocaleOption) ut.Translator {
	config := &Config{locale: LocaleEN}
	localeOpt(config)
	return getTranslator(config.locale)
}

// translate key with params
func T(localeOpt LocaleOption, key string, opts ...TranslateOption) string {
	config := &Config{
		locale: LocaleEN,
	}

	localeOpt(config)
	for _, op := range opts {
		op(config)
	}

	trans, err := getTranslator(config.locale).T(key, config.params...)
	if err != nil {
		log.Warn().Msgf("translator not found for locale %s, key %s, err %v", config.locale, key, err)

		if config.fallbackFunc != nil {
			return config.fallbackFunc(key)
		}

		return strings.ReplaceAll(key, "_", " ")
	}

	return trans
}

func Is(source LocaleOption, target LocaleType) bool {
	config := &Config{}
	source(config)
	return config.locale == target
}
