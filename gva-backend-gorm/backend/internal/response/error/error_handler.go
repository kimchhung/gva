package rerror

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	apperror "backend/app/common/error"
	in_validator "backend/internal/bootstrap/validator"

	"backend/internal/bootstrap/lang"
	"backend/internal/response"
)

// Default error handler
func SanitizeError(ctx context.Context, err error) *apperror.Error {
	var resErr *apperror.Error

	switch e := err.(type) {
	case validator.ValidationErrors:
		// error from request input validation
		t := lang.DefaultTranslator().GetTranslator(lang.ForContext(ctx))
		translatedMsg := in_validator.RemoveTopStruct(e.Translate(t))

		resErr = apperror.NewError(
			apperror.ErrValidationError,
			apperror.WithMessage(translatedMsg),
		)

	case *apperror.Error:
		// throw from logical error for user to see
		if e.IsDisableTranslate() {
			resErr = e
		} else {
			resErr = apperror.NewError(e, apperror.WithMessageFunc(
				func(message string) string {
					return lang.DefaultTranslator().T(lang.ForContext(ctx), message)
				},
			))
		}

	case *echo.HTTPError:
		// wrong routing .....
		resErr = apperror.NewError(
			apperror.ErrBadRequest,
			apperror.WithMessageFunc(
				func(message string) string {
					return lang.DefaultTranslator().T(lang.ForContext(ctx), message)
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
			apperror.WithMessageFunc(
				func(message string) string {
					return lang.DefaultTranslator().T(lang.ForContext(ctx), message)
				},
			),
		)
	}

	return resErr
}

// Default error handler
func ParseError(c echo.Context, err error) (*apperror.Error, error) {
	resErr := SanitizeError(c.Request().Context(), err)
	return resErr, response.New(response.Error(resErr)).Parse(c)
}
