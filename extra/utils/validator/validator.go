package validator

import (
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	ent "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	Trans    ut.Translator
)

func init() {
	validate = validator.New()

	uni = ut.New(en.New())
	Trans, _ = uni.GetTranslator("en")

	if err := ent.RegisterDefaultTranslations(validate, Trans); err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}
}

func ValidateStruct(input any) error {
	return validate.Struct(input)
}

// Remove unnecessary fields from validator message
func RemoveTopStruct(fields map[string]string) string {
	res := []string{}
	for _, msg := range fields {
		res = append(res, msg)
	}
	return strings.Join(res, ", ")
}
