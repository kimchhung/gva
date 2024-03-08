package common

import (
	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/app/middleware"
	"github.com/kimchhung/gva/extra/internal/bootstrap"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"go.uber.org/fx"
)

var NewCommonModule = fx.Module("CommonModule",
	fx.Provide(bootstrap.NewLogger),
	fx.Provide(bootstrap.NewFiber),
	fx.Provide(database.NewDatabase),
	fx.Provide(middleware.NewMiddleware),

	// Services
	fx.Provide(services.NewPasswordService),
	fx.Provide(services.NewJwtService),
)
