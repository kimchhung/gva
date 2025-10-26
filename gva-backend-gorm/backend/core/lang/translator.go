package lang

import (
	"backend/core/env"
	"context"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/km"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
)

var (
	_Translator = &Translator{}
)

func IsInitialized() bool {
	return _Translator != nil
}

type Translator struct {
	env *env.Config
	log *zap.Logger

	localePaths []string
	uTranslator *ut.UniversalTranslator
}

func NewTranslator(env *env.Config, log *zap.Logger) *Translator {
	ent := en.New()
	kmt := km.New()
	zht := zh.New()

	t := &Translator{
		env: env,
		log: log.Named("Translator"),
		localePaths: []string{
			"./app/share/locale",
		},
		uTranslator: ut.New(ent, ent, kmt, zht),
	}

	return t
}

func (t *Translator) IsInitialized() bool {
	return IsInitialized()
}

func (t *Translator) UTranslator() *ut.UniversalTranslator {
	return t.uTranslator
}

func (t *Translator) Import(paths ...string) {
	if !t.IsInitialized() {
		return
	}

	for _, path := range paths {
		t.log.Info("importing translations", zap.String("path", path))
		defer t.log.Info("import translations completed", zap.String("path", path))

		if err := t.UTranslator().Import(ut.FormatJSON, path); err != nil {
			t.log.Panic("failed to import translator", zap.String("path", path), zap.Error(err))
			continue
		}

		if err := t.UTranslator().VerifyTranslations(); err != nil {
			t.log.Panic("failed to import translator", zap.String("path", path), zap.Error(err))
		}
	}
}

func (t *Translator) Initialize() error {
	for _, path := range t.localePaths {
		t.Import(path)
	}

	t.log.Info("translator is initialized")
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
