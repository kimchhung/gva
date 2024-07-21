package request

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	app_err "github.com/gva/app/common/error"
	"github.com/gva/internal/rql"
	"github.com/gva/utils"
	"github.com/gva/utils/json"
	ustrings "github.com/gva/utils/strings"

	"github.com/labstack/echo/v4"

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

	queries := "filter[money][$gt]=true&limit=10"
	expected := `{"filter":{"money":{"$gt":true}},"limit":100}`
*/
func ParseUrlValue(values url.Values) string {
	jsonStr := json.NewEmptyObject()
	var err error

	for path, vals := range values {
		path = replaceToDot(path)
		fmt.Println("path: ", path)
		fmt.Println("pathv: ", vals)

		val := ""
		if len(vals) == 1 {
			val = vals[0]
		} else {
			jsonStr, err = sjson.SetBytes(jsonStr, path, vals)
			utils.PanicIfErr("ParseUrlQuery", err)
			continue
		}

		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			jsonStr, err = sjson.SetBytes(jsonStr, path, f)
			utils.PanicIfErr("ParseUrlQuery", err)

			continue
		}

		b, err := strconv.ParseBool(val)
		if err == nil {
			jsonStr, err = sjson.SetBytes(jsonStr, path, b)
			utils.PanicIfErr("ParseUrlQuery", err)
			continue
		}

		jsonStr, err = sjson.SetBytes(jsonStr, path, val)
		utils.PanicIfErr("ParseUrlQuery", err)
	}

	fmt.Println("result: ", jsonStr.String())

	return jsonStr.String()
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
func RqlQueryParser(out *rql.Params, parser *rql.Parser) Parser {
	return func(c echo.Context) (any, error) {
		var (
			param *rql.Params
			err   error
		)

		urlValue := c.QueryParams()
		if urlValue.Encode() != "" {
			str := ParseUrlValue(urlValue)
			param, err = parser.Parse([]byte(str))
		} else {
			param, err = parser.ParseQuery(
				&rql.Query{
					Limit:  20,
					Offset: 0,
				},
			)
		}

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
func MustRqlParser(config rql.Config) *rql.Parser {
	config.FieldSep = "__"
	config.InterpretFieldSepAsNestedJsonbObjectMysql = true
	config.NameFn = func(s string) string {
		s = ustrings.ToCamel(s, config.FieldSep)
		return strings.ReplaceAll(s, ".", config.FieldSep)
	}
	config.DoNotLog = false
	return rql.MustNewParser(config)
}
