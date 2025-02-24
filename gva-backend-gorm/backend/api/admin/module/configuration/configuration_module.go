package configuration

import (
	"backend/app/common/controller"

	"go.uber.org/fx"
)

// Register bulkly
var ConfigurationModule = fx.Module("ConfigurationModule",
	// Register Service
	fx.Provide(NewConfigurationService),

	// Regiser Controller
	controller.ProvideAdminController(NewConfigurationController),
)
