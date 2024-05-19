package auth

import (
	"github.com/kimchhung/gva/extra/app/constant"
	"go.uber.org/fx"
)

// Register bulkly
var AuthModule = fx.Module("AuthModule",
	fx.Provide(NewAuthService),

	// Regiser Controller
	constant.ProvideAdminController(NewAuthController),
)
