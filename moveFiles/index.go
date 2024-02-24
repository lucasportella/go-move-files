package movefiles

import (
	"io"
	"io/fs"
	"log"
	"os"

	"github.com/lucasportella/go-move-files/types"
	"github.com/lucasportella/go-move-files/utils"
)

func GetPaths() types.Paths {
	paths, err := utils.ReadJSONFile()
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

func MoveFiles(paths types.Paths) {
	for level1Key, level2Map := range paths {
		for level2Key, level3Map := range level2Map {
			switch level1Key {
			case "default":
				MoveFilesDefault(level2Key, level3Map)
			case "withDate":
				MoveFilesWithDate(level2Key, level3Map)
			}
		}
	}
}

func ReadFilesFromSrcDir(srcPath string) []fs.DirEntry {
	files, err := os.ReadDir(srcPath)
	if err != nil {
		log.Fatalf("Fatal! Could not read source directory. Error: %v\n", err)
	}
	return files
}

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Error, could not delete the file in this path: %v. Error message: %v\n", filePath, err)
	}
}

func MoveFile(dstFile *os.File, srcFile *os.File) error {
	_, err := io.Copy(dstFile, srcFile)
	return err
}
