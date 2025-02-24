package middleware

import (
	"time"

	appctx "backend/app/common/context"
	apperror "backend/app/common/error"
	"backend/env"
	"backend/internal/bootstrap/database"
	"backend/internal/lang"
	"backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Middleware struct {
	app *echo.Echo
	cfg *env.Config
	log *zap.Logger
	db  *database.Database
}

func NewMiddleware(app *echo.Echo, cfg *env.Config, log *zap.Logger, db *database.Database) *Middleware {
	return &Middleware{
		app: app,
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (m *Middleware) Register() {
	mdCfg := m.cfg.Middleware

	m.app.Use(
		OperationLogger(m.db),
		appctx.Middleware(m.cfg),
	)

	m.app.Use(
		lang.Middleware(),
		middleware.RemoveTrailingSlash(),
	)

	// cors
	m.app.Use(middleware.CORS())

	// rate limitor
	m.app.Use(middleware.RateLimiterWithConfig(
		middleware.RateLimiterConfig{
			Skipper: utils.IsEnabled(mdCfg.Limiter.Enable),
			Store: middleware.NewRateLimiterMemoryStoreWithConfig(
				middleware.RateLimiterMemoryStoreConfig{
					Rate:      rate.Limit(mdCfg.Limiter.Max),
					Burst:     int(mdCfg.Limiter.Max + 30),
					ExpiresIn: time.Duration(mdCfg.Limiter.ExpirationSeconds) * time.Second,
				},
			),
			IdentifierExtractor: func(ctx echo.Context) (string, error) {
				id := ctx.RealIP()
				return id, nil
			},
			ErrorHandler: func(context echo.Context, err error) error {
				return apperror.NewError(apperror.ErrUnknownError, apperror.Join(err))
			},
			DenyHandler: func(context echo.Context, identifier string, err error) error {
				return apperror.ErrTooManyRetries
			},
		},
	))

	// compress
	m.app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   mdCfg.Compress.Level,
		Skipper: utils.IsEnabled(mdCfg.Compress.Enable),
	}))

	// monitor
	if mdCfg.Monitor.Enable {
		// m.app.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		// 	Skipper: func(c echo.Context) bool {
		// 		return strings.Contains(c.Path(), mdCfg.Monitor.Path)
		// 	},
		// })) // adds middleware to gather metrics
		// m.app.GET(mdCfg.Monitor.Path, echoprometheus.NewHandler()) // adds route to serve gathered metrics
	}

	m.app.Use(TraceDebug(m.cfg))
}
