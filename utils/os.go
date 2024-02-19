package utils

import (
	"encoding/json"
	"os"
	"log"
	"runtime"
)

func GetUserOS() string {
	return runtime.GOOS
}

func ReadJSONFile() (map[string]string, error) {
	filePath := "paths.json"

	jsonPaths, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the JSON file: %v", err)
	}

	var paths map[string]string
	err = json.Unmarshal(jsonPaths, &paths)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}
	return paths, err
}
