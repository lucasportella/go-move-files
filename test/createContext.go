package test

import (
	"fmt"
	"os"
	"path/filepath"
)

func Setup() (string, error) {
	projectRoot, err := FindProjectRoot()
	if err != nil {
		return "", err
	}

	directories := []string{
		filepath.Join(projectRoot, "testTemp"),
		filepath.Join(projectRoot, "testTemp", "src"),
		filepath.Join(projectRoot, "testTemp", "dst"),
	}

	files := [][]string{
		{"#bill-1", "content of bill 1"},
		{"#bill-2", "content of bill 2"},
		{"#doc-1", "content of doc 1"},
		{"#doc-2", "content of doc 2"},
		{"#book-1", "content of book 1"},
		{"#book-2", "content of book 2"},
		{"#music-1", "content of music 1"},
		{"#music-2", "content of music 2"},
		{"#tax-1", "content of tax 1"},
		{"#tax-2", "content of tax 2"},
	}

	for _, dir := range directories {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	for _, file := range files {
		osFile, err := os.Create(filepath.Join(projectRoot, "testTemp", "src", file[0]))

		if err != nil {
			return "", fmt.Errorf("error creating file %s: %w", file, err)
		}
		defer osFile.Close()

		_, err = osFile.WriteString(file[1])

		if err != nil {
			return "", fmt.Errorf("error while writing content in file %v", file[0])
		}
	}

	return projectRoot, nil
}

func Teardown() error {
	projectRoot, err := FindProjectRoot()
	testTempDir := filepath.Join(projectRoot, "testTemp")
	if err != nil {
		return err
	}
	err = os.RemoveAll(testTempDir)
	if err != nil {
		fmt.Printf("error removing directory %s: %v\n", testTempDir, err)
	}
	return nil
}

func FindProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			return cwd, nil
		}
		newCwd := filepath.Dir(cwd)
		if newCwd == cwd {
			return "", fmt.Errorf("project root not found")
		}
		cwd = newCwd
	}
}
