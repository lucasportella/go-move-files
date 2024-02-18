package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	paths := getPaths()
	if getUserOS() == "windows" {
		paths = fixWindowsSlashes(paths)
	}
	moveFiles(paths)

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

func moveFiles(paths map[string]string) {
	srcPath := paths["src_dir"]
	dstPath := paths["dst_dir"]
	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatalf("Fatal! Could not open source directory. Error: %v", err)
	}
	defer src.Close()

	files, err := src.ReadDir(-1)

	if err != nil {
		log.Fatalf("Fatal! Could not read source directory. Error: %v", err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "test") {
			os.Rename(srcPath+"\\"+file.Name(), dstPath+"\\"+file.Name())
		}
	}
}
