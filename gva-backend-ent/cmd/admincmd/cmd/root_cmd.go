package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adminroute",
	Short: "AdminRoute CLI",
	Long: `AdminRoute is a CLI application for managing routes.
It provides functionalities to pull and push routes.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
