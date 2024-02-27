package utils

import (
	"encoding/json"
	"log"
	"os"
	"runtime"

	types "github.com/lucasportella/go-move-files/types"
)

func GetUserOS() string {
	return runtime.GOOS
}

func ReadJSONFile() (types.Configuration, error) {
	filePath := "paths.json"

	jsonPaths, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the JSON file: %v", err)
	}

	var paths types.Configuration
	err = json.Unmarshal(jsonPaths, &paths)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}
	return paths, err
}
