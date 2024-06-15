package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	codegen "github.com/kimchhung/gva/backend-echo/internal/codegen"
	ustrings "github.com/kimchhung/gva/backend-echo/utils/strings"
)

func main() {

	args := os.Args[1:]
	if len(args) > 0 {
		entity := ustrings.ToPascalCase(args[0])
		params := codegen.CodeGenParams{
			EntityPascal:     entity,
			EntityCamel:      ustrings.PascalToCamel(entity),
			EntityAllLower:   strings.ReplaceAll(ustrings.PascalToSnake(entity), "_", ""),
			EntitySnake:      ustrings.PascalToSnake(entity),
			EntityUpperSnake: strings.ToUpper(ustrings.PascalToSnake(entity)),
			EntityKebab:      strings.ReplaceAll(ustrings.PascalToSnake(entity), "_", "-"),
			Table:            ustrings.PascalToSnake(entity) + "s",
		}

		codegen.GenerateCodes(params)

		output, err := exec.Command("go", "generate", "cmd/ent/generate.go").CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run command: %v\n", err)
			return
		}
		fmt.Println(string(output))

		output, err = exec.Command("swag", "init").CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run command: %v\n", err)
			return
		}
		fmt.Println(string(output))

	}
}

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target=../../internal/ent ../../app/database/schema --feature sql/versioned-migration
