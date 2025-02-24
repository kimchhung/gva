package request

import (
	"github.com/gva/internal/response"

	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
)

// Nothing to describe this fucking variable.
var IsProduction bool

func init() {
	zerolog.ErrorStackMarshaler = MarshalStackSkip(3)
}

/*
A fuction to return beautiful and structured responses.

	Response(c,response.Data(...))

	{
		code:0,
		message:"OK"
		data:any
	}

	Response(c)

	{
		code:0,
		message:"OK"
		data:any
	}

	Response(c,response.Error(...))

	{
		code:-5,
		message:".....error"
	}
*/
func Response(c echo.Context, opts ...response.ReponseOption) error {
	return response.New(opts...).Parse(c)
}
