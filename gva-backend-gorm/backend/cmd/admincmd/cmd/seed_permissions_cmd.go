package cmd

import (
	"context"

	"backend/app/common/permission"
	"backend/app/common/service"
	"backend/env"
	"backend/internal/bootstrap"
	"backend/internal/bootstrap/database"

	"github.com/spf13/cobra"
)

var seedPermissionCmd = &cobra.Command{
	Use:   "seed.permission",
	Short: "Seed permission from seeder permission",

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

		// seeds all permission
		db.SeedModels(ctx, permission.AllSeeders()...)
	},
}

func init() {
	rootCmd.AddCommand(seedPermissionCmd)
}
