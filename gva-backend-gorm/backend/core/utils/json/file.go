package json

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// using base path ./app/....
//
// can't use relative path
func ReadJsonFile(path string) (JSON, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("can't get absolute path %v", err)
	}

	jsonFile, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("can't open file %v", err)
	}

	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("can't read jsonFile %v", err)
	}

	return JSON(bytes), nil
}

func WriteJsonToFile(data JSON, path string) error {
	// Convert the base path to an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("can't get absolute path %v", err)
	}

	// Create the file
	file, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("can't create file %v", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("can't write to file %v", err)
	}

	return nil
}
