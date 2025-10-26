package cmd

import (
	"context"

	"backend/app/admin/permission"
	"backend/app/share/service"
	"backend/core/database"
	"backend/core/env"
	"backend/internal/logger"

	"github.com/spf13/cobra"
)

var seedPermissionCmd = &cobra.Command{
	Use:   "seed.permission",
	Short: "Seed permission from seeder permission",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := env.NewConfig()
		db := database.NewDatabase(cfg, logger.G())
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
