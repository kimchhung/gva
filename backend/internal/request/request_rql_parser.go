package request

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	appctx "github.com/gva/app/common/context"
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
	whiteList := make(url.Values, len(values))
	for _, k := range []string{"limit", "offset", "select", "sort", "filter"} {
		whiteList[k] = values[k]
	}

	jsonStr := json.NewEmptyObject()
	var err error

	for path, vals := range whiteList {
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
		params := c.QueryParams()
		query := &rql.Query{
			Limit:  25,
			Offset: 0,
		}

		if str := params.Get("limit"); str != "" {
			if num, err := strconv.Atoi(str); err == nil {
				query.Limit = num
			}
		}

		if str := params.Get("offset"); str != "" {
			if num, err := strconv.Atoi(str); err == nil {
				query.Offset = num
			}
		}

		if str := params.Get("select"); str != "" {
			for _, sel := range strings.Split(str, ",") {
				sel = strings.TrimSpace(sel)
				if sel != "" {
					query.Select = append(query.Select, sel)
				}
			}
		}

		if str := params.Get("filter"); str != "" {
			decodedBytes, err := base64.URLEncoding.DecodeString(str)
			if err != nil {
				return nil, app_err.NewError(
					app_err.ErrBadRequest,
					app_err.Join(err),
				)
			}

			if err := json.JSON(decodedBytes).Out(&query.Filter); err != nil {
				return nil, app_err.NewError(
					app_err.ErrBadRequest,
					app_err.Join(err),
				)
			}
		}

		rqlParam, err := parser.ParseQuery(query)
		if err != nil {
			return nil, app_err.NewError(
				app_err.ErrBadRequest,
				app_err.Join(err),
			)
		}

		rctx := appctx.MustRequestContext(c.Request().Context())
		if !rctx.IsProd() {
			c.Response().Header().Set("X-Filter", fmt.Sprintf("%v", json.MustJSON(query.Filter)))
		}

		*out = *rqlParam
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
