package request

import (
	"github.com/gofiber/fiber/v2"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/rql"
)

/*
@Required: meta.use(RQL(...))

https://github.com/a8m/rql

	{
	  "limit": 25,
	  "filter": {
	    "admin": false,
	    "created_at": {
	      "$gt": "2018-01-01T16:00:00.000Z",
	      "$lt": "2018-04-01T16:00:00.000Z"
	    }
	    "$or": [
	      { "address": "TLV" },
	      { "address": "NYC" }
	    ]
	  }
	  "sort": ["-created_at"]
	}
*/
func RqlParser(out *rql.Params, parser *rql.Parser) Parser {
	return func(c *fiber.Ctx) (any, error) {
		param, err := parser.Parse(c.Context().QueryArgs().QueryString())
		if err != nil {
			return nil, app_err.NewError(
				app_err.ErrValidationError,
				app_err.Join(c.UserContext(), err),
			)
		}

		*out = *param
		return nil, nil
	}
}
