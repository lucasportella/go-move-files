package movefiles

import (
	"io"
	"io/fs"
	"log"
	"os"

	"github.com/lucasportella/go-move-files/types"
	"github.com/lucasportella/go-move-files/utils"
)

func GetPaths() types.Configuration {
	paths, err := utils.ReadJSONFile()
	if err != nil {
		log.Fatal(err)
	}
	return paths
}

func MoveFiles(configuration types.Configuration) {
	for key, paths := range configuration.Default {
		MoveFilesDefault(key, paths)
	}

		MoveFilesWithDate(configuration)


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
	//fix: add delete fn here
	return err
}
