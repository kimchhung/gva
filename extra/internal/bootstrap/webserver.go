package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/gosuri/uitable"
	appctx "github.com/kimchhung/gva/extra/app/common/context"
	"github.com/kimchhung/gva/extra/app/middleware"
	"github.com/kimchhung/gva/extra/app/router"
	"github.com/kimchhung/gva/extra/lang"
	"github.com/kimchhung/gva/extra/utils/color"
	"github.com/kimchhung/gva/extra/utils/validator"

	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"

	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func NewEcho(cfg *config.Config) *echo.Echo {
	// Setup Webserver

	// Pass production config to check it
	request.IsProduction = cfg.App.Production

	e := echo.New()
	e.Server.IdleTimeout = cfg.App.IdleTimeout * time.Second
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
			row = append(row, color.Cyan(url+"/doc"))
		}

		table.AddRow(row...)
	}

	if cfg.API.Admin.Enable {
		url := host + ":" + port + cfg.API.Admin.BasePath
		row := []any{"Admin", color.Cyan(url)}

		if cfg.Middleware.Swagger.Enable {
			row = append(row, color.Cyan(url+"/doc"))
		}

		table.AddRow(row...)
	}

	fmt.Printf("\n --------------- %s --------------- \n\n", color.White(cfg.App.Name)+" is running at the moment!")
	fmt.Println(table)
	fmt.Print("\n")
}

func printRoutes(routes []*echo.Route) {
	// Create a new table
	table := uitable.New()

	// Set the table headers

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
	lifecycle.Append(fx.StartHook(
		func(ctx context.Context) error {
			// Register middlewares & routes
			middlewares.Register()
			routers.Register(app, cfg, database)

			// Initailize validator and translator
			if err := lang.InitializeTranslator(); err != nil {
				return err
			}

			if err := validator.InitializeValidator(); err != nil {
				return err
			}

			// Connect db
			if err := database.ConnectDatabase(); err != nil {
				return err
			}

			if cfg.App.PrintRoutes {
				printRoutes(app.Routes())
			}

			printStartupMessage(cfg)

			// Listen the app (with TLS Support)
			if cfg.App.TLS.Enable {
				log.Info().Msg("TLS support was enabled.")

				go func() {
					if err := app.StartTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
						log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()
			} else {
				go func() {
					if err := app.Start(cfg.App.Port); err != nil {
						log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()
			}

			return nil
		},
	))

	lifecycle.Append(fx.StopHook(
		func(ctx context.Context) error {
			log.Info().Msg("Shutting down the app...")
			if err := app.Shutdown(ctx); err != nil {
				log.Panic().Err(err).Msg("")
			}

			log.Info().Msg("Running cleanup tasks...")
			log.Info().Msg("1- Shutdown the database")

			if err := database.ShutdownDatabase(); err != nil {
				return err
			}

			log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
			log.Info().Msg("\u001b[96mSee you again👋\u001b[0m")
			return nil
		},
	))

}
