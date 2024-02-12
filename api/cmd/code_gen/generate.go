package main

import (
	"fmt"
	"gva/internal/code_gen"
	ustrings "gva/utils/strings"
	"os"
	"os/exec"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) > 0 {
		entity := ustrings.ToPascalCase(args[0])
		params := code_gen.CodeGenParams{
			Entity:           entity,
			EntityLower:      ustrings.PascalToCamel(entity),
			EntityAllLower:   strings.ReplaceAll(ustrings.PascalToSnake(entity), "_", ""),
			EntitySnake:      ustrings.PascalToSnake(entity),
			EntityUpperSnake: strings.ToUpper(ustrings.PascalToSnake(entity)),
			Table:            ustrings.PascalToSnake(entity) + "s",
		}

		code_gen.GenerateCodes(params)

		if _, err := exec.Command("go", "generate", "cmd/ent/generate.go").CombinedOutput(); err != nil {
			fmt.Printf("Failed to run command: %v\n", err)
			return
		}
	}
}

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target=../../internal/ent ../../app/database/schema --feature sql/versioned-migration
