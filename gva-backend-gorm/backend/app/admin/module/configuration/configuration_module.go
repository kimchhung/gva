package configuration

import (
	adminrouter "backend/app/admin/router"

	"go.uber.org/fx"
)

// Register bulkly
var ConfigurationModule = fx.Module("ConfigurationModule",
	// Register Service
	fx.Provide(NewConfigurationService),

	// Regiser Controller
	adminrouter.Controller.Add(NewConfigurationController),
)
