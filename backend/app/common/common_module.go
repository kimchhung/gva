package common

import (
	"github.com/gva/app/common/repository"
	"github.com/gva/app/common/service"
	"github.com/gva/app/middleware"
	"github.com/gva/internal/bootstrap"
	"github.com/gva/internal/bootstrap/database"

	"go.uber.org/fx"
)

var NewCommonModule = fx.Module("CommonModule",
	fx.Provide(bootstrap.NewLogger),
	fx.Provide(bootstrap.NewEcho),
	fx.Provide(database.NewDatabase),
	fx.Provide(middleware.NewMiddleware),

	// Services
	fx.Provide(service.NewPasswordService),
	fx.Provide(service.NewJwtService),

	// Database Repository | repositores
	fx.Provide(repository.NewAdminRepository),
	fx.Provide(repository.NewMenuRepository),
	fx.Provide(repository.NewPermissionRepository),
	fx.Provide(repository.NewDepartmentRepository),
	fx.Provide(repository.NewTodoRepository),
	// #inject:repository (do not remove this comment, it is used by the code generator)
)
