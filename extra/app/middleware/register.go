package middleware

import (
	"strings"
	"time"

	appctx "github.com/kimchhung/gva/extra/app/common/context"
	apperror "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/lang"
	"github.com/kimchhung/gva/extra/utils"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"github.com/rs/zerolog"
)

type Middleware struct {
	app *echo.Echo
	cfg *config.Config
	log *zerolog.Logger
}

func NewMiddleware(app *echo.Echo, cfg *config.Config, log *zerolog.Logger) *Middleware {
	return &Middleware{
		app: app,
		cfg: cfg,
		log: log,
	}
}

func (m *Middleware) Register() {
	mdCfg := m.cfg.Middleware

	if mdCfg.Filesystem.Enable {
		m.app.Static("/public/", "storage/public")
	}

	m.app.Pre(middleware.RemoveTrailingSlash())

	// language and recover error handling
	m.app.Pre(
		appctx.NewRequestContext(),
		lang.Middleware(),
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
					ExpiresIn: time.Duration(mdCfg.Limiter.ExpSecs) * time.Second,
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
		Level: int(mdCfg.Compress.Level),
		Skipper: func(c echo.Context) bool {
			if !mdCfg.Compress.Enable {
				return false
			}

			return strings.Contains(c.Path(), "/doc")
		},
	}))

	// monitor
	if mdCfg.Monitor.Enable {
		m.app.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
			Skipper: func(c echo.Context) bool {

				return strings.Contains(c.Path(), "/doc") || strings.Contains(c.Path(), mdCfg.Monitor.Path)
			},
		})) // adds middleware to gather metrics
		m.app.GET(mdCfg.Monitor.Path, echoprometheus.NewHandler()) // adds route to serve gathered metrics
	}

}
