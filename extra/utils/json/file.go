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
