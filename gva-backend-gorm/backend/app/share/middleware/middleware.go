package middleware

import (
	"time"

	apperror "backend/app/share/error"
	"backend/core/database"
	"backend/core/env"
	"backend/core/lang"
	coretype "backend/core/type"
	"backend/core/utils"
	"backend/core/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

var _ interface{ coretype.Middleware } = (*Middleware)(nil)

type Middleware struct {
	app        *echo.Echo
	cfg        *env.Config
	log        *zap.Logger
	db         *database.Database
	translator *lang.Translator
	validator  *validator.Validator
}

func NewMiddleware(
	app *echo.Echo,
	cfg *env.Config,
	log *zap.Logger,
	db *database.Database,
	translator *lang.Translator,
	validator *validator.Validator,

) *Middleware {
	return &Middleware{
		app:        app,
		cfg:        cfg,
		log:        log,
		db:         db,
		translator: translator,
		validator:  validator,
	}
}

func (m *Middleware) RegisterMiddleware(mr coretype.MiddlewareRouter) {
	mdCfg := m.cfg.Middleware

	mr.Use(
		middleware.RemoveTrailingSlash(),
		m.CORS(),
		m.Compress(),
	)

	mr.Use(
		m.RequestContext(),
		m.ResponseHook(),
		middleware.Recover(),
		m.RateLimit(),
		m.Translation(),
		// m.Debug(),
	)

	// monitor
	if mdCfg.Monitor.Enable {
		// mr.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		// 	Skipper: func(c echo.Context) bool {
		// 		return strings.Contains(c.Path(), mdCfg.Monitor.Path)
		// 	},
		// })) // adds middleware to gather metrics
		// mr.GET(mdCfg.Monitor.Path, echoprometheus.NewHandler()) // adds route to serve gathered metrics
	}
}

func (m *Middleware) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.DefaultCORSConfig)
}

func (m *Middleware) Compress() echo.MiddlewareFunc {
	return middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   m.cfg.Middleware.Compress.Level,
		Skipper: utils.IsEnabled(m.cfg.Middleware.Compress.Enable),
	})
}

func (m *Middleware) RateLimit() echo.MiddlewareFunc {
	cfg := m.cfg.Middleware

	return middleware.RateLimiterWithConfig(
		middleware.RateLimiterConfig{
			Skipper: utils.IsEnabled(cfg.Limiter.Enable),
			Store: middleware.NewRateLimiterMemoryStoreWithConfig(
				middleware.RateLimiterMemoryStoreConfig{
					Rate:      rate.Limit(cfg.Limiter.Max),
					Burst:     int(cfg.Limiter.Max + 30),
					ExpiresIn: time.Duration(cfg.Limiter.ExpirationTTL) * time.Second,
				},
			),
			IdentifierExtractor: func(ctx echo.Context) (string, error) {
				id := ctx.RealIP()
				return id, nil
			},
			DenyHandler: func(context echo.Context, identifier string, err error) error {
				return apperror.ErrTooManyRetries
			},
		},
	)
}
