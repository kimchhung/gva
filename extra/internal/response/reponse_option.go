package response

import (
	"encoding/json"
	"reflect"

	apperror "github.com/kimchhung/gva/extra/app/common/error"
)

/*
return as data

	{
		code:0,
		message:"OK"
		data:any
	}
*/
func Data(data any) ReponseOption {
	return func(resp *Response) {
		resp.Data = data
	}
}

/*
return as data in format

Case: Map | Struct

	{
		code:0,
		message:"OK"
		meta:{
			key:value
		}
	}

Case: Slices

	{
		code:0,
		message:"OK"
		meta:{
			list:[]
			[key]:value
		}
	}
*/
func Meta(meta any, keyValues ...map[string]any) ReponseOption {
	return func(resp *Response) {
		if reflect.TypeOf(meta).Kind() == reflect.Slice {
			// Directly assign the slice to the map under a specific key, e.g., "list"
			obj := map[string]any{
				"list": meta,
			}

			for _, keyValue := range keyValues {
				for k, v := range keyValue {
					obj[k] = v
				}
			}

			resp.Meta = obj
			return
		}

		if len(keyValues) > 0 {
			obj := map[string]any{}

			b, err := json.Marshal(meta)
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

			resp.Meta = obj
			return
		}

		resp.Meta = meta
	}
}

/*
Ovewrite default status

	Status(http.StatusCreated)
*/
func Status(status int) ReponseOption {
	return func(resp *Response) {
		resp.httpStatus = status
	}
}

func Error(err *apperror.Error) ReponseOption {
	return func(resp *Response) {
		resp.Code = err.ErrorCode
		resp.Message = err.Error()
		resp.httpStatus = err.HttpCode
	}
}

func Message(msg string) ReponseOption {
	return func(resp *Response) {
		resp.Message = msg
	}
}
