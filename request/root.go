// Package request is the main app core
package request

import (
	"os"
	"star-request/utils"

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
		files, _ := utils.FindConfigFile(dir)
		pterm.Info.Println(files[0])
	}
}
