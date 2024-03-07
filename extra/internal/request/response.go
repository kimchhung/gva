package request

import (
	"github.com/go-playground/validator/v10"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/lang"
	in_validator "github.com/kimchhung/gva/extra/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Nothing to describe this fucking variable.
var IsProduction bool

// Default error handler
func ErrorHandler(c *fiber.Ctx, err error) error {
	var resErr *app_err.Error

	switch e := err.(type) {
	case validator.ValidationErrors:
		t := lang.Get(lang.WithFiberCtx(c))
		msg := in_validator.RemoveTopStruct(e.Translate(t))

		resErr = app_err.NewError(
			app_err.ErrValidationError,
			app_err.WithMessage(msg),
		)

		resErr.SetTranslated()

	case *app_err.Error:
		resErr = e
	default:
		resErr = app_err.NewError(app_err.ErrUnknownError,
			app_err.Join(c.UserContext(), err),
		)
	}

	if !IsProduction {
		log.Error().Err(err).Msg("From: Fiber's error handler")
	}

	resErr.Translate(c.UserContext())
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
