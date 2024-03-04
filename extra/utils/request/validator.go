package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/utils/validator"
)

// use as middleware Validate(BodyParser(&body),ParamsParser(out))
func Validate(parser Parser, parsers ...Parser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, parse := range append(parsers, parser) {

			data, err := parse(c)
			if err != nil {
				return err
			}

			if err = validator.ValidateStruct(data); err != nil {
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
