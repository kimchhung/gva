package request

import (
	"github.com/kimchhung/gva/extra/internal/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Nothing to describe this fucking variable.
var IsProduction bool

func init() {
	zerolog.ErrorStackMarshaler = MarshalStackSkip(3)
}

// A fuction to return beautiful and structured responses.
func Response(c *fiber.Ctx, opt response.ReponseOption, opts ...response.ReponseOption) error {
	return response.New(opt, opts...).Parse(c)
}
