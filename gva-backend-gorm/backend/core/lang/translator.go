package lang

import (
	"context"
	"strings"

	"backend/env"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/km"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
)

var (
	_UTranslator *ut.UniversalTranslator
	_Translator  = &Translator{}
)

func IsInitialized() bool {
	return _UTranslator != nil
}

type LocaleType string

const (
	LocaleEN LocaleType = "en"
	LocaleKM LocaleType = "km"
	LocaleZH LocaleType = "zh"
)

type Translator struct {
	cfg *env.Config
	log *zap.Logger
}

func NewTranslator(cfg *env.Config, log *zap.Logger) *Translator {
	t := &Translator{}
	t.cfg = cfg
	t.log = log.Named("translator")

	return t
}

func (t *Translator) IsInitialized() bool {
	return IsInitialized()
}

func (t *Translator) UTranslator() *ut.UniversalTranslator {
	return _UTranslator
}

func (t *Translator) Initialize() error {
	en := en.New()
	km := km.New()
	zh := zh.New()
	_UTranslator = ut.New(en, en, km, zh)

	if err := t.UTranslator().Import(ut.FormatJSON, "./lang"); err != nil {
		return err
	}

	if err := t.UTranslator().VerifyTranslations(); err != nil {
		t.log.Panic("failed to initialize translator")
	}

	t.log.Info("translator is initialized")
	return nil
}

func (t *Translator) SetAsDefaultTranslator() error {
	_Translator = t
	return nil
}

func (t *Translator) getTranslator(locale LocaleType) ut.Translator {
	trans, found := t.UTranslator().GetTranslator(string(locale))
	if !found {
		t.log.Sugar().Warnf("translator not found for locale %s, using default", "en")
	}

	return trans
}

// get translators
func (t *Translator) GetTranslator(localeOpt LocaleOption) ut.Translator {
	config := &Config{locale: LocaleEN}
	localeOpt(config)
	return t.getTranslator(config.locale)
}

func (t *Translator) T(localeOpt LocaleOption, key string, opts ...TranslateOption) string {
	if !IsInitialized() {
		return key
	}

	config := &Config{
		locale: LocaleEN,
	}

	localeOpt(config)
	for _, op := range opts {
		op(config)
	}

	trans, err := t.getTranslator(config.locale).T(key, config.params...)
	if err != nil {

		if config.fallbackFunc != nil {
			return config.fallbackFunc(key)
		}

		return strings.ReplaceAll(key, "_", " ")
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

// translate key with params

func Is(source LocaleOption, target LocaleType) bool {
	config := &Config{}
	source(config)
	return config.locale == target
}

func WithContext(ctx context.Context, locale LocaleType) context.Context {
	return context.WithValue(ctx, langKey{}, locale)
}

func FromContext(ctx context.Context) LocaleType {
	if locale, ok := ctx.Value(langKey{}).(LocaleType); ok {
		return locale
	}
	panic("locale not found in context")
}

func DefaultTranslator() *Translator {
	return _Translator
}
