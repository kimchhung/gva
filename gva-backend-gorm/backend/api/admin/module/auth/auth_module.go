package auth

import (
	"backend/app/common/controller"

	"go.uber.org/fx"
)

// Register bulkly
var AuthModule = fx.Module("AuthModule",
	fx.Provide(NewAuthService),

	// Regiser Controller
	controller.ProvideAdminController(NewAuthController),
)
