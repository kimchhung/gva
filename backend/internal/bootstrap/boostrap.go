package bootstrap

import (
	"context"
	"time"

	"github.com/gva/app/middleware"
	"github.com/gva/app/router"
	"github.com/gva/env"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/lang"
	"github.com/gva/utils/validator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type Bootstrap struct {
	lc          fx.Lifecycle
	cfg         *env.Config
	routers     *router.Router
	app         *echo.Echo
	middlewares *middleware.Middleware
	db          *database.Database
	log         *zerolog.Logger

	listeners      []chan struct{}
	isShuttingdown bool
}

func NewBootstrap(
	lc fx.Lifecycle,
	cfg *env.Config,
	routers *router.Router,
	app *echo.Echo,
	middlewares *middleware.Middleware,
	db *database.Database,
	log *zerolog.Logger,
) *Bootstrap {
	return &Bootstrap{
		lc:          lc,
		cfg:         cfg,
		routers:     routers,
		app:         app,
		middlewares: middlewares,
		db:          db,
		log:         log,
	}
}

func (b *Bootstrap) Done() <-chan struct{} {
	done := make(chan struct{})
	b.listeners = append(b.listeners, done)
	return done
}

func (b *Bootstrap) notifyDone() {
	for _, l := range b.listeners {
		l <- struct{}{}
		close(l)
	}
}

func (b *Bootstrap) setup() {
	b.lc.Append(
		fx.StartStopHook(
			b.start,
			b.stop,
		),
	)
}

func (b *Bootstrap) start(ctx context.Context) {

	// Initailize validator and translator
	if err := lang.InitializeTranslator(); err != nil {
		b.log.Panic().Err(err).Msg("failed to initialize translator!")
	}

	if err := validator.InitializeValidator(); err != nil {
		b.log.Panic().Err(err).Msg("failed to initialize validator!")
	}

	// Connect db
	if err := b.db.ConnectDatabase(); err != nil {
		b.log.Panic().Err(err).Msg("failed to connect to db!")
	}

	// Register middlewares & routes
	b.middlewares.Register()
	b.routers.Register(b.app, b.cfg)

	b.app.Server.RegisterOnShutdown(func() {
		log.Info().Msg("1- Shutdown the database")
		if err := b.db.ShutdownDatabase(); err != nil {
			log.Err(err).Msg("failed to shutdown db!")
		}
	})

	if b.cfg.App.PrintRoutes {
		printRoutes(b.app.Routes())
	}

	printStartupMessage(b.cfg)

	// notify bootstrap as done
	b.notifyDone()

	// Listen the app (with TLS Support)
	if b.cfg.App.TLS.Auto {
		log.Info().Msg("Auto TLS support was enabled.")

		go func() {
			if err := b.app.StartAutoTLS(b.cfg.App.Port); err != nil && !b.isShuttingdown {
				log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
			}
		}()
	} else if b.cfg.App.TLS.Enable {
		log.Info().Msg("TLS support was enabled.")

		go func() {
			if err := b.app.StartTLS(b.cfg.App.Port, b.cfg.App.TLS.CertFile, b.cfg.App.TLS.KeyFile); err != nil && !b.isShuttingdown {
				log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
			}
		}()
	} else {
		go func() {
			if err := b.app.Start(b.cfg.App.Port); err != nil && !b.isShuttingdown {
				log.Panic().Err(err).Msg("An unknown error occurred when to run server!")
			}
		}()
	}
}

func (b *Bootstrap) stop(ctx context.Context) {
	b.isShuttingdown = true
	ctx, cancel := context.WithTimeout(ctx, time.Duration(b.cfg.App.ShutdownTimeout)*time.Second)
	defer cancel()

	log.Info().Msg("Shutting down the app...")
	log.Info().Msg("Running cleanup tasks...")

	if err := b.app.Shutdown(ctx); err != nil {
		log.Panic().Err(err).Msg("")
	}

	log.Info().Msgf("%s was successful shutdown.", b.cfg.App.Name)
	log.Info().Msg("\u001b[96mSee you again👋\u001b[0m")
}
