package bootstrap

import (
	"fmt"
	"time"

	appctx "backend/app/common/context"
	"backend/env"
	"backend/utils/color"

	"github.com/gosuri/uitable"

	"backend/internal/request"

	"github.com/labstack/echo/v4"
)

func NewEcho(cfg *env.Config) *echo.Echo {
	// Setup Webserver

	// Pass production config to check it
	request.IsProduction = cfg.IsProd()

	e := echo.New()
	e.Server.IdleTimeout = time.Duration(cfg.App.IdleTimeout) * time.Second
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		appctx.ErrorHandler(err, c)
	}

	return e
}

func printStartupMessage(cfg *env.Config) {

	// Custom Startup Messages
	host, port := env.ParseAddress(cfg.App.Port)
	if host == "" {
		host = "http://localhost"
	}

	table := uitable.New()
	table.AddRow("API Module", "BasePath", "Document")

	if cfg.API.Bot.Enable {
		url := host + ":" + port + cfg.API.Bot.BasePath
		row := []any{"Bot", color.Cyan(url)}

		if cfg.Middleware.Swagger.Enable {
			row = append(row, color.Cyan(url+cfg.Middleware.Swagger.Path))
		}

		table.AddRow(row...)
	}

	if cfg.API.Admin.Enable {
		url := host + ":" + port + cfg.API.Admin.BasePath
		row := []any{"Admin", color.Cyan(url)}

		if cfg.Middleware.Swagger.Enable {
			row = append(row, color.Cyan(url+cfg.Middleware.Swagger.Path))
		}

		table.AddRow(row...)
	}

	fmt.Printf("\n --------------- %s --------------- \n\n", color.White(cfg.App.Name)+" is running at the moment!")
	fmt.Println(table)
	fmt.Print("\n")
}
