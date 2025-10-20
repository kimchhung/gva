package bootstrap

import (
	"context"
	"time"

	"backend/app/common/permission"
	"backend/app/common/seeds"
	"backend/app/common/service"
	"backend/app/middleware"
	"backend/app/router"
	"backend/env"
	"backend/internal/bootstrap/database"
	"backend/internal/ctxutil"

	"backend/internal/logger"

	"github.com/labstack/echo/v4"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Bootstrap struct {
	lc          fx.Lifecycle
	cfg         *env.Config
	routers     *router.Router
	app         *echo.Echo
	middlewares *middleware.Middleware
	db          *database.Database
	redis       *database.Redis
	log         *zap.Logger

	startedListeners   []chan struct{}
	shutdownListerners []chan struct{}
	isShuttingdown     bool
}

func NewBootstrap(
	lc fx.Lifecycle,
	cfg *env.Config,
	routers *router.Router,
	app *echo.Echo,
	middlewares *middleware.Middleware,
	db *database.Database,
	redis *database.Redis,
	log *zap.Logger,
) *Bootstrap {
	return &Bootstrap{
		lc:          lc,
		cfg:         cfg,
		routers:     routers,
		app:         app,
		middlewares: middlewares,
		db:          db,
		log:         log,
		redis:       redis,
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
	if !b.cfg.IsProd() {
		logger.Log(b.cfg)
	}

	if err := b.db.Connect(); err != nil {
		b.log.Panic("b.db.Connect", zap.Error(err))
	}

	b.RunSeed(ctx)

	if err := b.redis.Connect(); err != nil {
		b.log.Panic("b.redis.Connect", zap.Error(err))
	}

	// Register middlewares & routes
	b.middlewares.Register()
	b.routers.Register(ctxutil.Add(ctx, b.app, b.cfg))

	b.app.Server.RegisterOnShutdown(func() {
		b.log.Info("1- Shutdown the database")
		if err := b.db.Close(); err != nil {
			b.log.Error("b.db.Close", zap.Error(err))
		}

		b.log.Info("2- Shutdown the redis")
		if err := b.redis.Close(); err != nil {
			b.log.Error("b.db.Close", zap.Error(err))
		}
	})

	if b.cfg.App.PrintRoutes {
		printRoutes(b.app.Routes())
	}

	printStartupMessage(b.cfg)
	time.Sleep(time.Second)

	// notify bootstrap as done
	b.notifyStarted()

	// Listen the app (with TLS Support)
	if b.cfg.App.TLS.Auto {
		b.log.Info("Auto TLS support was enabled.")

		go func() {
			if err := b.app.StartAutoTLS(b.cfg.App.Port); err != nil && !b.isShuttingdown {
				b.log.Panic("An unknown error occurred when to run server!", zap.Error(err))
			}
		}()
	} else if b.cfg.App.TLS.Enable {
		b.log.Info("TLS support was enabled.")

		go func() {
			if err := b.app.StartTLS(b.cfg.App.Port, b.cfg.App.TLS.CertFile, b.cfg.App.TLS.KeyFile); err != nil && !b.isShuttingdown {
				b.log.Panic("An unknown error occurred when to run server!", zap.Error(err))
			}
		}()
	} else {
		go func() {
			if err := b.app.Start(b.cfg.App.Port); err != nil && !b.isShuttingdown {
				b.log.Panic("An unknown error occurred when to run server!", zap.Error(err))
			}
		}()
	}
}

func (b *Bootstrap) stop(ctx context.Context) {
	b.isShuttingdown = true
	ctx, cancel := context.WithTimeout(ctx, time.Duration(b.cfg.App.ShutdownTimeout)*time.Second)
	defer cancel()

	b.log.Info("Shutting down the app...")
	b.log.Info("Running cleanup tasks...")

	b.notifyShuttingDown()

	if err := b.app.Shutdown(ctx); err != nil {
		b.log.Panic("b.app.Shutdown", zap.Error(err))
	}

	b.log.Sugar().Infof("%s was successful shutdown.", b.cfg.App.Name)
	b.log.Info("\u001b[96mSee you again👋\u001b[0m")
}

func (b *Bootstrap) RunSeed(ctx context.Context) {
	if !b.cfg.DB.Seed.Enable {
		return
	}

	// dependencies for seeding
	ctx = ctxutil.Add(ctx, b.cfg, service.NewPasswordService(b.cfg))
	seeders := append(seeds.AllSeeders(), permission.AllSeeders()...)
	b.db.SeedModels(ctx, seeders...)
}
