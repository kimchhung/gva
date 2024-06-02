package rerror

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	apperror "github.com/kimchhung/gva/extra/app/common/error"
	in_validator "github.com/kimchhung/gva/extra/utils/validator"

	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/lang"
)

// Default error handler
func ParseError(c echo.Context, err error) (*apperror.Error, error) {

	var resErr *apperror.Error

	switch e := err.(type) {
	case validator.ValidationErrors:
		// error from request input validation
		t := lang.GetTranslator(lang.ForContext(c.Request().Context()))

		translatedMsg := in_validator.RemoveTopStruct(e.Translate(t))
		resErr = apperror.NewError(
			apperror.ErrValidationError,
			apperror.Message(translatedMsg),
		)

	case *apperror.Error:
		// throw from logical error for user to see
		if e.IsDisableTranslate() {
			resErr = e
		} else {
			resErr = apperror.NewError(e, apperror.MessageFunc(
				func(message string) string {
					return lang.T(lang.ForContext(c.Request().Context()), message)
				},
			))
		}

	case *echo.HTTPError:
		// wrong routing .....
		resErr = apperror.NewError(
			apperror.ErrBadRequest,
			apperror.MessageFunc(
				func(message string) string {
					return lang.T(lang.ForContext(c.Request().Context()), message)
				},
			),
			apperror.Join(err),
		)
		resErr.HttpCode = e.Code

	default:

		// unexpected error, crashed etc...
		// StackHandler will invoke too
		resErr = apperror.NewError(
			apperror.ErrUnknownError,
			apperror.MessageFunc(
				func(message string) string {
					return lang.T(lang.ForContext(c.Request().Context()), message)
				},
			),
		)
	}

	return resErr, response.New(response.Error(resErr)).Parse(c)
}
