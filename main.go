// Package main is the main app package
package main

import (
	"flag"
	"os"
	"star-request/utils"

	"github.com/pterm/pterm"
)

var (
	configPath   *string
	excludeFiles *string
	files        []string
)

func init() {
	pterm.EnableDebugMessages()

	// get program arguments
	configPath = flag.String("config", "", "Path to your request config directory. Can be relative or absolute path.")
	excludeFiles = flag.String("exclude", "", "Filenames from given/config directory to exclude. Regex supported.")

	flag.Parse()
}

func main() {

	pterm.Debug.Println(flag.CommandLine.Args())

	// If no args, skip exe
	if len(os.Args) < 2 {
		pterm.Debug.Println("Looking in current directory...")
		dir, _ := os.Getwd()
		// Find *.sr.[json,yaml] in current directory
		files = utils.FindConfigFile(dir)
	}

	if *configPath != "" {
		pterm.Debug.Println("Looking in user defined directory...")
		// Find *.sr.[json,yaml] in given directory(s)
		files = utils.FindConfigFile(*configPath)
	}

	pterm.Info.Println("Attempting to run requests found in:", files)
}
