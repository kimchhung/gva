package bootstrap

import (
	"backend/internal/bootstrap/database"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("bootstrap",
		// Define logger
		fx.Provide(NewLogger),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Provide(NewEcho),
		fx.Provide(database.NewDatabase),
		fx.Provide(database.NewRedis),
		fx.Provide(NewBootstrap),
		fx.Invoke(
			func(b *Bootstrap) {
				b.setup()
			},
		),
	)
}
