// Package request is the main app core
package request

import (
	"os"
	"regexp"

	"github.com/pterm/pterm"
)

var (
	// The config file to read request properties from
	config string
)

// Execute starts the app
func Execute(args []string) {
	if len(args) < 2 {
		dir, _ := os.Getwd()
		// Find *.sr.[json,yaml] in current directory
		files, _ := findConfigFile(dir)
		pterm.Info.Println(files[0])
	}
}

func findConfigFile(dir string) ([]string, error) {
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
