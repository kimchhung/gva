package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	module_template "github.com/kimchhung/gva/extra/internal/codegen/module"
)

type CodeGenParams struct {
	EntityPascal     string
	EntityAllLower   string
	EntityCamel      string
	EntitySnake      string
	EntityUpperSnake string
	EntityKebab      string
	Table            string
}

func GenerateCodes(params CodeGenParams) {
	// GenerateCodeByTemplate(params, "schema", "app/database/schema", module_template.Schema)
	GenerateModule(params, "module", "api/admin/module", module_template.Module)
	GenerateModuleChild(params, "dto", "api/admin/module", "request", module_template.DtoRequest)
	GenerateModuleChild(params, "dto", "api/admin/module", "response", module_template.DtoResponse)
	GenerateModuleChild(params, "repository", "api/admin/module", "repository", module_template.Repository)
	GenerateModuleChild(params, "service", "api/admin/module", "service", module_template.Service)
	GenerateModuleChild(params, "controller", "api/admin/module", "controller", module_template.Controller)
	// Appends(params)
}

func Appends(params CodeGenParams) {
	InjectCodeToPos("main.go", map[string]string{
		"// #inject:module ":      fmt.Sprintf("%v.New%vModule,\n", params.EntitySnake, params.EntityPascal),
		"// #inject:moduleImport": fmt.Sprintf(`"github.com/kimchhung/gva/extra/api/admin/module/%v"`+"\n", params.EntitySnake),
	}, true)
}

func GenerateCodeByTemplate(params CodeGenParams, templateName string, directory string, templateContent string) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		panic(err)
	}

	fullPath := directory + "/" + params.EntitySnake + ".go"

	file := createFullPathFile(fullPath)
	defer file.Close()

	err = tmpl.Execute(file, params)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated " + fullPath)
}

func GenerateModule(params CodeGenParams, templateName string, directory string, templateContent string) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		panic(err)
	}

	fullPath := directory + "/" + params.EntitySnake + "/" + params.EntitySnake + "_module.go"

	file := createFullPathFile(fullPath)
	defer file.Close()

	err = tmpl.Execute(file, params)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated " + fullPath)
}

func GenerateModuleChild(params CodeGenParams, templateName, directory, suffix, templateContent string) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		panic(err)
	}

	fullPath := directory + "/" + params.EntitySnake + "/" + templateName + "/" + params.EntitySnake + "_" + suffix + ".go"

	file := createFullPathFile(fullPath)
	defer file.Close()

	err = tmpl.Execute(file, params)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated " + fullPath)
}

func createFullPathFile(fullPath string) *os.File {
	dirPath := filepath.Dir(fullPath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			panic(err)
		}
	}

	file, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	return file
}
