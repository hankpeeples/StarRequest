// Package request is the main builder and runner of parsed requests
package request

import (
	"star-request/utils"

	"github.com/pterm/pterm"
)

// BuildRequest creates the base request parsed from the file
func BuildRequest(requests utils.JSONRequests) {
	for _, req := range requests.Requests {
		switch req.Method {
		case GET:
			pterm.Debug.Println("GET request")
			break
		case POST:
			pterm.Debug.Println("POST request")
			break
		case PUT:
			pterm.Debug.Println("PUT request")
			break
		case DELETE:
			pterm.Debug.Println("DELETE request")
			break
		default:
			pterm.Error.Println("Unknown request method:", req.Method)
			break
		}
	}
}
