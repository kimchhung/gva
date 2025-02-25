package blog

import (
	"backend/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var BlogModule = fx.Module("BlogModule",
	// Register Service
	fx.Provide(NewBlogService),

	// Regiser Controller
	controller.ProvideAdminController(NewBlogController),
)
