package movefiles

import (
	"log"
	"os"

	"github.com/lucasportella/go-move-files/types"
)

func MoveFilesDefault(configuration types.Configuration) {
	for key, paths := range configuration.Default {

		files, err := os.ReadDir(paths.SrcDir)
		if err != nil {
			log.Fatalf("Fatal! Could not read source directory. Error: %v\n", err)
		}

		for _, file := range files {
			if !FileNameContainsKey(file.Name(), key) {
				continue
			}
			err := MkdirNewFolders(file, paths.DstDir, key)
			if err != nil {
				log.Println(err)
				continue
			}
			MoveFile(paths.DstDir, paths.SrcDir, file.Name())
		}
	}
}
