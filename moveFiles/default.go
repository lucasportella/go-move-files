package movefiles

import (
	"log"
	"os"
	"strings"

	"github.com/lucasportella/go-move-files/types"
)

func MoveFilesDefault(key string, innerPaths types.InnerPaths) {
	// get paths from json
	srcPath := innerPaths.Src_dir
	dstPath := innerPaths.Dst_dir

	files := ReadFilesFromSrcDir(srcPath)

	for _, file := range files {

		//openFile in the src dir
		oldFilePath := srcPath + "/" + file.Name()
		srcFile, err := OpenFile(oldFilePath)
		if err != nil {
			log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
		}

		// create file in dst dir if file matches the key
		newFilePath := dstPath + "/" + file.Name()
		if strings.HasPrefix(file.Name(), key) {
			dstFile, err := os.Create(newFilePath)
			if err != nil {
				log.Printf("Error while creating file in destiny folder: %v\n", err)
				dstFile.Close()
				continue
			}

			err = MoveFile(dstFile, srcFile)
			srcFile.Close()
			dstFile.Close()
			if err != nil {
				DeleteFile(newFilePath)
			} else {
				DeleteFile(oldFilePath)
			}
		}
	}
}
