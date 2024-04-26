package cmd

import (
	"github.com/kimchhung/gva/extra/api"
	"github.com/kimchhung/gva/extra/api/admin"
	"github.com/kimchhung/gva/extra/api/web"
	"github.com/spf13/cobra"
)

var serveWebCmd = &cobra.Command{
	Use:   "web",
	Short: "server web api",
	Run: func(cmd *cobra.Command, args []string) {
		web.Run()
	},
}

var serveAdminCmd = &cobra.Command{
	Use:   "admin",
	Short: "serve admin api",
	Run: func(cmd *cobra.Command, args []string) {
		admin.Run()
	},
}

var serveAllCmd = &cobra.Command{
	Use:   "all",
	Short: "serve admin | web api",
	Run: func(cmd *cobra.Command, args []string) {
		api.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveAdminCmd)
	rootCmd.AddCommand(serveWebCmd)
	rootCmd.AddCommand(serveAllCmd)
}
