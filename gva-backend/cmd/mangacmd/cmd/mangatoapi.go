package cmd

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/gva/internal/lang"
	"github.com/gva/internal/scrapper"
	"github.com/gva/internal/scrapper/mangatoapi"
	"github.com/gva/internal/validator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var manganatoCmd = &cobra.Command{
	Use:   "manganato.search",
	Short: "search manga CLI",
	Run: func(_ *cobra.Command, args []string) {
		name := strings.TrimSpace(args[0])
		if name == "" {
			panic("first argument is required")
		}

		for _, arg := range args {
			switch true {
			case strings.Contains(arg, "name"):
				name = strings.TrimPrefix(arg, "name=")
			}
		}

		ctx := context.Background()
		lang.InitializeTranslator()
		validator.InitializeValidator()
		collector := scrapper.NewCollector()

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		mangaApi := mangatoapi.NewManganatoApi(collector, log.Logger)

		t1 := time.Now()
		mangas, errs := mangaApi.SearchManga(ctx, name)

		log.Debug().
			Dur("latency_ms", time.Since(t1)).
			Errs("errs", errs).
			Any("mangas", mangas).
			Msg("mangaApi.SearchManga ended.")
	},
}

var manganatoDetailCmd = &cobra.Command{
	Use:   "manganato.detail",
	Short: "get manga detail CLI",
	Run: func(_ *cobra.Command, args []string) {
		id := strings.TrimSpace(args[0])
		if id == "" {
			panic("first argument is required")
		}

		for _, arg := range args {
			switch true {
			case strings.Contains(arg, "id"):
				id = strings.TrimPrefix(arg, "id=")
			}
		}

		ctx := context.Background()
		lang.InitializeTranslator()
		validator.InitializeValidator()
		collector := scrapper.NewCollector()

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		mangaApi := mangatoapi.NewManganatoApi(collector, log.Logger)

		t1 := time.Now()
		detail, errs := mangaApi.FetctMangaDetail(ctx, "https://chapmanganato.to", id)

		log.Debug().
			Dur("latency_ms", time.Since(t1)).
			Errs("errs", errs).
			Any("detail", detail).Msg("mangaApi.FetctMangaDetail ended.")
	},
}

func init() {
	rootCmd.AddCommand(manganatoCmd)
	rootCmd.AddCommand(manganatoDetailCmd)
}
