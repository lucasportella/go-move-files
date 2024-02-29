package movefiles

import (
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
	paths, err := utils.ReadJSONFile()
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

func ReadFilesFromSrcDir(srcPath string) []fs.DirEntry {
	files, err := os.ReadDir(srcPath)
	if err != nil {
		log.Fatalf("Fatal! Could not read source directory. Error: %v\n", err)
	}
	return files
}

func BuildDstFilePath(file fs.DirEntry, paths types.Paths, key string) error {
	srcPath := paths.SrcDir
	dstPath := paths.DstDir
	formattedFileName := strings.ToLower(file.Name())
	formattedKey := strings.ToLower(key)
	oldFilePath := srcPath + "/" + formattedFileName

	//openFile in the src dir
	srcFile, err := os.Open(oldFilePath)
	if err != nil {
		log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
	}
	defer srcFile.Close()

	// create file in dst dir if file matches the key
	newFilePath := dstPath + "/" + formattedFileName
	if strings.Contains(formattedFileName, formattedKey) {
		if !PathExists(dstPath) {
			CreateFolders(dstPath)
		}
		dstFile, err := os.Create(newFilePath)
		if err != nil {
			log.Printf("Error while creating file in destiny folder: %v\n", err)
			return err
		}
		defer dstFile.Close()

		MoveFile(dstFile, srcFile, oldFilePath, newFilePath)
	}
	return nil
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Error, could not delete the file in this path: %v. Error message: %v\n", filePath, err)
	}
}

func MoveFile(dstFile *os.File, srcFile *os.File, oldFilePath string, newFilePath string) error {
	_, err := io.Copy(dstFile, srcFile)
	srcFile.Close()
	dstFile.Close()
	if err != nil {
		DeleteFile(newFilePath)
	} else {
		DeleteFile(oldFilePath)
	}
	return err
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
