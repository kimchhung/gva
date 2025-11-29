package bootstrap

import (
	"backend/core/database"
	"backend/core/lang"
	"backend/core/validator"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("bootstrap",
		// Define logger
		fx.Provide(NewLogger),
		fx.Provide(func(log *Logger) *zap.Logger {
			log.Initailized()
			return log.logger
		}),

		fx.Provide(lang.NewTranslator),
		fx.Provide(validator.NewValidator),
		fx.Provide(database.NewDatabase),
		fx.Provide(database.NewRedis),
		fx.Provide(NewEcho),
		fx.Provide(NewBootstrap),
		fx.Invoke(
			func(b *Bootstrap) {
				b.setup()
			},
		),
	)
}
