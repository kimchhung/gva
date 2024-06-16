package cmd

import (
	"context"

	"github.com/kimchhung/gva/backend/env"
	"github.com/kimchhung/gva/backend/internal/bootstrap"
	"github.com/kimchhung/gva/backend/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend/internal/menu"
	"github.com/spf13/cobra"
)

var pullRouteCmd = &cobra.Command{
	Use:   "pull.route",
	Short: "Pull routes from the database",
	Long:  `This command pulls routes from the database and performs necessary operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := env.NewConfig()
		log := bootstrap.NewLogger(cfg)

		log.Info().Msg("Pulling routes...")
		defer log.Info().Msg("Pull routes is completed")

		db := database.NewDatabase(env.NewConfig(), log)
		db.ConnectDatabase()

		menu.PullRoutes(ctx, db.Client, routeDataPath)
	},
}

func init() {
	rootCmd.AddCommand(pullRouteCmd)
}
