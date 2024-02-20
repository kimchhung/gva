package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/utils/validator"
)

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
