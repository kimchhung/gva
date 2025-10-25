package request

import (
	"github.com/labstack/echo/v4"

	"backend/core/validator"
)

// use as middleware Validate(BodyParser(&body),ParamsParser(out))
func Validate(parser Parser, parsers ...Parser) echo.HandlerFunc {
	return func(c echo.Context) error {
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
func Parse(parser Parser, parsers ...Parser) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, parse := range append(parsers, parser) {
			if _, err := parse(c); err != nil {
				return err
			}
		}

		return nil
	}
}
