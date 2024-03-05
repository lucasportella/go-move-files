package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"

	types "github.com/lucasportella/go-move-files/types"
)

func GetUserOS() string {
	return runtime.GOOS
}

func ReadJSONFile(filePath string) (types.Configuration, error) {
	jsonPaths, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("error reading the JSON file: %v", err)
		return types.Configuration{}, fmt.Errorf("error reading the JSON file: %v", err)
	}

	var paths types.Configuration
	err = json.Unmarshal(jsonPaths, &paths)
	if err != nil {
		log.Printf("error parsing JSON file: %v", err)
		return types.Configuration{}, fmt.Errorf("error parsing the JSON file: %v", err)
	}
	return paths, err
}
