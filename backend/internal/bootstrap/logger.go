package bootstrap

import (
	"os"

	"github.com/kimchhung/gva/backend/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(cfg *env.Config) *zerolog.Logger {
	zerolog.TimeFieldFormat = cfg.Logger.TimeFormat

	if cfg.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(cfg.Logger.Level)
	return &log.Logger
}
