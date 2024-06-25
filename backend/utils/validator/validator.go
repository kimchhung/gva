package validator

import (
	"strings"

	"github.com/gva/internal/lang"

	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"

	ent "github.com/go-playground/validator/v10/translations/en"

	zht "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *validator.Validate // Use a single validator instance
)

func InitializeValidator() error {
	validate = validator.New()
	enTranslator := lang.GetTranslator(lang.Locale(lang.LocaleEN))
	zhTranslator := lang.GetTranslator(lang.Locale(lang.LocaleZH))

	if err := ent.RegisterDefaultTranslations(validate, enTranslator); err != nil {
		return err
	}

	if err := zht.RegisterDefaultTranslations(validate, zhTranslator); err != nil {
		return err
	}

	log.Info().Msg("Validator is initialized")
	return nil
}

func ValidateStruct(input any) error {
	return validate.Struct(input)
}

func RemoveTopStruct(fields map[string]string) string {
	res := []string{}
	for _, msg := range fields {
		res = append(res, msg)
	}
	return strings.Join(res, ", ")
}
