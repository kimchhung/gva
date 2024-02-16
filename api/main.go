package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/kimchhung/gva/app/middleware"
	"github.com/kimchhung/gva/app/module"

	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/bootstrap"
	"github.com/kimchhung/gva/internal/bootstrap/database"
	"github.com/kimchhung/gva/internal/control_route"

	"github.com/kimchhung/gva/app/module/admin"

	"github.com/kimchhung/gva/app/module/permission"
	"github.com/kimchhung/gva/app/module/role"
	// #inject:moduleImport (do not remove this comment, it is used by the code generator)
)

func main() {
	fx.New(
		// Provide patterns
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewLogger),
		fx.Provide(bootstrap.NewFiber),
		fx.Provide(database.NewDatabase),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(
			fx.Annotate(module.NewRouter,
				fx.As(new(control_route.Router)),
				fx.ParamTags(`group:"routers"`),
			),
		),

		admin.NewAdminModule,
		permission.NewPermissionModule,
		role.NewRoleModule,
		// #inject:module (do not remove this comment, it is used by the code generator)

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger((func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		})),
	).Run()
}
