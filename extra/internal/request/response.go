package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/pkg/errors"

	"github.com/kimchhung/gva/extra/lang"
	in_validator "github.com/kimchhung/gva/extra/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Nothing to describe this fucking variable.
var IsProduction bool

func init() {
	zerolog.ErrorStackMarshaler = MarshalStackSkip(3)
}

func StackHandler(c *fiber.Ctx, panic any) {
	if _, ok := panic.(*app_err.Error); ok {
		return
	}

	err := errors.Wrap(fmt.Errorf("%v", panic), "from panic")
	log.Error().Stack().Err(err).Msg("Stack handler")
}

// Default error handler
func ErrorHandler(c *fiber.Ctx, err error) error {
	var resErr *app_err.Error

	switch e := err.(type) {
	case validator.ValidationErrors:
		// error from request input validation
		t := lang.GetTranslator(lang.FiberCtx(c))
		translatedMsg := in_validator.RemoveTopStruct(e.Translate(t))
		resErr = app_err.NewError(
			app_err.ErrValidationError,
			app_err.Message(translatedMsg),
		)

	case *app_err.Error:
		// throw from logical error for user to see
		resErr = e
		resErr.Message = lang.T(lang.FiberCtx(c), resErr.Message)

	case *fiber.Error:
		// wrong routing .....
		resErr = app_err.NewError(
			app_err.ErrBadRequest,
			app_err.MessageFunc(
				func(message string) string {
					return lang.T(lang.FiberCtx(c), message)
				},
			),
			app_err.Join(err),
		)
		resErr.HttpCode = e.Code

	default:
		// unexpected error, crashed etc...
		resErr = app_err.NewError(
			app_err.ErrUnknownError,
			app_err.MessageFunc(
				func(message string) string {
					return lang.T(lang.FiberCtx(c), message)
				},
			),
			app_err.Join(err),
		)
	}

	if !IsProduction {
		log.Err(err).Str("from", "fiber error").Msg("Error handler")
	}

	return Response(c, response.Error(resErr))
}

// A fuction to return beautiful and structured responses.
func Response(c *fiber.Ctx, opt response.ReponseOption, opts ...response.ReponseOption) error {
	return response.New(opt, opts...).Parse(c)
}
