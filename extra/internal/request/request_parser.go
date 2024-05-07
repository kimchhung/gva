package request

import (
	"github.com/gofiber/fiber/v2"
)

type Parser func(*fiber.Ctx) (any, error)

/*
BodyParser(&Person{})

	type Person struct {
	    Name string `json:"name" xml:"name" form:"name"`
	    Pass string `json:"pass" xml:"pass" form:"pass"`
	}

Swagger

	// @Param 		<name> <params-type> <value-type> <required> <description>
	// @Param 		info body dto.RouteRequest true "Route Info"
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

Swagger

	// @Param       <name> <params-type> <value-type> <required> <description>
	// @Param       person query dto.Person true "person info"
*/
func QueryParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.QueryParser(out)
	}
}

/*
ParamsParser(&param{})

	param := struct {ID uint `params:"id"`}{}

Swagger

	// @Param       <name> <params-type> <value-type> <required> <description>
	// @Param       id path int true "Route ID"
*/
func ParamsParser(out any) Parser {
	return func(c *fiber.Ctx) (any, error) {
		return out, c.ParamsParser(out)
	}
}
