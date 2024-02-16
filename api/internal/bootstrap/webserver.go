package bootstrap

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/kimchhung/gva/app/middleware"
	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/bootstrap/database"
	"github.com/kimchhung/gva/internal/control_route"
	"github.com/kimchhung/gva/utils/response"

	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFiber(cfg *config.Config) *fiber.App {
	// Setup Webserver
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
	})

	// Pass production config to check it
	response.IsProduction = cfg.App.Production

	return app
}

func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, routers control_route.Router, middlewares *middleware.Middleware, database *database.Database, log *zap.Logger) {
	sugar := log.Sugar()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				routers.Register()
				fmt.Print("register middlewares")
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
					log.Error("An unknown error occurred when to print ASCII art!", zap.Error(err))
				}

				fmt.Println("")
				for _, line := range strings.Split(futils.UnsafeString(ascii), "\n") {
					fmt.Println(line)
				}
				fmt.Println("")

				// Information message
				sugar.Info(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					sugar.Debugln("Version: %s", "-")
					sugar.Debugln("Host: %s", host)
					sugar.Debugln("Port: %s", port)
					sugar.Debugln("Prefork: %s", prefork)
					sugar.Debugln("Handlers: %d", fiber.HandlersCount())
					sugar.Debugln("Processes: %d", procs)
					sugar.Debugln("PID: %d", os.Getpid())
				}

				// Listen the app (with TLS Support)
				if cfg.App.TLS.Enable {
					log.Info("TLS support was enabled.")

					if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
						log.Error("An unknown error occurred when to run server!", zap.Error(err))
					}
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error("An unknown error occurred when to run server!", zap.Error(err))
					}
				}()

				database.ConnectDatabase()

				// database.MigrateModels()
				// database.SeedModels(seeds.ArticleSeed)

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic("fiber.Shutdown()", zap.Error(err))
				}

				log.Info("Running cleanup tasks...")
				log.Info("1- Shutdown the database")
				database.ShutdownDatabase()
				log.Info("app was successful shutdown.", zap.String("app", cfg.App.Name))
				return nil
			},
		},
	)
}
