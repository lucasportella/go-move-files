package movefiles

import (
	"github.com/lucasportella/go-move-files/types"
)

func MoveFilesDefault(configuration types.Configuration) {
	for key, paths := range configuration.Default {
		files := ReadFilesFromSrcDir(paths.SrcDir)
		for _, file := range files {
			if !FileNameContainsKey(file.Name(), key) {
				continue
			}
			err := MkdirNewFolders(file, paths.DstDir, key)
			if err != nil {
				continue
			}
			MoveFile(paths.SrcDir, paths.DstDir)
		}
	}
}
