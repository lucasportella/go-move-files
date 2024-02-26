package movefiles

import (
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
		log.Fatalf("Error while getting data from the following file: %v", path)
		return "", err
	}
	formattedTime := fileInfo.ModTime().Format(time.RFC3339)
	relevantTime := strings.Split(formattedTime, "T")[0]
	return relevantTime, nil
}

func CreateDateFolders(key string) {
	
}