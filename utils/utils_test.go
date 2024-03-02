package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	files := []string{"one.sr.log", "two.se.json", "three.sr.yaml", "fake.sr.js"}
	ret := getFilePath(files)
	if ret[0] != "three.sr.yaml" {
		t.Errorf("Result incorrect, got: %s, want: %s.", ret[0], files[2])
	}
}

func TestJSONParse(t *testing.T) {
	name := "Request Name"
	type JSON struct {
		Name        string `json:"name"`
		URL         string `json:"URL"`
		Method      string `json:"method"`
		ContentType string `json:"content-type"`
	}

	var jsonStruct JSON

	jsonFile, err := os.Open("/Users/hankpeeples/go/src/github.com/hankpeeples/StarRequest/example.sr.json")
	if err != nil {
		t.Errorf("Unable to open file: %v", err)
	}
	defer jsonFile.Close()

	goodJSON, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Read JSON file error: %s", err)
	}

	err = json.Unmarshal(goodJSON, &jsonStruct)
	if err != nil {
		t.Errorf("Unmarshal error: %s", err)
	}

	fmt.Println(jsonStruct)
	if jsonStruct.Name != name {
		t.Errorf("got '%s', expected '%s'", jsonStruct.Name, name)
	}
}
