// Package request is the main builder and runner of parsed requests
package request

import (
	"io"
	"net/http"
	"os"
	"star-request/utils"

	"github.com/pterm/pterm"
)

// SendRequest creates the base request parsed from the file
func SendRequest(requests utils.JSONRequests) {
	for _, req := range requests.Requests {
		pterm.Info.Printf("Name: %s ~ %s request to %s\n", req.Name, req.Method, req.URL)
		buildRequest(req)
	}
}

func buildRequest(req utils.Request) {
	switch req.Method {
	case GET:
		sendGetRequest(req)
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

func sendGetRequest(req utils.Request) {
	pterm.Debug.Println("Running GET request")

	res, err := http.Get(req.URL)
	if err != nil {
		pterm.Error.Println(err)
	}

	if !checkStatusCode(res.StatusCode, res.Status) {
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	pterm.Info.Printf("Status Code: %s ~ Data: %v\n", res.Status, body)
}

func checkStatusCode(statusCode int, status string) bool {
	if statusCode != 200 {
		pterm.Warning.Printf("Bad request ... Status: %s\n", status)
		return false
	}
	return true
}
