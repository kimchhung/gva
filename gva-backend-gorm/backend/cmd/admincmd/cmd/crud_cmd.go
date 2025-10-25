package cmd

import (
	"fmt"
	"strings"

	"backend/internal/codegen"

	"github.com/spf13/cobra"
)

func GenerateByCmd(params codegen.CodeGenParams, cmds ...string) {

	fmt.Println("----------- Generating:", params.EntityAllLower, "------------")
	fmt.Println("")
	defer fmt.Println("\n-----------Completed:", params.EntityAllLower, "------------")

	seperator := ","

	all := []string{}
	for _, text := range cmds {
		parts := strings.Split(text, seperator)
		for _, p := range parts {
			if p != "" {
				all = append(all, p)
			}
		}
	}

	isGenerateAll := len(all) == 0
	if isGenerateAll {
		all = []string{
			// common base
			"model", "repository", "permission",

			// module base
			"dto", "service", "controller", "module",
		}
	}

	for _, cmd := range all {
		cmd = strings.ReplaceAll(cmd, "-", "")
		switch cmd {
		case "MD", "model":
			codegen.GenerateFiles(params, "model")
		case "M", "module":
			codegen.GenerateFiles(params, "module")
		case "R", "repository":
			codegen.GenerateFiles(params, "repository")
		case "P", "permission":
			codegen.GenerateFiles(params, "permission")
		case "D", "dto":
			codegen.GenerateFiles(params, "dto")
		case "S", "service":
			codegen.GenerateFiles(params, "service")
		case "C", "controller":
			codegen.GenerateFiles(params, "controller")
		default:
			panic(fmt.Errorf("unknown option: %v", cmd))
		}
	}

}

var crudCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate crud template",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("name is required")
		}

		name := args[0]
		options := args[1:]

		entity := codegen.NewCodeGenParams(name)
		GenerateByCmd(entity, options...)

	},
}

func init() {
	rootCmd.AddCommand(crudCmd)
}
