package bootstrap

import (
	"context"
	"os"
	"runtime"
	"strings"

	"time"

	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"

	"github.com/kimchhung/gva/extra/app/middleware"
	"github.com/kimchhung/gva/extra/app/router"
	"github.com/kimchhung/gva/extra/lang"
	"github.com/kimchhung/gva/extra/utils/validator"

	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	rerror "github.com/kimchhung/gva/extra/internal/response/error"

	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func NewFiber(cfg *config.Config) *fiber.App {
	// Setup Webserver
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		ErrorHandler:          rerror.ErrorHandler,
		DisableStartupMessage: true,
		Immutable:             true,
	})

	// Pass production config to check it
	request.IsProduction = cfg.App.Production
	return app
}

func printStartupMessage(cfg *config.Config, fiber *fiber.App) {
	// Information message
	log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

	// Custom Startup Messages
	host, port := config.ParseAddr(cfg.App.Port)
	if host == "" {
		if fiber.Config().Network == "tcp6" {
			host = "[::1]"
		} else {
			host = "0.0.0.0"
		}
	}

	// ASCII Art
	ascii, err := os.ReadFile("./storage/ascii_art.txt")
	if err != nil {
		log.Debug().Err(err).Msg("An unknown error occurred when to print ASCII art!")
	}
	for _, line := range strings.Split(futils.UnsafeString(ascii), "\n") {
		log.Info().Msg(line)
	}

	// Debug informations
	if !cfg.App.Production {
		prefork := "Enabled"
		procs := runtime.GOMAXPROCS(0)
		if !cfg.App.Prefork {
			procs = 1
			prefork = "Disabled"
		}

		log.Info().Msgf("Version: %s", "-")
		log.Info().Msgf("Host: %s", host)
		log.Info().Msgf("Port: %s", port)
		log.Info().Msgf("Prefork: %s", prefork)
		log.Info().Msgf("Handlers: %d", fiber.HandlersCount())
		log.Info().Msgf("Processes: %d", procs)
		log.Info().Msgf("PID: %d", os.Getpid())
	}
}

func Start(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	routers *router.Router,
	fiber *fiber.App,
	middlewares *middleware.Middleware,
	database *database.Database,
	log *zerolog.Logger,
) {
	lifecycle.Append(fx.StartHook(
		func(ctx context.Context) error {
			// Register middlewares & routes
			middlewares.Register()
			routers.Register(fiber, cfg, database)

			// Listen the app (with TLS Support)
			if cfg.App.TLS.Enable {
				log.Debug().Msg("TLS support was enabled.")

				if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
					log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
				}
			}

			go func() {
				if err := fiber.Listen(cfg.App.Port); err != nil {
					log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
				}
			}()

			go func() {
				time.Sleep(time.Second)
				printStartupMessage(cfg, fiber)
			}()

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

			return nil
		},
	))

	lifecycle.Append(fx.StopHook(
		func(ctx context.Context) error {
			log.Info().Msg("Shutting down the app...")
			if err := fiber.Shutdown(); err != nil {
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
