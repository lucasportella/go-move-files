package movefiles

import "testing"

type PathData struct {
	path   string
	result bool
}

func TestPathExists(t *testing.T) {
	pathsData := []PathData{
		{path: "C:/Users/lucas/Pictures/", result: true},
		{path: "C:/Users/lucas/Videos/jpwnnpngiewahns", result: false},
	}
	for _, data := range pathsData {
		result := PathExists(data.path)
		if data.result != result {
			t.Errorf("Expected %v, got: %v", data.result, result)
		}
	}
}
