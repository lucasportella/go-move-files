package movefiles

import (

	"github.com/lucasportella/go-move-files/types"
)



func MoveFilesDefault(configuration types.Configuration) {
	for key, paths := range configuration.Default {
		files := ReadFilesFromSrcDir(paths.SrcDir)
		for _, file := range files {
			BuildDstFilePath(file, paths, key)
		}
	}
}
