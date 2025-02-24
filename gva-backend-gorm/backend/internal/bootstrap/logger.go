package bootstrap

import (
	"backend/env"
	"backend/internal/logger"

	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger(env *env.Config) *zap.Logger {
	var config zap.Config
	if env.IsProd() {
		config = zap.NewProductionConfig()
	} else {
		if env.Logger.Prettier {
			config = prettyconsole.NewConfig()
		} else {
			config = zap.NewDevelopmentConfig()
		}
	}

	config.Level = zap.NewAtomicLevelAt(zapcore.Level(env.Logger.Level))

	log, _ := config.Build()
	log = zap.New(
		log.Core(),
		zap.AddStacktrace(zap.ErrorLevel),
		logger.WithCore(env),
	)
	zap.ReplaceGlobals(log)
	return log
}

func NewLogger(env *env.Config) *zap.Logger {
	initLogger(env)
	return logger.G()
}
