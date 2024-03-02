// Package utils provides basic utilities
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
	Requests []Request `json:"requests"`
}

// Request defines a single request
type Request struct {
	// Name - A requests name
	Name string `json:"name"`
	// ULR - Request endpoint
	URL string `json:"url"`
	// Method - Request method
	Method string `json:"method"`
	// ContentType - Request conetent type
	ContentType string `json:"content-type"`
}

// ParseConfig reads the request config file(s)
func ParseConfig(file string) JSONRequests {
	var requests JSONRequests
	splitFile := strings.Split(file, ".")
	if splitFile[len(splitFile)-1] == "json" {
		requests = parseJSONConfig(file)
	}
	return requests
}

func parseJSONConfig(file string) JSONRequests {
	log := GetLogger()
	log.Infof("Parsing JSON config file: %v", file)

	jsonFile, err := os.Open(file)
	if err != nil {
		pterm.Fatal.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	// read json as a byte array.
	byteArr, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("JSON file read error: %v", err)
	}

	var requests JSONRequests

	err = json.Unmarshal(byteArr, &requests)
	if err != nil {
		log.Fatal("JSON unmarshal error:", err)
	}

	for n, req := range requests.Requests {
		if req.Name == "" {
			log.Warnf("Request #%d has no name..", n+1)
			requests.Requests[n].Name = "N/A"
		}
		if req.URL == "" {
			log.Fatalf("Required field 'URL' not found on request '%s'\n", req.Name)
		}
	}

	return requests
}
