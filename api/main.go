package main

import (
	"go.uber.org/fx"

	"gva/app/middleware"
	"gva/app/module/admin"
	"gva/app/module/article"
	"gva/config"
	"gva/internal/bootstrap"
	"gva/internal/bootstrap/database"

	fxzerolog "github.com/efectn/fx-zerolog"
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

		article.NewArticleModule,
		admin.NewAdminModule,
		// #inject:module (do not remove this comment, it is used by the code generator)

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
