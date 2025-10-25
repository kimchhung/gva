package request

import (
	"backend/core/utils/json"
	"backend/internal/pagi"
	"backend/internal/relay"
	"backend/internal/relay/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/sjson"
)

type (
	Parser      func(c echo.Context) (any, error)
	RawQueryDto struct {
		Page    string                    `json:"page" query:"page"`
		Limit   string                    `json:"limit" query:"limit"`
		Filters map[string]map[string]any `json:"filters" query:"filters"`
		Sorts   string                    `json:"sorts" query:"sorts"`
		Search  string                    `json:"search" query:"search"`
		Selects string                    `json:"selects" query:"selects"`
	}
	RawCursorDto struct {
		First   string                    `json:"first" query:"first"`
		Last    string                    `json:"last" query:"last"`
		After   string                    `json:"after" query:"after"`
		Before  string                    `json:"before" query:"before"`
		Sorts   string                    `json:"sorts" query:"sorts"`
		IsCount string                    `json:"isCount" query:"isCount"`
		Filters map[string]map[string]any `json:"filters" query:"filters"`
	}
)

var binder = &echo.DefaultBinder{}

/*
BodyParser(&Person{})

	type User struct {
	  ID string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	}

Swagger

	// @Param 		<name> <params-type> <value-type> <required> <description>
	// @Param 		info body dto.MenuRequest true "Route Info"
*/
func BodyParser(out any) Parser {
	return func(c echo.Context) (any, error) {
		return out, binder.BindBody(c, out)
	}
}

/*
Only for GET/DELETE methods

	type Person struct {
		    Name     string     `query:"name"`
		    Age     int     `query:"age"`
	}

	{
		persion := &Person{}
		QueryParser(persion)
	}

Swagger

	// @Param       <name> <params-type> <value-type> <required> <description>
	// @Param       person query dto.Person true "person info"
*/
func QueryParser(out any) Parser {
	return func(c echo.Context) (any, error) {
		return out, binder.BindQueryParams(c, out)
	}
}

/*
ParamsParser(&param{})

	param := new(struct {ID uint `param:"id"`})

Swagger

	// @Param       <name> <params-type> <value-type> <required> <description>
	// @Param       id path int true "Route ID"
*/
func ParamsParser(out any) Parser {
	return func(c echo.Context) (any, error) {
		return out, binder.BindPathParams(c, out)
	}
}

/*
PaginateParser(&pagi.QueryDto{})
*/
func PaginateParser(out *pagi.QueryDto) Parser {
	return func(c echo.Context) (any, error) {
		var (
			dto     RawQueryDto
			jsonraw json.JSON
		)
		params := c.QueryParams()

		for key, values := range params {
			key = replaceToDot(key)
			value := values[0]

			jsonraw, _ = sjson.SetBytes(jsonraw, key, value)
		}

		err := json.MustJSON(jsonraw).Out(&dto)
		if err != nil {
			return nil, err
		}

		var filter pagi.FilterMap
		if err := copier.Copy(&filter, dto.Filters); err != nil {
			return nil, fmt.Errorf("error copying filter: %v", err)
		}

		out.Filters = filter
		out.Page, _ = strconv.Atoi(dto.Page)
		if out.Page == 0 {
			out.Page = 1
		}
		out.Limit, _ = strconv.Atoi(dto.Limit)
		if out.Limit == 0 {
			out.Limit = 20
		}
		out.Sorts = strings.Split(dto.Sorts, ",")
		out.DefaultSelects(dto.Selects)
		out.Search = dto.Search

		return out, nil
	}
}

func CursorParser(out *relay.CursorDto) Parser {
	return func(c echo.Context) (any, error) {
		var (
			dto     RawCursorDto
			jsonraw json.JSON
		)
		params := c.QueryParams()

		for key, values := range params {
			key = replaceToDot(key)
			value := values[0]

			jsonraw, _ = sjson.SetBytes(jsonraw, key, value)
		}

		err := json.MustJSON(jsonraw).Out(&dto)
		if err != nil {
			return nil, err
		}

		var filter map[string]map[string]any
		if err := copier.Copy(&filter, dto.Filters); err != nil {
			return nil, fmt.Errorf("error copying filter: %v", err)
		}

		out.Filters = filter

		first, err := strconv.Atoi(dto.First)
		if err == nil {
			out.First = &first
		}

		last, err := strconv.Atoi(dto.Last)
		if err == nil {
			out.Last = &last
		}

		if out.First == nil && out.Last == nil {
			out.First = utils.ToPointer(20)
		}

		out.Sorts = strings.Split(dto.Sorts, ",")

		if dto.After != "" {
			out.After = &dto.After
		}

		if dto.Before != "" {
			out.Before = &dto.Before
		}

		return out, nil
	}
}

// filter[money][gt]:100 to map
func replaceToDot(v string) string {
	v = strings.ReplaceAll(v, "]", "")
	v = strings.ReplaceAll(v, "[", ".")
	return v
}
