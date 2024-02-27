package request

import (
	"github.com/go-playground/validator/v10"
	app_err "github.com/kimchhung/gva/app/common/error"
	in_validator "github.com/kimchhung/gva/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
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
		resp.Code = app_err.ErrValidationError.ErrorCode
		resp.Message = app_err.ErrValidationError.Message
		resp.Errors = []any{in_validator.RemoveTopStruct(e.Translate(in_validator.Trans))}
		resp.HttpStatus = app_err.ErrForbidden.HttpCode
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
