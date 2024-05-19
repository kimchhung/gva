package request

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/rql"
	"github.com/kimchhung/gva/extra/utils"
	"github.com/kimchhung/gva/extra/utils/json"
	ustrings "github.com/kimchhung/gva/extra/utils/strings"

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

func isValidKey(key string) bool {
	for _, validKey := range []string{"limit", "offset", "filter", "select", "sort"} {
		if strings.Contains(validKey, key) {
			return true
		}
	}
	return false
}

/*
ParseUrlQuery

	queries := "filter[money][gt]=true&limit=10"
	expected := `{"filter":{"money":{"gt":true}},"limit":100}`
*/
func ParseUrlQuery(v string) string {
	values := strings.Split(v, "&")
	jsonStr := json.NewEmptyObject()
	var err error

	for _, part := range values {
		path, val := splitPathAndValue(part)
		if !isValidKey(path) {
			continue
		}

		if strings.ContainsRune(val, ',') {
			jsonStr, err = sjson.SetBytes(jsonStr, path, strings.Split(val, ","))
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
