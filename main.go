package main

import (
	movefiles "github.com/lucasportella/go-move-files/moveFiles"
)

func main() {
	paths := movefiles.GetPaths()
	movefiles.MoveFiles(paths)
}
