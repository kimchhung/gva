package cmd

import (
	"context"

	"github.com/kimchhung/gva/extra/app/common/service"
	"github.com/kimchhung/gva/extra/app/database/seeds"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seeds",
	Short: "Run all seeds",
	Long:  `This command pulls routes from the database and performs necessary operations.`,
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
		ctx = context.WithValue(ctx, service.PasswordService{}, service.NewPasswordService(cfg))

		db.SeedModels(ctx,
			seeds.RouterSeeder{},
			seeds.SuperAdminSeeder{},
		)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
