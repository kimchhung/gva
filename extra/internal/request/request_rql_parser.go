package request

import (
	"github.com/gofiber/fiber/v2"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/rql"
)

type rqlParserKey struct{}

/*
@Usage:

	meta.use(RQL(...))

https://github.com/a8m/rql
*/
func RQL(config rql.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(rqlParserKey{}, rql.MustNewParser(config))
		return c.Next()
	}
}

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
func RqlParser(out *rql.Params) Parser {
	return func(c *fiber.Ctx) (any, error) {
		parser, ok := c.Locals(rqlParserKey{}).(*rql.Parser)
		if !ok {
			panic("RqlMiddleware is required")
		}

		param, err := parser.Parse(c.Context().QueryArgs().QueryString())
		if err != nil {
			return nil, app_err.NewError(
				app_err.ErrValidationError,
				app_err.Join(err),
			)
		}

		*out = *param

		// don't need validate so return nil
		return nil, nil
	}
}

/*
https://github.com/a8m/rql

@Required:

	meta.use(RQL(...))



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
func RqlBodyParser(out *rql.Params) Parser {
	return func(c *fiber.Ctx) (any, error) {
		parser, ok := c.Locals(rqlParserKey{}).(*rql.Parser)
		if !ok {
			panic("RqlMiddleware is required")
		}

		param, err := parser.Parse(c.Body())
		if err != nil {
			return nil, app_err.NewError(
				app_err.ErrValidationError,
				app_err.Join(err),
			)
		}

		*out = *param
		// don't need validate so return nil
		return nil, nil
	}
}
