// Package main is the main app package
package main

import (
	"os"
	"star-request/request"
)

func main() {
	request.Execute(os.Args)
}
