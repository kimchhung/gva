package app

import (
	"fmt"
	"runtime"
	"time"

	share "backend/app/share"
	"backend/core/router"
	"backend/env"
	"backend/internal/bootstrap"
	"backend/internal/logger"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(cfg *env.Config, opts ...fx.Option) *fx.App {

	return fx.New(

		// Provide config
		fx.Supply(cfg),

		/* Common Module */
		share.NewShareModule(),

		/* add web or admin modules */
		fx.Options(opts...),

		// global router
		fx.Provide(router.NewRouter),

		// Start Application, execute on run
		bootstrap.NewModule(),
	)
}

func PrintMemUsage() {
	go func() {
		time.Sleep(time.Second)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		logger.G().Info("Memory",
			zap.String("alloc", fmt.Sprintf("%v MiB", bToMb(m.Alloc))),
			zap.String("totalAlloc", fmt.Sprintf("%v MiB", bToMb(m.TotalAlloc))),
			zap.String("sys", fmt.Sprintf("%v MiB", bToMb(m.Sys))),
			zap.String("NumGC", fmt.Sprintf("%v", m.NumGC)),
		)

	}()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
