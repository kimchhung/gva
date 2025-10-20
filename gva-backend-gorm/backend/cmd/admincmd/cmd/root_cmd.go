package cmd

import (
	"backend/env"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var makeEnvCMD = &cobra.Command{
	Use:   "env.create",
	Short: "Generate .env from env/config.go",

	Run: func(cmd *cobra.Command, args []string) {
		if err := env.GenerateEnvFromDefaultConfig(false); err != nil {
			panic(fmt.Errorf("ReadEnvOrGenerate: %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(makeEnvCMD)
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Root CLI",
	Long: `CLI is a CLI application for managing routes.
It provides functionalities to pull and push routes.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
