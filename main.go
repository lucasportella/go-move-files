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
	// get paths from json
	srcPath := innerPaths.Src_dir
	dstPath := innerPaths.Dst_dir

	// read files from src dir
	files, err := os.ReadDir(srcPath)
	if err != nil {
		log.Fatalf("Fatal! Could not read source directory. Error: %v\n", err)
	}

	//loop through src dir files
	for _, file := range files {

		//openFile in the src dir
		oldFilePath := srcPath + "/" + file.Name()
		srcFile, err := openFile(oldFilePath)
		if err != nil {
			log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
		}
		defer srcFile.Close()

		// create file in dst dir if file matches the key
		newFilePath := dstPath + "/" + file.Name()
		if strings.HasPrefix(file.Name(), key) {
			dstFile, err := os.Create(newFilePath)
			if err != nil {
				log.Printf("Error while creating file in destiny folder: %v\n", err)
				dstFile.Close()
				continue
			}
			defer dstFile.Close()
			moveFile(dstFile, srcFile, newFilePath, oldFilePath)

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

func moveFile(dstFile *os.File, srcFile *os.File, newFilePath string, oldFilePath string) {
	// copy the file from src to dst and delete src file if success or delete dst file if fails
	_, err := io.Copy(dstFile, srcFile)
	srcFile.Close()
	if err != nil {
		fmt.Printf("Error while moving file: %v\n", err)
		deleteFile(newFilePath)
	} else {
		deleteFile(oldFilePath)
	}
}
