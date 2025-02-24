package lang

import (
	"context"
	"backend/internal/logger"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
)

var (
	uTranslator *ut.UniversalTranslator
)

type LocaleType string

const (
	LocaleEN LocaleType = "en"
	LocaleZH LocaleType = "zh"
)

func InitializeTranslator() error {
	en := en.New()
	zh := zh.New()
	uTranslator = ut.New(en, en, zh)

	if err := uTranslator.Import(ut.FormatJSON, "./lang"); err != nil {
		return err
	}

	if err := uTranslator.VerifyTranslations(); err != nil {
		return err
	}

	logger.G().Info("translator is initialized")
	return nil
}

func getTranslator(locale LocaleType) ut.Translator {
	trans, found := uTranslator.GetTranslator(string(locale))
	if !found {
		logger.G().Sugar().Error("translator not found for locale %s, using default", "en")
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
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			preferredLanguage := LocaleType(c.Request().Header.Get("Accept-Language"))
			switch preferredLanguage {
			case "zh", "zh-CN":
				preferredLanguage = LocaleZH
			default:
				preferredLanguage = LocaleEN
			}

			ctx := context.WithValue(c.Request().Context(), langKey{}, string(preferredLanguage))
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
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
		logger.G().Sugar().Warn("translator not found for locale %s, key %s, err %v", config.locale, key, err)

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
