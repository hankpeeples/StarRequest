// Package utils provides basic utilities
package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/pterm/pterm"
)

// unknownArgs show the unknown args error message
func unknownArgs(args []string, index int) {
	fmt.Println("Unknown args:")
	for i, arg := range args {
		if i == index {
			fmt.Printf("\t%v\n", arg)
		}
	}
	os.Exit(1)
}

// FindConfigFile finds all request configs in the given directory, Either the current
// dir or a config file/files via given args.
func FindConfigFile(dir string) []string {
	filenames := []string{}
	pterm.Debug.Printf("Looking for config file in '%s'\n", dir)

	// open directory for reading
	thisDir, err := os.Open(dir)
	if err != nil {
		pterm.Error.Println("Could not open current director:\n\t", err.Error())
		os.Exit(17)
	}

	// read dir and get all files
	files, err := thisDir.ReadDir(-1)
	if err != nil {
		pterm.Error.Println("Could not read current directory:\n\t", err.Error())
		os.Exit(18)
	}

	// get only filenames
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return getFilePath(filenames)
}

// getFilePath matches correct filenames
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
