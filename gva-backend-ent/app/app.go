package app

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gva/app/common"
	"github.com/gva/env"
	"github.com/gva/internal/bootstrap"

	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func New(cfg *env.Config, opts ...fx.Option) *fx.App {
	return fx.New(
		// boostrap

		// Provide config
		fx.Supply(cfg),

		/* Common Module */
		common.NewCommonModule,

		// Define logger
		fx.WithLogger(fxzerolog.InitPtr()),

		/* add web or admin modules */
		fx.Module("api", opts...),

		// Start Application, execute on run
		bootstrap.Module,
	)
}

func PrintMemUsage() {
	go func() {
		time.Sleep(time.Second)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		log.Info().
			Str("alloc", fmt.Sprintf("%v MiB", bToMb(m.Alloc))).
			Str("totalAlloc", fmt.Sprintf("%v MiB", bToMb(m.TotalAlloc))).
			Str("sys", fmt.Sprintf("%v MiB", bToMb(m.Sys))).
			Str("NumGC", fmt.Sprintf("%v", m.NumGC)).
			Msg("Memory")
	}()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
