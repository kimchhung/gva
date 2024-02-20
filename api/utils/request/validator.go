package request

import (
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
	trans    ut.Translator
)

func init() {
	validate = validator.New()

	uni = ut.New(en.New())
	trans, _ = uni.GetTranslator("en")

	if err := ent.RegisterDefaultTranslations(validate, trans); err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}
}

func ValidateStruct(input any) error {
	return validate.Struct(input)
}

type Parser func(*fiber.Ctx) (any, error)

func BodyParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.BodyParser(out)
	}
}

func QueryParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.QueryParser(out)
	}
}

func ParamsParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.ParamsParser(out)
	}
}

// use as middleware Validate(BodyParser(&body),ParamsParser(out))
func Validate(parser Parser, parsers ...Parser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, parse := range append(parsers, parser) {

			data, err := parse(c)
			if err != nil {
				return err
			}

			if err = ValidateStruct(data); err != nil {
				return err
			}
		}

		return nil
	}
}

// simple parser Parse(BodyParser(&body),ParamsParser(out))
func Parse(parser Parser, parsers ...Parser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, parse := range append(parsers, parser) {
			if _, err := parse(c); err != nil {
				return err
			}
		}

		return nil
	}
}
