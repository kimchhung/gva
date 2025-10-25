package adminrouter

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	docs "backend/app/admin/docs"
	adminmiddleware "backend/app/admin/middleware"
	"backend/core/lang"
	"backend/core/router"
	coretype "backend/core/type"
	"backend/core/utils"
	"backend/core/utils/swagger"
	"backend/env"
	"backend/internal/ctr"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	_ interface{ coretype.AppRouter } = (*Router)(nil)
)

const (
	Controller router.Group = "admin-controllers"
)

type RouterParam struct {
	fx.In
	App *echo.Echo
	Env *env.Config

	Translator      *lang.Translator
	AdminMiddleware *adminmiddleware.Middleware
	Controllers     []ctr.CTR `group:"admin-controllers"`
}

type Router struct {
	app *echo.Echo
	env *env.Config

	translator      *lang.Translator
	adminMiddleware *adminmiddleware.Middleware
	controllers     []ctr.CTR
}

func NewRouter(p RouterParam) *Router {
	return &Router{
		controllers:     p.Controllers,
		app:             p.App,
		env:             p.Env,
		translator:      p.Translator,
		adminMiddleware: p.AdminMiddleware,
	}
}

func (r *Router) Register(ctx context.Context) {
	//default value if not exist in env config
	utils.SetIfEmpty(&r.env.Admin.BasePath, "/admin/v1")
	docs.SwaggerInfoadmin.BasePath = r.env.Admin.BasePath
	docs.SwaggerInfoadmin.Host = r.env.App.Host

	swagger.Register(r.app, r.env.Admin.BasePath, r.env.Middleware.Swagger.Path,
		echoSwagger.InstanceName(docs.SwaggerInfoadmin.InstanceName()),
		echoSwagger.PersistAuthorization(true),
		echoSwagger.SyntaxHighlight(true),
	)

	api := r.app.Group(r.env.Admin.BasePath)

	// register admin middleware
	r.adminMiddleware.RegisterMiddleware(api)
	r.translator.Import("./app/admin/locale")

	if err := router.RegisterEcho(api, r.controllers); err != nil {
		panic(err)
	}
}
