// Package main is the main app package
package main

import (
	"flag"
	"os"
	"star-request/request"
	"star-request/utils"

	"github.com/pterm/pterm"
)

var (
	configPath   *string
	excludeFiles *string
	debug        *bool
	recursive    *bool
	files        []string
)

func init() {
	// get program arguments
	configPath = flag.String("config", "", "Path to your request config directory. Can be relative or absolute path.")
	excludeFiles = flag.String("exclude", "", "(Not Implemented!) Filenames from given/config directory to exclude. Regex supported.")
	debug = flag.Bool("debug", true, "Whether to show debug log messages.")
	recursive = flag.Bool("recursive", false, "(Not Implemented!) Whether to look in all sub directories of given path.")

	flag.Parse()
}

func main() {
	if *debug {
		pterm.EnableDebugMessages()
	}

	pterm.Debug.Println("configPath:", *configPath)
	pterm.Debug.Println("excludeFiles:", *excludeFiles)
	pterm.Debug.Println("debug:", *debug)
	pterm.Debug.Println("recursive:", *recursive)

	// If no args, skip exe
	if len(os.Args) < 2 {
		dir, _ := os.Getwd()
		// Find *.sr.[json,yaml] in current directory
		files = utils.FindConfigFile(dir)
	}

	if *configPath != "" {
		// Find *.sr.[json,yaml] in given directory(s)
		files = utils.FindConfigFile(*configPath)
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, utils.GetFile(file))
	}

	pterm.Info.Println("Attempting to run requests found in:\n\t", utils.BuildFileList(fileList))

	for _, file := range files {
		request.BuildRequest(utils.ParseConfig(file))
	}
}
