package middleware

import (
	"strings"
	"time"

	appctx "github.com/gva/app/common/context"
	apperror "github.com/gva/app/common/error"
	"github.com/gva/env"
	"github.com/gva/internal/lang"
	"github.com/gva/internal/utils"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type Middleware struct {
	app *echo.Echo
	cfg *env.Config
	log *zerolog.Logger
}

func NewMiddleware(app *echo.Echo, cfg *env.Config, log *zerolog.Logger) *Middleware {
	return &Middleware{
		app: app,
		cfg: cfg,
		log: log,
	}
}

func (m *Middleware) Register() {
	mdCfg := m.cfg.Middleware

	m.app.Pre(
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
		m.app.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
			Namespace: "gva",
			Skipper: func(c echo.Context) bool {
				return strings.Contains(c.Path(), mdCfg.Monitor.Path)
			},
		})) // adds middleware to gather metrics
		m.app.GET(mdCfg.Monitor.Path, echoprometheus.NewHandler()) // adds route to serve gathered metrics
	}

	m.app.Use(TraceDebug(m.cfg))
}
