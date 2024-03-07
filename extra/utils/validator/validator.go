package validator

import (
	"strings"

	"github.com/kimchhung/gva/extra/lang"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"

	ent "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"

	zht "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *validator.Validate // Use a single validator instance
)

func init() {
	validate = validator.New()

	if err := ent.RegisterDefaultTranslations(validate, lang.GetTranslator(lang.LocaleEN)); err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}
	if err := zht.RegisterDefaultTranslations(validate, lang.GetTranslator(lang.LocaleZH)); err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}
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
