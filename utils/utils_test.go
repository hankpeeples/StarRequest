package utils

import (
	"testing"
)

func TestGetFilePath(t *testing.T) {
	files := []string{"one.sr.log", "two.se.json", "three.sr.yaml", "fake.sr.js"}
	ret := getFilePath(files)
	if ret[0] != "three.sr.yaml" {
		t.Errorf("Result incorrect, got: %s, want: %s.", ret[0], files[2])
	}
}
