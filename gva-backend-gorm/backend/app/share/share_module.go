package share

import (
	"backend/app/share/middleware"
	"backend/app/share/repository"
	"backend/app/share/service"

	"go.uber.org/fx"
)

func NewShareModule() fx.Option {
	return fx.Module("ShareModule",
		fx.Provide(middleware.NewMiddleware),

		// share Services can use in every apps
		service.NewServiceModule(),

		// repositores
		repository.NewRepositoyModule(),
	)
}
