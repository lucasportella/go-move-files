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

	for _, dir := range directories {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("error creating directory %s: %w", dir, err)
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
