package request

import (
	"github.com/labstack/echo/v4"
)

type Parser func(c echo.Context) (any, error)

var binder = &echo.DefaultBinder{}

/*
BodyParser(&Person{})

	type User struct {
	  ID string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	}

Swagger

	// @Param 		<name> <params-type> <value-type> <required> <description>
	// @Param 		info body dto.MenuRequest true "Route Info"
*/
func BodyParser(out any) Parser {
	return func(c echo.Context) (any, error) {
		return out, binder.BindBody(c, out)
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
	return func(c echo.Context) (any, error) {
		return out, binder.BindQueryParams(c, out)
	}
}

/*
ParamsParser(&param{})

	param := new(struct {ID uint `param:"id"`})

Swagger

	// @Param       <name> <params-type> <value-type> <required> <description>
	// @Param       id path int true "Route ID"
*/
func ParamsParser(out any) Parser {
	return func(c echo.Context) (any, error) {
		return out, binder.BindPathParams(c, out)
	}
}
