package movefiles

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lucasportella/go-move-files/types"
)

func MoveFilesWithDate(key string, innerPaths types.InnerPaths) {
	// get paths from json
	srcPath := innerPaths.Src_dir
	// dstPath := innerPaths.Dst_dir

	files := ReadFilesFromSrcDir(srcPath)

	for _, file := range files {
		//openFile in the src dir
		oldFilePath := srcPath + "/" + file.Name()
		fileInfo, err := GetFileDate(oldFilePath)
		if err != nil {
			continue
		}
		fmt.Println(fileInfo)
		_, err = OpenFile(oldFilePath)
		if err != nil {
			log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
		}

	}
}

// GetFileDate gets the date of the file and returns a string in the format YYYY-MM-DD
func GetFileDate(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("error while getting data from the following file: %v", path)
	}
	formattedTime := fileInfo.ModTime().Format(time.RFC3339)
	relevantTime := strings.Split(formattedTime, "T")[0]
	return relevantTime, nil
}

// verify if correct folder exists
// if exists create it, if not, do nothing
// update newFilePath
// moveFile
func GetDatePaths(key string, config string, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := CreateDateFolders(path)
		if err != nil {
			return fmt.Errorf("failed to create folder(s): %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error while checking folder existence: %v", err)
	}
	return nil
}

func CreateDateFolders(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfFolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
