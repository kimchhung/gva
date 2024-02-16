package response

import (
	"strings"

	app_err "github.com/kimchhung/gva/app/error"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"messages,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  []any  `json:"errors,omitempty"`

	HttpStatus int `json:"-"`
}

// Nothing to describe this fucking variable.
var IsProduction bool

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Response{
		Code:       app_err.ErrUnknownError.ErrorCode,
		Message:    app_err.ErrUnknownError.Message,
		HttpStatus: app_err.ErrUnknownError.HttpCode,
	}
	// Handle errors
	if e, ok := err.(validator.ValidationErrors); ok {
		resp.Code = app_err.ErrForbidden.ErrorCode
		resp.Message = app_err.ErrForbidden.Message
		resp.Errors = []any{removeTopStruct(e.Translate(trans))}
	} else if e, ok := err.(*app_err.Error); ok {
		resp.Code = e.ErrorCode
		resp.Message = e.Message
		resp.HttpStatus = e.HttpCode
	} else {
		resp.Message = err.Error()
	}

	if !IsProduction {
		log.Error().Err(err).Msg("From: Fiber's error handler")
	}

	return Resp(c, resp)
}

// A fuction to return beautiful responses.
func Resp(c *fiber.Ctx, resp Response) error {
	// Set status
	if resp.HttpStatus == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.HttpStatus)

	// Return JSON
	return c.JSON(resp)
}

// Remove unnecessary fields from validator message
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, msg := range fields {
		stripStruct := field[strings.Index(field, ".")+1:]
		res[stripStruct] = msg
	}
	return res
}
