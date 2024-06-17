package comic

import (
	"go.uber.org/fx"
)

// Register bulkly
var ComicModuleModule = fx.Module("ComicModule",
	fx.Provide(NewComicService),
)
