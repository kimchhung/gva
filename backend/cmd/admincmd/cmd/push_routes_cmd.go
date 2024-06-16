package cmd

import (
	"context"

	"github.com/kimchhung/gva/backend/env"
	"github.com/kimchhung/gva/backend/internal/bootstrap"
	"github.com/kimchhung/gva/backend/internal/bootstrap/database"
	"github.com/kimchhung/gva/backend/internal/menu"
	"github.com/spf13/cobra"
)

const (
	routeDataPath = "./app/database/data/routes_data.json"
)

var pushRouteCmd = &cobra.Command{
	Use:   "push.route",
	Short: "push routes from json to database, delete and recreate base on file",
	Long:  `This command pulls routes from the database and performs necessary operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := env.NewConfig()
		log := bootstrap.NewLogger(cfg)

		log.Info().Msg("Pushing routes...")
		defer log.Info().Msg("Push routes is completed")

		db := database.NewDatabase(env.NewConfig(), log)
		db.ConnectDatabase()

		menu.PushRouters(ctx, db.Client, routeDataPath)
	},
}

func init() {
	rootCmd.AddCommand(pushRouteCmd)
}
