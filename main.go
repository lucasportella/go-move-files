package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	types "github.com/lucasportella/go-move-files/types"
	utils "github.com/lucasportella/go-move-files/utils"
)

func main() {
	paths := getPaths()
	moveFiles(paths)
}

func getPaths() types.Paths {
	paths, err := utils.ReadJSONFile()
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

func moveFiles(paths types.Paths) {
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
}

func WalkThroughPaths(paths types.Paths) {

}

func moveFilesDefault(key string, innerPaths types.InnerPaths) {
	srcPath := innerPaths.Src_dir
	dstPath := innerPaths.Dst_dir
	src, err := os.Open(srcPath)

	if err != nil {
		log.Fatalf("Fatal! Could not open source directory. Error: %v\n", err)
	}
	defer src.Close()

	files, err := src.ReadDir(-1)

	if err != nil {
		log.Fatalf("Fatal! Could not read source directory. Error: %v\n", err)
	}

	for _, file := range files {

		oldFilePath := srcPath + "/" + file.Name()
		srcFile, err := openFile(oldFilePath)
		if err != nil {
			log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
		} else {
			defer srcFile.Close()
		}

		newFilePath := dstPath + "/" + file.Name()
		if strings.HasPrefix(file.Name(), key) {
			dstFile, err := os.Create(newFilePath)
			if err != nil {
				log.Printf("Error while creating file in destiny folder: %v\n", err)
				continue
			}

			_, err = io.Copy(dstFile, srcFile)
			srcFile.Close()
			if err != nil {
				fmt.Printf("Error while moving file: %v\n", err)
				deleteFile(newFilePath)
			} else {
				deleteFile(oldFilePath)
			}
		}
	}
}

func moveFilesWithDate(key string, innerPaths types.InnerPaths) {

}

func deleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Error, could not delete the file in this path: %v. Error message: %v\n", filePath, err)
	}
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
