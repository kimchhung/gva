package middleware

import (
	"net/http"
	"time"

	"github.com/kimchhung/gva/extra/app/common/contexts"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/lang"

	"github.com/kimchhung/gva/extra/utils"

	"github.com/rs/zerolog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

type Middleware struct {
	app *fiber.App
	cfg *config.Config
	log *zerolog.Logger
}

func NewMiddleware(app *fiber.App, cfg *config.Config, log *zerolog.Logger) *Middleware {
	return &Middleware{
		app: app,
		cfg: cfg,
		log: log,
	}
}

func (m *Middleware) Register() {

	m.app.Use(cors.New())

	m.app.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(m.cfg.Middleware.Limiter.Enable),
		Max:        m.cfg.Middleware.Limiter.Max,
		Expiration: m.cfg.Middleware.Limiter.ExpSecs * time.Second,
	}))

	// add locale to context
	m.app.Use(
		contexts.NewRequestContext(),
		lang.Register("locale"),
	)

	m.app.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(m.cfg.Middleware.Compress.Enable),
		Level: m.cfg.Middleware.Compress.Level,
	}))

	// m.app.Use(recover.New(recover.Config{
	// 	Next:              utils.IsEnabled(m.cfg.Middleware.Recover.Enable),
	// 	EnableStackTrace:  true,
	// 	StackTraceHandler: request.StackHandler,
	// }))

	m.app.Use(pprof.New(pprof.Config{
		Next: utils.IsEnabled(m.cfg.Middleware.Pprof.Enable),
	}))

	m.app.Use(filesystem.New(filesystem.Config{
		Next:   utils.IsEnabled(m.cfg.Middleware.Filesystem.Enable),
		Root:   http.Dir(m.cfg.Middleware.Filesystem.Root),
		Browse: m.cfg.Middleware.Filesystem.Browse,
		MaxAge: m.cfg.Middleware.Filesystem.MaxAge,
	}))

	m.app.Get(m.cfg.Middleware.Monitor.Path, monitor.New(monitor.Config{
		Next: utils.IsEnabled(m.cfg.Middleware.Monitor.Enable),
	}))

}
