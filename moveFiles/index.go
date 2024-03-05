package movefiles

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/lucasportella/go-move-files/types"
	"github.com/lucasportella/go-move-files/utils"
)

func MoveFiles(configuration types.Configuration) {
	MoveFilesDefault(configuration)
	MoveFilesWithDate(configuration)
}

func GetPaths() types.Configuration {
	paths, err := utils.ReadJSONFile("paths.json")
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

func FileNameContainsKey(fileName string, key string) bool {
	return strings.Contains(fileName, key)
}

func MkdirNewFolders(file fs.DirEntry, dstPath string, key string) error {
	if !PathExists(dstPath) {
		err := CreateFolders(dstPath)
		if err != nil {
			return fmt.Errorf("error while creating destiny folders: %v", err)
		}
	}
	return nil
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Error, could not delete the file in this path: %v. Error message: %v\n", filePath, err)
	}
}

func MoveFile(dstPath string, srcPath string, fileName string) {
	oldFilePath := srcPath + "/" + fileName
	newFilePath := dstPath + "/" + fileName

	dstFile, err := os.Create(newFilePath)
	if err != nil {
		log.Printf("could not open file %v. Skipping to next file", newFilePath)
	}
	defer dstFile.Close()

	srcFile, err := os.Open(oldFilePath)
	if err != nil {
		log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		dstFile.Close()
		DeleteFile(dstPath)
	} else {
		srcFile.Close()
		DeleteFile(oldFilePath)
	}

	if err != nil {
		log.Printf("error while moving file: %v", err)
	} else {
		log.Printf("file %v moved to %v with success!", srcPath, newFilePath)
	}

}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
