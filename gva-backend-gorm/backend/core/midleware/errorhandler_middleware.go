package coremiddleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	coreerror "backend/core/error"
	"backend/core/lang"
	in_validator "backend/core/validator"
)

// Default error handler
func sanitizeError(ctx context.Context, err error, log *zap.Logger) (resErr *coreerror.Error) {
	translator := func(message string) string {
		if lang.IsInitialized() {
			return lang.DefaultTranslator().T(lang.ForContext(ctx), message)
		}
		return message
	}

	switch e := err.(type) {
	case validator.ValidationErrors:
		translator = func(message string) string {
			if lang.IsInitialized() {
				t := lang.DefaultTranslator().GetTranslator(lang.ForContext(ctx))
				return in_validator.RemoveTopStruct(e.Translate(t))
			}
			return message
		}

		resErr = coreerror.ErrValidationError.Copy(coreerror.Translate(translator))

	case *coreerror.Error:
		// throw from logical error for user to see
		if e.IsDisableTranslate() {
			resErr = e
		} else if lang.IsInitialized() {
			resErr = coreerror.NewError(e, coreerror.Translate(translator))
		}

	case *echo.HTTPError:
		// wrong routing .....
		resErr = coreerror.ErrBadRequest.Copy(coreerror.Translate(translator), coreerror.Join(err))

	default:

		// unexpected error, crashed etc...
		resErr = coreerror.ErrUnknownError.Copy(coreerror.Translate(translator))

		if log != nil && !resErr.IsPublic() {
			log.Error(resErr.Message,
				zap.Int("errorCode", resErr.ErrorCode),
				zap.Error(err),
			)
		}
	}

	return resErr
}

func (m *Middleware) ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			returnErr := next(c)
			ctx := c.Request().Context()

			if r := recover(); r != nil {
				if r == http.ErrAbortHandler {
					panic(r)
				}
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				var log *zap.Logger
				if !m.cfg.IsProd() {
					log = m.log.Named("Error")
				}

				coreError := sanitizeError(ctx, err, log)
				return coreError
			}

			return returnErr
		}
	}
}
