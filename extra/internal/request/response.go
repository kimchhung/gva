package request

import (
	"time"

	"github.com/kimchhung/gva/extra/internal/response"

	"github.com/gofiber/fiber/v2"
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
func Response(c *fiber.Ctx, opts ...response.ReponseOption) error {
	return response.New(opts...).Parse(c)
}

func ReponseNow(c *fiber.Ctx) error {
	return Response(c, response.Data(time.Now().UTC().Format(time.RFC3339)))
}
