package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/utils/validator"
)

type Parser func(*fiber.Ctx) (any, error)

/*
BodyParser(&Person{})

	type Person struct {
	    Name string `json:"name" xml:"name" form:"name"`
	    Pass string `json:"pass" xml:"pass" form:"pass"`
	}
*/
func BodyParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.BodyParser(out)
	}
}

/*
QueryParser(&Person{})

	type Person struct {
	    Name     string     `query:"name"`
	    Pass     string     `query:"pass"`
	    Products []string   `query:"products"`
	}
*/
func QueryParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.QueryParser(out)
	}
}

/*
ParamsParser(&param{})

param := struct {ID uint `params:"id"`}{}
*/
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
