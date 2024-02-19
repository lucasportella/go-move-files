package main

import (
	"log"
	"os"
	"strings"

	utils "github.com/lucasportella/go-move-files/utils"
)


func main() {
	paths := getPaths()
	moveFiles(paths)

}

func getPaths() map[string]string {
	paths, err := utils.ReadJSONFile()
	if err != nil {
		log.Fatal(err)
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
			os.Rename(srcPath+"/"+file.Name(), dstPath+"/"+file.Name())
		}
	}
}
