package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

func main() {
	paths := getPaths()
	if getUserOS() == "windows" {
		paths = fixWindowsSlashes(paths)
	}
	fmt.Println(paths)
}

func getUserOS() string {
	return runtime.GOOS
}

func replaceSlashes(path string) string {
	return strings.ReplaceAll(path, `/`, `\`)
}

func fixWindowsSlashes(paths map[string]string) map[string]string {
	updatedPaths := make(map[string]string)
	for key, value := range paths {
		updatedPaths[key] = replaceSlashes(value)
	}
	return updatedPaths
}

func getPaths() map[string]string {
	filePath := "paths.json"

	jsonPaths, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the JSON file: %v", err)
	}

	var paths map[string]string
	err = json.Unmarshal(jsonPaths, &paths)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}
	return paths
}
