package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/gosuri/uitable"
	appctx "github.com/kimchhung/gva/backend-echo/app/common/context"
	"github.com/kimchhung/gva/backend-echo/app/middleware"
	"github.com/kimchhung/gva/backend-echo/app/router"
	"github.com/kimchhung/gva/backend-echo/lang"
	"github.com/kimchhung/gva/backend-echo/utils/color"
	"github.com/kimchhung/gva/backend-echo/utils/validator"

	"github.com/kimchhung/gva/backend-echo/config"
	"github.com/kimchhung/gva/backend-echo/internal/bootstrap/database"

	"github.com/kimchhung/gva/backend-echo/internal/request"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func NewEcho(cfg *config.Config) *echo.Echo {
	// Setup Webserver

	// Pass production config to check it
	request.IsProduction = cfg.App.Production

	e := echo.New()
	e.Server.IdleTimeout = time.Duration(cfg.App.IdleTimeout) * time.Second
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		appctx.ErrorHandler(err, c)
	}

	return e
}

func printStartupMessage(cfg *config.Config) {

	// Custom Startup Messages
	host, port := config.ParseAddr(cfg.App.Port)
	if host == "" {
		host = "http://localhost"
	}

	table := uitable.New()
	table.AddRow("API Module", "BasePath", "Document")

	if cfg.API.Web.Enable {
		url := host + ":" + port + cfg.API.Web.BasePath
		row := []any{"Web", color.Cyan(url)}

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

func printRoutes(routes []*echo.Route) {
	table := uitable.New()
	table.AddRow("Method", "Path", "Name")
	for _, r := range routes {
		table.AddRow(color.MethodColor(r.Method), color.Yellow(r.Path), color.Cyan(r.Name))
	}

	table.Wrap = true
	// Print the table
	fmt.Print("\n ------------- Routes --------------- \n\n")
	fmt.Println(table)
	fmt.Print("\n")
}

func Start(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	routers *router.Router,
	app *echo.Echo,
	middlewares *middleware.Middleware,
	database *database.Database,
	log *zerolog.Logger,
) {

	isShuttingdown := false
	lifecycle.Append(fx.StartHook(
		func(ctx context.Context) error {
			// Register middlewares & routes
			middlewares.Register()
			routers.Register(app, cfg)

			// Initailize validator and translator
			if err := lang.InitializeTranslator(); err != nil {
				log.Panic().Err(err).Msg("failed to initialize translator!")
			}

			if err := validator.InitializeValidator(); err != nil {
				log.Panic().Err(err).Msg("failed to initialize validator!")
			}

			// Connect db
			if err := database.ConnectDatabase(); err != nil {
				log.Panic().Err(err).Msg("failed to connect to db!")
			}

			app.Server.RegisterOnShutdown(func() {
				log.Info().Msg("1- Shutdown the database")
				if err := database.ShutdownDatabase(); err != nil {
					log.Err(err).Msg("failed to shutdown db!")
				}
			})

			if cfg.App.PrintRoutes {
				printRoutes(app.Routes())
			}

			printStartupMessage(cfg)

			// Listen the app (with TLS Support)
			if cfg.App.TLS.Auto {
				log.Info().Msg("Auto TLS support was enabled.")

				go func() {
					if err := app.StartAutoTLS(cfg.App.Port); err != nil && !isShuttingdown {
						log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()
			} else if cfg.App.TLS.Enable {
				log.Info().Msg("TLS support was enabled.")

				go func() {
					if err := app.StartTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil && !isShuttingdown {
						log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()
			} else {
				go func() {
					if err := app.Start(cfg.App.Port); err != nil && !isShuttingdown {
						log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()
			}

			return nil
		},
	))

	lifecycle.Append(fx.StopHook(
		func(ctx context.Context) error {
			isShuttingdown = true
			ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.App.ShutdownTimeout)*time.Second)
			defer cancel()

			log.Info().Msg("Shutting down the app...")
			log.Info().Msg("Running cleanup tasks...")

			if err := app.Shutdown(ctx); err != nil {
				log.Panic().Err(err).Msg("")
			}

			log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
			log.Info().Msg("\u001b[96mSee you again👋\u001b[0m")
			return nil
		},
	))

}
