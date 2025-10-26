package request

import (
	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"

	"backend/core/validator"
)

func SetDefaultValue(v any) echo.HandlerFunc {
	return func(c echo.Context) error {
		return defaults.Set(v)
	}
}

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
