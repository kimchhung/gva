package request

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	in_validator "github.com/kimchhung/gva/extra/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
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

// Nothing to describe this fucking variable.
var IsProduction bool

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	var resErr *app_err.Error

	if e, ok := err.(validator.ValidationErrors); ok {
		resErr = app_err.NewError(
			app_err.ErrValidationError,
			app_err.WithMessage(in_validator.RemoveTopStruct(e.Translate(in_validator.Trans))),
		)
	} else if e, ok := err.(*app_err.Error); ok {
		resErr = e
	} else {
		resErr = app_err.NewError(app_err.ErrUnknownError,
			// * count as server error can send message to webhook chat or save as log
			app_err.WithMessage(err.Error()),
		)
	}

	if !IsProduction {
		log.Error().Err(err).Msg("From: Fiber's error handler")
	}

	return Resp(c, Error(resErr))
}

func defaultReponse() *Response {
	return &Response{
		Code:       SuccessCode,
		Message:    SuccessMessage,
		HttpStatus: SuccessHttpCode,
		Data:       map[string]any{},
	}
}

// A fuction to return beautiful responses.
func Resp(c *fiber.Ctx, opt ReponseOption, opts ...ReponseOption) error {
	resp := defaultReponse()
	opt(resp)

	for _, op := range opts {
		op(resp)
	}

	return c.Status(resp.HttpStatus).JSON(resp)
}

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

			resp.Data = data
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

			resp.Data = data
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
