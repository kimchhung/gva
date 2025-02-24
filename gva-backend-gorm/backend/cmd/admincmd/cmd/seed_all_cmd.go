package cmd

import (
	"context"

	"backend/app/common/permission"
	"backend/app/common/seeds"
	"backend/app/common/service"

	"backend/env"
	"backend/internal/bootstrap"
	"backend/internal/bootstrap/database"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed.all",
	Short: "Run all seeds",
	Long:  `This command pulls routes from the database and performs necessary operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := env.NewConfig()
		log := bootstrap.NewLogger(cfg)
		db := database.NewDatabase(cfg, log)
		db.Connect()
		defer db.Close()

		// dependencies for seeding
		ctx = context.WithValue(ctx, env.Config{}, cfg)
		ctx = context.WithValue(ctx, service.PasswordService{}, service.NewPasswordService(cfg))

		seeders := append(seeds.AllSeeders(), permission.AllSeeders()...)
		db.SeedModels(ctx, seeders...)
		log.Info("Run seed is completed")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
