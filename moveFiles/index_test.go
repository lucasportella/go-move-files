package movefiles

import (
	"path/filepath"
	"testing"

	"github.com/lucasportella/go-move-files/test"
)

type PathData struct {
	path   string
	result bool
}

func TestPathExists(t *testing.T) {
	projectRoot, err := test.Setup()
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}
	pathsData := []PathData{
		{path: filepath.Join(projectRoot, "test", "src"), result: true},
		{path: filepath.Join(projectRoot, "test", "dst"), result: true},
		{path: filepath.Join(projectRoot, "test", "notExistPath"), result: false},
	}
	for _, data := range pathsData {
		result := PathExists(data.path)
		if data.result != result {
			t.Errorf("Expected %v, got: %v", data.result, result)
		}
	}
	err = test.Teardown()
	if err != nil {
		t.Fatalf("Teardown failed: %v", err)
	}
}
