package main

import (
	"fmt"
	"log"

	types "github.com/lucasportella/go-move-files/types"
	utils "github.com/lucasportella/go-move-files/utils"
)



func main() {
	paths := getPaths()
	for level1Key, level2Map := range paths {
		for level2Key, level3Map := range level2Map {
			switch level1Key {
			case "default":
				moveFilesDefault(level2Key, level3Map)
			case "withDate":
				moveFilesWithDate(level2Key, level3Map)
			}
		}
	}
	// moveFiles(paths)

}

func getPaths() types.Paths {
	paths, err := utils.ReadJSONFile()
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

// func moveFiles(paths types.Paths, pathKey string) {
// 	srcPath := paths[pathKey]["src_dir"]
// 	dstPath := paths["dst_dir"]
// 	src, err := os.Open(srcPath)

// 	if err != nil {
// 		log.Fatalf("Fatal! Could not open source directory. Error: %v", err)
// 	}
// 	defer src.Close()

// 	files, err := src.ReadDir(-1)

// 	if err != nil {
// 		log.Fatalf("Fatal! Could not read source directory. Error: %v", err)
// 	}

// 	for _, file := range files {
// 		if strings.HasPrefix(file.Name(), "test") {
// 			os.Rename(srcPath+"/"+file.Name(), dstPath+"/"+file.Name())
// 		}
// 	}
// }

func WalkThroughPaths(paths types.Paths) {

}

func moveFilesDefault(key string, mapPaths types.InnerPaths) {
	fmt.Println("moveFilesDefault called!")
}

func moveFilesWithDate(key string, mapPaths types.InnerPaths) {
	fmt.Println("moveFilesWithDate called!")
}
