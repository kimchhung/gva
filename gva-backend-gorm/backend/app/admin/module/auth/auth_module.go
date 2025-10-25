package auth

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var AuthModule = fx.Module("AuthModule",
	fx.Provide(NewAuthService),

	// Regiser Controller
	adminrouter.Controller.Add(NewAuthController),
)
