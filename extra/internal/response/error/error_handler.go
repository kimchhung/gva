package rerror

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	app_err "github.com/kimchhung/gva/extra/app/common/error"
	in_validator "github.com/kimchhung/gva/extra/utils/validator"

	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/lang"
)

// Default error handler
func ParseError(c *fiber.Ctx, err error) (*app_err.Error, error) {

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
		if e.IsDisableTranslate() {
			resErr = e
		} else {
			resErr = app_err.NewError(e, app_err.MessageFunc(
				func(message string) string {
					return lang.T(lang.FiberCtx(c), message)
				},
			))
		}

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
		// StackHandler will invoke too
		resErr = app_err.NewError(
			app_err.ErrUnknownError,
			app_err.MessageFunc(
				func(message string) string {
					return lang.T(lang.FiberCtx(c), message)
				},
			),
		)
	}

	return resErr, response.New(response.Error(resErr)).Parse(c)
}
