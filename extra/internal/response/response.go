package response

import (
	"encoding/json"
	"net/http"
	"reflect"

	app_err "github.com/kimchhung/gva/extra/app/common/error"
)

const (
	SuccessCode     = http.StatusOK
	SuccessMessage  = "OK"
	SuccessHttpCode = http.StatusOK
)

type (
	ReponseOption func(resp *Response)

	Response struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
		Data    any    `json:"data,omitempty"`

		HttpStatus int `json:"-"`
	}
)

/*
return as data in format

Case: Map | Struct

	{
		code:[SuccessCode],
		message:"[SuccessMessage]"
		data:{
			key:value
		}
	}

Case: Slices

	{
		code:[SuccessCode],
		message:"[SuccessMessage]"
		data:{
			list:[]
			[key]:value
		}
	}
*/
func Data(data any, keyValues ...map[string]any) ReponseOption {
	return func(resp *Response) {
		if reflect.TypeOf(data).Kind() == reflect.Slice {
			// Directly assign the slice to the map under a specific key, e.g., "list"
			obj := map[string]any{
				"list": data,
			}

			for _, keyValue := range keyValues {
				for k, v := range keyValue {
					obj[k] = v
				}
			}

			resp.Data = obj
			return
		}

		if len(keyValues) > 0 {
			obj := map[string]any{}

			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}

			if err := json.Unmarshal(b, &obj); err != nil {
				panic(err)
			}

			for _, keyValue := range keyValues {
				for k, v := range keyValue {
					obj[k] = v
				}
			}

			resp.Data = obj
			return
		}

		resp.Data = data
	}
}

/*
return as data

	{
		code:200,
		message:"ok"
	}
*/
func Error(err *app_err.Error) ReponseOption {
	return func(resp *Response) {
		resp.Code = err.ErrorCode
		resp.Message = err.Message
		resp.HttpStatus = err.HttpCode
	}
}

/*
return as data

	{
		code:200,
		message:"ok"
	}
*/
func Message(msg string) ReponseOption {
	return func(resp *Response) {
		resp.Message = msg
	}
}
