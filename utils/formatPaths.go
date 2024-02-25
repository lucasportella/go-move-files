package utils

import (
	"strings"
)

func ReplaceSlashes(path string) string {
	return strings.ReplaceAll(path, `/`, `\`)
}

func FixWindowsSlashes(paths map[string]string) map[string]string {
	updatedPaths := make(map[string]string)
	for key, value := range paths {
		updatedPaths[key] = ReplaceSlashes(value)
	}
	return updatedPaths
}

func FormatPaths(paths map[string]string) map[string]string {
	userOS := GetUserOS()
	var formattedPaths map[string]string
	if userOS == "windows" {
		formattedPaths = FixWindowsSlashes(paths)
	}
	return formattedPaths
}
