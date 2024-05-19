package request

import (
	"github.com/gofiber/fiber/v2"

	appctx "github.com/kimchhung/gva/extra/app/common/context"
	"github.com/kimchhung/gva/extra/internal/ent"
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

// required middleware jwtService.RequiredAdmin()
func MustAdminContext(out *appctx.AdminContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		*out = *appctx.MustAdminContext(c.UserContext())
		return nil
	}
}

// required middleware jwtService.RequiredAdmin()
func MustAdmin(out *ent.Admin) fiber.Handler {
	return func(c *fiber.Ctx) error {
		*out = *appctx.MustAdminContext(c.UserContext()).Admin
		return nil
	}
}
