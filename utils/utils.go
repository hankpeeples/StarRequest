// Package utils provides basic utilities
package utils

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/kpango/glg"
	"github.com/pterm/pterm"
)

var SEP string

type Log struct {
	*glg.Glg
}

var l *Log

func CreateLogger(debug bool) {
	var level glg.LEVEL
	level = glg.TagStringToLevel("error")
	if debug {
		level = glg.TagStringToLevel("debug")
	}

	// Initialize logger to print to terminal and log file
	logFile := glg.FileWriter("./sr.log", 0666)
	logger := glg.Get().SetMode(glg.BOTH).SetLineTraceMode(glg.TraceLineShort).
		InitWriter().
		AddWriter(logFile).
		SetWriter(logFile).
		AddLevelWriter(glg.LOG, logFile).
		AddLevelWriter(glg.INFO, logFile).
		AddLevelWriter(glg.WARN, logFile).
		AddLevelWriter(glg.ERR, logFile).
		SetLevelWriter(glg.LOG, logFile).
		SetLevelWriter(glg.INFO, logFile).
		SetLevelWriter(glg.WARN, logFile).
		SetLevelWriter(glg.ERR, logFile).
		SetLevel(level)

	l = &Log{logger}
}

func GetLogger() *Log {
	return l
}

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
	if runtime.GOOS == "windows" {
		SEP = "\\"
	} else {
		SEP = "/"
	}

	var filenames []string
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
		filenames = append(filenames, dir+SEP+file.Name())
	}

	return getFilePath(filenames)
}

// getFilePath matches correct filenames
func getFilePath(files []string) []string {
	var found []string

	// find which filenames are valid config files
	for _, file := range files {
		match, _ := regexp.MatchString(".sr.(json|yaml)", file)
		if match {
			found = append(found, file)
		}
	}

	return found
}

// BuildFileList creates an easily readable list of found request config files
func BuildFileList(files []string) string {
	var str string
	listLen := len(files) - 1

	for i, file := range files {
		if i == listLen {
			str += file
		} else {
			str += file + ", "
		}
	}

	return str
}

// GetFile returns only the given directories *.sr.[json,yaml] file
func GetFile(dir string) string {
	var dirArr []string

	dirArr = strings.Split(dir, SEP)

	return dirArr[len(dirArr)-1]
}
