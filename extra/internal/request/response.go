package request

import (
	"github.com/go-playground/validator/v10"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/response"
	in_validator "github.com/kimchhung/gva/extra/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
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

	return Response(c, response.Error(resErr))
}

func defaultResponse() *response.Response {
	return &response.Response{
		Code:       response.SuccessCode,
		Message:    response.SuccessMessage,
		HttpStatus: response.SuccessHttpCode,
		Data:       map[string]any{},
	}
}

// A fuction to return beautiful responses.
func Response(c *fiber.Ctx, opt response.ReponseOption, opts ...response.ReponseOption) error {
	resp := defaultResponse()
	opt(resp)

	for _, op := range opts {
		op(resp)
	}

	if v, ok := resp.Data.(map[string]any); ok && len(v) == 0 {
		resp.Data = nil
	}

	return c.Status(resp.HttpStatus).JSON(resp)
}
