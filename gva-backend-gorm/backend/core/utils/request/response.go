package request

import (
	"backend/core/utils/response"
	"backend/internal/logger"

	"github.com/labstack/echo/v4"
)

// Nothing to describe this fucking variable.

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
func Response(c echo.Context, opts ...response.ReponseOption) error {
	resp := response.New(opts...)
	hook, err := GetHook(c.Request().Context())
	if err == nil {
		for _, callbacks := range hook.onBeforeResponses {
			callbacks(c, resp)
		}

		if r := recover(); r != nil {
			logger.Log("panic", r)
		}
	}

	return resp.Parse(c)
}
