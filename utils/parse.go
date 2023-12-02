package utils

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

// JSONRequests is the structure of the json request configs.
type JSONRequests struct {
	// Requests - The array of requests
	Requests []request `json:"requests"`
}

type request struct {
	// Name - A requests name
	Name string `json:"name"`
	// ULR - Request endpoint
	URL string `json:"url"`
	// Method - Request method
	Method string `json:"method"`
}

// ParseConfig reads the request config file(s)
func ParseConfig(file string) {
	splitFile := strings.Split(file, ".")
	if splitFile[len(splitFile)-1] == "json" {
		parseJSONConfig(file)
	}
}

func parseJSONConfig(file string) {
	pterm.Debug.Println("Parsing JSON config file:", file)

	jsonFile, err := os.Open(file)
	if err != nil {
		pterm.Fatal.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	// read json as a byte array.
	byteArr, err := io.ReadAll(jsonFile)
	if err != nil {
		pterm.Error.Println("JSON file read error:", err)
	}

	var requests JSONRequests

	err = json.Unmarshal(byteArr, &requests)
	if err != nil {
		pterm.Error.Println("JSON unmarshal error:", err)
	}

	pterm.Info.Println(requests)
}
