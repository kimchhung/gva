package validator

import (
	"strings"

	"backend/env"
	"backend/internal/bootstrap/lang"
	validateEnTranslation "backend/internal/bootstrap/validator/translations/en"
	validateKmTranslation "backend/internal/bootstrap/validator/translations/km"
	validateZhTranslation "backend/internal/bootstrap/validator/translations/zh"

	"github.com/go-playground/validator/v10"
)

var _validate *validator.Validate

func IsInitialied() bool {
	return _validate != nil
}

type Validator struct {
	cfg        *env.Config
	translator *lang.Translator
}

func NewValidator(cfg *env.Config, translator *lang.Translator) *Validator {
	v := &Validator{
		cfg:        cfg,
		translator: translator,
	}

	v.Initialize()
	return v
}

func (v *Validator) Initialize() {
	_validate = validator.New()
}

func (v *Validator) RegisterValidatorTranslation() {
	var (
		enTranslator = v.translator.GetTranslator(lang.Locale(lang.LocaleEN))
		khTranslator = v.translator.GetTranslator(lang.Locale(lang.LocaleKM))
		zhTranslator = v.translator.GetTranslator(lang.Locale(lang.LocaleZH))
	)

	validateEnTranslation.RegisterDefaultTranslations(v.Validate(), enTranslator)
	validateKmTranslation.RegisterDefaultTranslations(v.Validate(), khTranslator)
	validateZhTranslation.RegisterDefaultTranslations(v.Validate(), zhTranslator)
}

func (v *Validator) ValidateStruct(input any) error {
	return v.Validate().Struct(input)
}

func (v *Validator) Validate() *validator.Validate {
	return _validate
}

func RemoveTopStruct(fields map[string]string) string {
	res := []string{}
	for _, msg := range fields {
		res = append(res, msg)
	}
	return strings.Join(res, ", ")
}

func ValidateStruct(input any) error {
	return _validate.Struct(input)
}
