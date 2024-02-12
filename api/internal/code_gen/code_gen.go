package code_gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type CodeGenParams struct {
	Entity           string
	EntityAllLower   string
	EntityLower      string
	EntitySnake      string
	EntityUpperSnake string
	Table            string
}

func GenerateCodes(params CodeGenParams) {
	GenerateCodeByTemplate(params, "schema", "app/database/schema", schemaTemplate)
	GenerateModuleByTemplate(params, "module", "app/module", moduleTemplate)
	GenerateModuleInByTemplate(params, "request", "app/module", request_template)
	GenerateModuleInByTemplate(params, "repo", "app/module", repo_template)
	GenerateModuleInByTemplate(params, "service", "app/module", service_template)
	GenerateModuleInByTemplate(params, "controller", "app/module", contoller_template)
	Appends(params)
}

func Appends(params CodeGenParams) {
	InjectCodeToPos("main.go", map[string]string{
		"// #inject:module ":      fmt.Sprintf("%v.New%vModule,\n", params.EntitySnake, params.Entity),
		"// #inject:moduleImport": fmt.Sprintf(`"gva/app/module/%v"`+"\n", params.EntitySnake),
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

func GenerateModuleByTemplate(params CodeGenParams, templateName string, directory string, templateContent string) {
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

func GenerateModuleInByTemplate(params CodeGenParams, templateName string, directory string, templateContent string) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		panic(err)
	}

	fullPath := directory + "/" + params.EntitySnake + "/" + templateName + "/" + params.EntitySnake + ".go"

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
