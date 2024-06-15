package request

import (
	"github.com/labstack/echo/v4"

	appctx "github.com/kimchhung/gva/backend-echo/app/common/context"
	"github.com/kimchhung/gva/backend-echo/internal/ent"
	"github.com/kimchhung/gva/backend-echo/utils/validator"
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

// required middleware jwtService.RequiredAdmin()
func MustAdminContext(out *appctx.AdminContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		*out = *appctx.MustAdminContext(c.Request().Context())
		return nil
	}
}

// required middleware jwtService.RequiredAdmin()
func MustAdmin(out *ent.Admin) echo.HandlerFunc {
	return func(c echo.Context) error {
		*out = *appctx.MustAdminContext(c.Request().Context()).Admin
		return nil
	}
}
