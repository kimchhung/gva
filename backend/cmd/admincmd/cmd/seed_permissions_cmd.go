package cmd

import (
	"context"

	"github.com/gva/app/common/permission"
	"github.com/gva/app/common/service"
	"github.com/gva/env"
	"github.com/gva/internal/bootstrap"
	"github.com/gva/internal/bootstrap/database"

	"github.com/spf13/cobra"
)

var seedPermissionCmd = &cobra.Command{
	Use:   "seed.permission",
	Short: "Seed permission from seeder permission",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := env.NewConfig()
		log := bootstrap.NewLogger(cfg)

		log.Info().Msg("Run seeding...")
		defer log.Info().Msg("Run seed is completed")

		db := database.NewDatabase(env.NewConfig(), log)
		db.ConnectDatabase()
		defer db.Close()

		// dependencies for seeding
		ctx = context.WithValue(ctx, env.Config{}, cfg)
		ctx = context.WithValue(ctx, service.PasswordService{}, service.NewPasswordService(cfg))

		// seeds all permission
		db.SeedModels(ctx, permission.AllSeeders()...)
	},
}

func init() {
	rootCmd.AddCommand(seedPermissionCmd)
}
