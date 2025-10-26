package bootstrap

import (
	"context"
	"fmt"
	"time"

	coreerror "backend/core/error"
	"backend/core/lang"
	"backend/core/utils/color"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/env"

	in_validator "backend/core/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gosuri/uitable"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

func NewEcho(cfg *env.Config, logger *zap.Logger) *echo.Echo {
	// Setup Webserver
	log := logger.Named("Echo")
	e := echo.New()
	e.Server.IdleTimeout = time.Duration(cfg.App.IdleTimeout) * time.Second

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		coreErr := tranformError(c.Request().Context(), err)

		if !coreErr.IsPublic() {
			log.Error(coreErr.Message,
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.Int("errorCode", coreErr.ErrorCode),
				zap.Int("httpCode", coreErr.HttpCode),
				zap.Error(err),
			)
		}

		request.Response(c, response.Error(coreErr))
	}

	return e
}

// Default error handler
func tranformError(ctx context.Context, err error) (resErr *coreerror.Error) {
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
		if e.IsTranslated() {
			resErr = e
		} else if lang.IsInitialized() {
			resErr = coreerror.NewError(e, coreerror.Translate(translator))
		}

	case *echo.HTTPError:
		// wrong routing .....
		resErr = coreerror.ErrBadRequest.Copy(
			coreerror.Translate(translator),
			coreerror.AppendMessage(e.Error()),
		)

	default:

		// unexpected error, crashed etc...
		resErr = coreerror.ErrUnknownError.Copy(coreerror.Translate(translator))

	}

	return resErr
}

func printStartupMessage(cfg *env.Config) {

	// Custom Startup Messages
	host, port := env.ParseAddress(cfg.App.Port)
	if host == "" {
		host = "http://localhost"
	}

	table := uitable.New()
	table.AddRow("API Module", "BasePath", "Document")

	if cfg.Bot.Enable {
		url := host + ":" + port + cfg.Bot.BasePath
		row := []any{"Bot", color.Cyan(url)}

		if cfg.Middleware.Swagger.Enable {
			row = append(row, color.Cyan(url+cfg.Middleware.Swagger.Path))
		}

		table.AddRow(row...)
	}

	if cfg.Admin.Enable {
		url := host + ":" + port + cfg.Admin.BasePath
		row := []any{"Admin", color.Cyan(url)}

		if cfg.Middleware.Swagger.Enable {
			row = append(row, color.Cyan(url+cfg.Middleware.Swagger.Path))
		}

		table.AddRow(row...)
	}

	if cfg.Web.Enable {
		url := host + ":" + port + cfg.Web.BasePath
		row := []any{"Web", color.Cyan(url)}

		if cfg.Middleware.Swagger.Enable {
			row = append(row, color.Cyan(url+cfg.Middleware.Swagger.Path))
		}

		table.AddRow(row...)
	}

	fmt.Printf("\n --------------- %s --------------- \n\n", color.White(cfg.App.Name)+" is running at the moment!")
	fmt.Println(table)
	fmt.Print("\n")
}
