package codegen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func InjectCodeToPos(filePath string, pos map[string]string, format bool) {
	// Read the existing file
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Convert the file content to a string
	strContent := string(content)
	// Find the location to insert the new line
	for key, value := range pos {
		fileContext := strings.TrimSpace(strContent)
		text := strings.TrimSpace(value)

		if strings.Contains(fileContext, value) {
			fmt.Println("Skipping text injection in file:"+filePath+" for existing text:", text)
			continue
		}

		insertStruct := strings.LastIndex(strContent, key)
		value += "\n\t"
		strContent = strContent[:insertStruct] + value + strContent[insertStruct:]
	}

	// Write the modified content back to the file
	err = os.WriteFile(filePath, []byte(strContent), 0644)
	if err != nil {
		panic(err)
	}

	if format {
		formatFile(filePath)
	}

	fmt.Println("Modified code to file:", filePath)
}

func formatFile(filePath string) {
	cmd := exec.Command("gofmt", "-w", filePath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
