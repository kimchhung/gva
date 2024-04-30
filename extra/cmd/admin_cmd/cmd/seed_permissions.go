package cmd

import (
	"context"

	"github.com/kimchhung/gva/extra/app/common/permissions"
	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/spf13/cobra"
)

var seedPermissionCmd = &cobra.Command{
	Use:   "seed.permission",
	Short: "Seed permission from seeder permission",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := config.NewConfig()
		log := bootstrap.NewLogger(cfg)

		log.Info().Msg("Run seeding...")
		defer log.Info().Msg("Run seed is completed")

		db := database.NewDatabase(config.NewConfig(), log)
		db.ConnectDatabase()
		defer db.Close()

		// dependencies for seeding
		ctx = context.WithValue(ctx, config.Config{}, cfg)
		ctx = context.WithValue(ctx, services.PasswordService{}, services.NewPasswordService(cfg))

		db.SeedModels(ctx,
			permissions.AdminPermissionSeeder{},
			permissions.AdminRolePermissionSeeder{},
		)
	},
}

func init() {
	rootCmd.AddCommand(seedPermissionCmd)
}
