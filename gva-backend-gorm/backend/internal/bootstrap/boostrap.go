package bootstrap

import (
	"context"
	"sort"
	"strings"
	"time"

	"backend/app/middleware"
	"backend/app/router"
	"backend/env"
	"backend/internal/bootstrap/database"
	"backend/internal/lang"
	"backend/internal/logger"
	"backend/internal/treeprint"
	"backend/utils/color"
	"backend/utils/validator"

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

// notify when server started
func (b *Bootstrap) Started() (wait, done func()) {
	doneCh := make(chan struct{})
	b.startedListeners = append(b.startedListeners, doneCh)

	wait = func() {
		<-doneCh
	}

	done = func() {
		doneCh <- struct{}{}
	}

	return wait, done
}

// notify when server is shuting down
func (b *Bootstrap) ShuttingDown() (wait, done func()) {
	doneCh := make(chan struct{})
	b.startedListeners = append(b.shutdownListerners, doneCh)

	wait = func() {
		<-doneCh
	}

	done = func() {
		doneCh <- struct{}{}
	}

	return wait, done
}

func (b *Bootstrap) notifyShuttingDown() {
	resps := make([]chan struct{}, len(b.shutdownListerners))
	for i, req := range b.shutdownListerners {
		// notify shutdown
		req <- struct{}{}
		resps[i] = req
	}

	for _, resp := range resps {
		<-resp
		close(resp)
	}
}

func (b *Bootstrap) notifyStarted() {
	for _, req := range b.startedListeners {
		// notify server started
		req <- struct{}{}

		// wait process is done
		<-req
		close(req)
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

	// Initailize validator and translator
	if err := lang.InitializeTranslator(); err != nil {
		b.log.Panic("InitializeTranslator", zap.Error(err))
	}

	if err := validator.InitializeValidator(); err != nil {
		b.log.Panic("InitializeValidator", zap.Error(err))
	}

	if err := b.db.Connect(); err != nil {
		b.log.Panic("b.db.Connect", zap.Error(err))
	}

	if err := b.redis.Connect(); err != nil {
		b.log.Panic("b.redis.Connect", zap.Error(err))
	}

	// Register middlewares & routes
	b.middlewares.Register()
	b.routers.Register(b.app, b.cfg)

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

func printRoutes(routes []*echo.Route) {
	tree := treeprint.New("api")
	N := 4

	sort.Slice(routes, func(i, j int) bool {
		return len(strings.Split(routes[i].Path, "")) > len(strings.Split(routes[j].Path, ""))
	})

	maxLenth := calculateMaxLength(routes, N)
	for _, r := range routes {
		if r.Method == "echo_route_not_found" {
			continue
		}

		paths := []any{}
		for _, str := range strings.SplitAfterN(r.Path, "/", N) {
			str := strings.TrimSuffix(str, "/")
			if str == "" {
				continue
			}
			paths = append(paths, strings.TrimSuffix(str, "/"))
		}
		if len(paths) > N-2 {
			paths[N-2] = strings.ReplaceAll(strings.Split(paths[N-2].(string), "/")[0], "/", "")
		}
		httpPath := color.MethodColor(r.Method) + " " + r.Path
		space := calculateDynamicSpace(httpPath, maxLenth)
		paths = append(paths, httpPath+space+color.Cyan(r.Name))
		tree.AddPath(paths...)
	}

	treeprint.Print(tree)
}

func calculateMaxLength(routes []*echo.Route, N int) int {
	maxLength := 0

	for _, r := range routes {
		paths := []any{}
		for _, str := range strings.SplitAfterN(r.Path, "/", N) {
			paths = append(paths, strings.TrimSuffix(str, "/"))
		}

		length := 0
		for _, str := range paths {
			strs := strings.ReplaceAll(str.(string), "/", "")
			length += len(strs)
		}

		length += len(color.MethodColor(r.Method))
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func calculateDynamicSpace(path string, maxLength int) string {
	spaceNeeded := maxLength - len(path)
	if spaceNeeded <= 0 {
		spaceNeeded = 1
	}
	return strings.Repeat(" ", spaceNeeded)
}
