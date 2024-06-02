package common

import (
	"github.com/kimchhung/gva/extra/app/common/repository"
	"github.com/kimchhung/gva/extra/app/common/service"
	"github.com/kimchhung/gva/extra/app/middleware"
	"github.com/kimchhung/gva/extra/internal/bootstrap"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
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
	fx.Provide(repository.NewRouteRepository),
	fx.Provide(repository.NewPermissionRepository),
)
