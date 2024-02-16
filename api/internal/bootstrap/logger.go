package bootstrap

import (
	"github.com/kimchhung/gva/config"
	zPretier "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	// Create a new zap logger based on the environment (development or production)
	var config zap.Config
	if cfg.Logger.Prettier {
		config = zPretier.NewConfig()
		// slow 3000%
	} else {
		config = zap.NewProductionConfig()
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)

	// Set the log level
	logger = logger.WithOptions(
		zap.IncreaseLevel(zapcore.Level(cfg.Logger.Level)),
	)

	// Replace the default time encoder with a custom one if needed
	// logger = logger.WithOptions(zap.AddCallerSkip(1))

	return logger, nil
}

// Prefork hook for zerolog
// type PreforkHook struct{}

// func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
// 	if fiber.IsChild() {
// 		e.Discard()
// 	}
// }
