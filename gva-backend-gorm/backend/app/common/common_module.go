package common

import (
	"backend/app/common/repository"
	"backend/app/common/service"
	"backend/app/middleware"

	"go.uber.org/fx"
)

var NewCommonModule = fx.Module("CommonModule",
	fx.Provide(middleware.NewMiddleware),

	// Common Services
	service.NewServiceModule(),

	// repositores
	repository.NewRepositoyModule(),
)
