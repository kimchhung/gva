package bootstrap

import (
	"context"

	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"

	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/app/database/seeds"
	"github.com/kimchhung/gva/extra/app/middleware"
	"github.com/kimchhung/gva/extra/app/module"

	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	rerror "github.com/kimchhung/gva/extra/internal/response/error"

	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/rs/zerolog"
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

func Start(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	routers *module.Router,
	fiber *fiber.App,
	middlewares *middleware.Middleware,
	database *database.Database,
	log *zerolog.Logger,
) {
	onStart := func(ctx context.Context) error {
		// Register middlewares & routes
		middlewares.Register()
		routers.Register(fiber, cfg)

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

		// Information message
		log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

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

		database.ConnectDatabase()

		// dependencies for seeding
		ctx = context.WithValue(ctx, config.Config{}, cfg)
		ctx = context.WithValue(ctx, services.PasswordService{}, services.NewPasswordService(cfg))

		database.SeedModels(ctx,
			seeds.RouterSeeder{},
			seeds.SuperAdminSeeder{},
		)

		return nil
	}

	onStop := func(ctx context.Context) error {
		log.Info().Msg("Shutting down the app...")
		if err := fiber.Shutdown(); err != nil {
			log.Panic().Err(err).Msg("")
		}

		log.Info().Msg("Running cleanup tasks...")
		log.Info().Msg("1- Shutdown the database")

		database.ShutdownDatabase()

		log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
		log.Info().Msg("\u001b[96mSee you again👋\u001b[0m")

		return nil
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: onStart,
			OnStop:  onStop,
		},
	)
}
