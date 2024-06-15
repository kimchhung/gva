package bootstrap

import (
	"os"

	"github.com/kimchhung/gva/backend/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(cfg *config.Config) *zerolog.Logger {
	zerolog.TimeFieldFormat = cfg.Logger.TimeFormat

	if cfg.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(cfg.Logger.Level)
	return &log.Logger
}
