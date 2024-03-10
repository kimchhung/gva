package request

import (
	"github.com/gofiber/fiber/v2"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/rql"
	"github.com/rs/zerolog/log"
)

/*
@Required: meta.use(RQL(...))

https://github.com/a8m/rql

	{
	  "limit": 25,
	  "filter": {
	    "admin": false,
	    "createdAt": {
	      "$gt": "2018-01-01T16:00:00.000Z",
	      "$lt": "2018-04-01T16:00:00.000Z"
	    }
	    "$or": [
	      { "address": "TLV" },
	      { "address": "NYC" }
	    ]
	  }
	  "sort": ["-createdAt"]
	}
*/
func RqlParser(out *rql.Params, parser *rql.Parser) Parser {
	return func(c *fiber.Ctx) (any, error) {
		param, err := parser.Parse(
			[]byte(`{
				"limit": 25,
				"offset": 0,
				"sort": ["-createdAt"]
			}`),
		)

		if err != nil {
			return nil, app_err.NewError(
				app_err.ErrBadRequest,
				app_err.Join(err),
			)
		}

		*out = *param
		return nil, nil
	}
}

/*
Model is the resource definition. The parser is configured based on the its definition. For example, given the following struct definition:

	type User struct {
	    Age	 int	`rql:"filter,sort"`
	    Name string	`rql:"filter"`
	}

Json

	{
	  "limit": 25,
	  "offset": 0,
	  "filter": {
	    "isEnable": true,
	    "createdAt": {
	      "$gt": "2018-01-01T16:00:00.000Z",
	      "$lt": "2018-04-01T16:00:00.000Z"
	    }
	    "$or": [
	      { "isEnable": false },
	      { "address": "NYC" }
	    ]
	  }
	  "sort": ["-id","createdAt","+isEnable"]
	}
*/
func MustRqlParser(model any, table string, mapColumnName map[string]string) *rql.Parser {
	return rql.MustNewParser(rql.Config{
		Model:         model,
		Log:           log.Info().Msgf,
		DefaultSort:   []string{"-createdAt"}, // desc
		DoNotLog:      false,
		ColumnFn:      rql.PascalToCamelCase,
		ColumnNameFn:  rql.CamelCaseToSnakeCase,
		MapColumnName: mapColumnName,
		Table:         table,
	})
}
