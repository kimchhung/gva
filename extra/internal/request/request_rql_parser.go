package request

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/rql"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/sjson"
)

// filter[money][gt]:100 to map
func replaceToDot(v string) string {
	v = strings.ReplaceAll(v, "]", "")
	v = strings.ReplaceAll(v, "[", ".")
	return v
}

/*
Parse string queries depth to json string

	"filter[money][gt]:100" -> "{"filter":{"money":{"gt":"100"}}}"
*/
func splitPathAndValue(v string) (path string, value string) {
	//filter[money][gt]:100
	parts := strings.Split(v, "=")
	if len(parts) != 2 {
		return "", v
	}

	if !strings.Contains(parts[0], "[") {
		return parts[0], parts[1]
	}

	//filter.money.gt
	dotPath := replaceToDot(parts[0])
	return dotPath, parts[1]
}

/*
ParseUrlQuery

	queries := "filter[money][gt]=true&limit=10"
	expected := `{"filter":{"money":{"gt":true}},"limit":100}`
*/
func ParseUrlQuery(v string) string {
	values := strings.Split(v, "&")
	jsonStr := ""

	for _, part := range values {
		path, val := splitPathAndValue(part)

		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			jsonStr, _ = sjson.Set(jsonStr, path, f)
			continue
		}

		b, err := strconv.ParseBool(val)
		if err == nil {
			jsonStr, _ = sjson.Set(jsonStr, path, b)
			continue
		}

		jsonStr, _ = sjson.Set(jsonStr, path, val)
	}

	return jsonStr
}

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
		urlQ := string(c.Request().URI().QueryString())
		str := ParseUrlQuery(urlQ)

		param, err := parser.Parse([]byte(str))
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
note: for admin only

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
func MustRqlParser(table string, model any, mapColumnName ...map[string]string) *rql.Parser {
	var mapColumn map[string]string
	if len(mapColumnName) > 0 {
		mapColumn = mapColumnName[0]
	}

	return rql.MustNewParser(rql.Config{
		Model:         model,
		Log:           log.Debug().Msgf,
		ColumnFn:      rql.PascalToCamelCase,
		ColumnNameFn:  rql.CamelCaseToSnakeCase,
		MapColumnName: mapColumn,
		Table:         table,
		DoNotLog:      true,
		LimitMaxValue: 500,
	})
}
