package movefiles

import (
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/lucasportella/go-move-files/types"
)

func BuildDstFilePath(file fs.DirEntry, paths types.Paths, key string) error {
	srcPath := paths.SrcDir
	dstPath := paths.DstDir
	//openFile in the src dir
	oldFilePath := srcPath + "/" + file.Name()
	srcFile, err := OpenFile(oldFilePath)
	if err != nil {
		log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
	}

	// create file in dst dir if file matches the key
	newFilePath := dstPath + "/" + file.Name()
	if strings.Contains(strings.ToLower(file.Name()), strings.ToLower(key)) {
		
		dstFile, err := os.Create(newFilePath)
		if err != nil {
			log.Printf("Error while creating file in destiny folder: %v\n", err)
			dstFile.Close()
			return err
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
	return nil
}

func MoveFilesDefault(configuration types.Configuration) {
	for key, paths := range configuration.Default {
		files := ReadFilesFromSrcDir(paths.SrcDir)
		for _, file := range files {
			BuildDstFilePath(file, paths, key)
		}
	}
}
