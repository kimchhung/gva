package cmd

import (
	"fmt"
	"os"

	"github.com/kimchhung/gva/extra/api"
	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "all",
	Short: "Run both web and admin api",
	Run: func(cmd *cobra.Command, args []string) {
		api.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
