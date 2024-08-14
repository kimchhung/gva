package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	module_template "github.com/gva/internal/codegen/module"
	ustrings "github.com/gva/utils/strings"
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

func NewCodeGenParams(name string) CodeGenParams {
	params := CodeGenParams{
		EntityPascal:     ustrings.ToPascalCase(name),
		EntityCamel:      ustrings.PascalToCamel(name),
		EntityAllLower:   strings.ReplaceAll(ustrings.PascalToSnake(name), "_", ""),
		EntitySnake:      ustrings.PascalToSnake(name),
		EntityUpperSnake: strings.ToUpper(ustrings.PascalToSnake(name)),
		EntityKebab:      strings.ReplaceAll(ustrings.PascalToSnake(name), "_", "-"),
		Table:            ustrings.PascalToSnake(name) + "s",
	}
	return params
}

func GenerateFiles(params CodeGenParams, opts ...string) {
	for _, opt := range opts {
		switch opt {
		case "schema":
			GenerateCodeByTemplate(params, "schema", "app/database/schema", module_template.Schema)
		case "module":
			GenerateModule(params, "module", "api/admin/module", module_template.Module)
			InjectCodeToPos("api/admin/module/module.go", map[string]string{
				"// #inject:module":       fmt.Sprintf("%v.%vModule,\n", params.EntityAllLower, params.EntityPascal),
				"// #inject:moduleImport": fmt.Sprintf(`"github.com/gva/api/admin/module/%v"`+"\n", params.EntityAllLower),
			}, true)
		case "repository":
			GenerateModuleChildNoFolder(params, "repository", "app/common", "repository", module_template.Repository)
			InjectCodeToPos("app/common/common_module.go", map[string]string{
				"// #inject:repository": fmt.Sprintf("fx.Provide(repository.New%vRepository),\n", params.EntityPascal),
			}, true)
		case "permission":
			GenerateModuleChildNoFolder(params, "permission", "app/common", "permission", module_template.Permission)
		case "dto":
			GenerateModuleChild(params, "dto", "api/admin/module", "request", module_template.DtoRequest)
			GenerateModuleChild(params, "dto", "api/admin/module", "response", module_template.DtoResponse)
		case "service":
			GenerateModuleChild(params, "", "api/admin/module", "service", module_template.Service)
		case "controller":
			GenerateModuleChild(params, "", "api/admin/module", "controller", module_template.Controller)
		default:
			panic(fmt.Errorf("unknown option: %v", opt))
		}
	}
}

func GenerateCodeByTemplate(params CodeGenParams, templateName string, directory string, templateContent string) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		panic(err)
	}

	fullPath := directory + "/" + params.EntityAllLower + ".go"

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

	fullPath := directory + "/" + params.EntityAllLower + "/" + params.EntityAllLower + "_module.go"

	file := createFullPathFile(fullPath)
	defer file.Close()

	err = tmpl.Execute(file, params)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated " + fullPath)
}

func GenerateModuleChildNoFolder(params CodeGenParams, templateName, directory, suffix, templateContent string) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		panic(err)
	}

	fullPath := directory + "/" + templateName + "/" + params.EntityAllLower + "_" + suffix + ".go"

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

	fullPath := directory + "/" + params.EntityAllLower + "/" + templateName + "/" + params.EntityAllLower + "_" + suffix + ".go"

	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		fmt.Printf("Skipping generation for existing file: %s\n", fullPath)
		return
	}

	file := createFullPathFile(fullPath)
	defer file.Close()

	if err = tmpl.Execute(file, params); err != nil {
		panic(err)
	}

	fmt.Println("Generated " + fullPath)
}

func createFullPathFile(fullPath string) *os.File {
	dirPath := filepath.Dir(fullPath)
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
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
