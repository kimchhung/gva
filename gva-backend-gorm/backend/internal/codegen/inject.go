package codegen

import (
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
		if strings.Contains(strContent, value) {
			continue
		}

		insertStruct := strings.LastIndex(strContent, key)
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
}

func formatFile(filePath string) {
	cmd := exec.Command("gofmt", "-w", filePath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
