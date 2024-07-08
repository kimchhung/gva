package cmd

import (
	"context"

	"github.com/gva/env"
	"github.com/gva/internal/bootstrap"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/menu"

	"github.com/spf13/cobra"
)

const (
	routeDataPath = "./app/database/data/menu_data.json"
)

var pushRouteCmd = &cobra.Command{
	Use:   "push.Menu",
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
