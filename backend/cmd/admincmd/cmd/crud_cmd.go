package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gva/internal/codegen"
	"github.com/spf13/cobra"
)

func GenerateByCmd(params codegen.CodeGenParams, cmds ...string) {
	seperator := ","

	all := []string{}
	for _, text := range cmds {
		parts := strings.Split(text, seperator)
		all = append(all, parts...)
	}

	if len(all) == 0 {
		all = []string{
			"M", "R", "P", "D", "S", "SC",
		}
	}

	for _, cmd := range all {
		cmd = strings.ReplaceAll(cmd, "-", "")
		switch cmd {
		case "SC", "schema":
			codegen.GenerateFiles(params, "schema")
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
	Use:   "crud",
	Short: "generate crud template",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("name is required")
		}

		name := args[0]
		options := args[1:]

		entity := codegen.NewCodeGenParams(name)
		GenerateByCmd(entity, options...)

		// generate ent
		output, err := exec.Command("go", "generate", "cmd/ent/generate.go").CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run command: %v\n", err)
			return
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(crudCmd)
}
