// Package utils provides basic utilities
package utils

import (
	"os"
	"regexp"

	"github.com/pterm/pterm"
)

// FindConfigFile finds all request configs in the given directory, Either the current
// dir or a config file/files via given args.
func FindConfigFile(dir string) ([]string, error) {
	filenames := []string{}
	pterm.Info.Printf("Looking for config file in '%s'\n", dir)

	// open directory for reading
	thisDir, err := os.Open(dir)
	if err != nil {
		pterm.Error.Println("Could not open current directory ...", err)
	}

	// read dir and get all files
	files, err := thisDir.ReadDir(-1)
	if err != nil {
		pterm.Error.Println("Could not read current directory ...", err)
	}

	// get only filenames
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return getFilePath(filenames), nil
}

// GetFilePath matches correct filenames
func getFilePath(files []string) []string {
	found := []string{}

	// find which filenames are valid config files
	for _, file := range files {
		match, _ := regexp.MatchString(".sr.(json|yaml)", file)
		if match {
			found = append(found, file)
		}
	}

	return found
}
