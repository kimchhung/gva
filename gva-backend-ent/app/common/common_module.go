package common

import (
	"github.com/gva/app/common/repository"
	"github.com/gva/app/common/service"
	"github.com/gva/app/middleware"

	"go.uber.org/fx"
)

var NewCommonModule = fx.Module("CommonModule",
	fx.Provide(middleware.NewMiddleware),

	// Services
	fx.Provide(service.NewPubsubService),
	fx.Provide(service.NewPasswordService),
	fx.Provide(service.NewJwtService),
	fx.Provide(service.NewS3Service),
	// Database Repository | repositores
	fx.Provide(repository.NewAdminRepository),
	fx.Provide(repository.NewPermissionRepository),
	fx.Provide(repository.NewDepartmentRepository),
	// #inject:repository (do not remove this comment, it is used by the code generator)
)
