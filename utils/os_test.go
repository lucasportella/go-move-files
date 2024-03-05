package utils

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/lucasportella/go-move-files/test"
)

func TestReadJSONFile(t *testing.T) {
	projectRoot, err := test.Setup()
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}
	defer test.Teardown()

	_, err = ReadJSONFile(filepath.Join(projectRoot, "tests1.json"))
	if err != nil {
		t.Fatalf("TestReadJSONFile failed. %v", err)
	}

	_, err = ReadJSONFile("wrongPath123.json")
	if !strings.Contains(err.Error(), "error reading the JSON file:") {
		t.Fatalf("TestReadJSONFile failed. %v", err)
	}

	_, err = ReadJSONFile(filepath.Join(projectRoot, "tests2.json"))
	if err != nil {
		if !strings.Contains(err.Error(), "error parsing the JSON file:") {
			t.Fatalf("TestReadJSONFile failed. %v", err)
		}
	}

}
