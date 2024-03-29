// Package request is the main builder and runner of parsed requests
package request

import (
	"io"
	"net/http"
	"time"

	"star-request/jsonhelpers"
	"star-request/utils"

	"github.com/pterm/pterm"
)

// SendRequest creates the base request parsed from the file
func SendRequest(requests utils.JSONRequests) {
	log := utils.GetLogger()
	for _, req := range requests.Requests {
		log.Infof("Running: name: %s ~ %s request to %s ...", req.Name, req.Method, req.URL)
		// pterm.Info.Printf("Name: %s ~ %s request to %s\n", req.Name, req.Method, req.URL)
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

	request, err := http.NewRequest(http.MethodGet, req.URL, nil)
	if err != nil {
		pterm.Error.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(request)
	if err != nil {
		pterm.Error.Println(err)
	}

	body := getRequestData(resp)

	jsonhelpers.Format(body)
}

func checkStatusCode(statusCode int, status string) {
	if statusCode == 404 {
		pterm.Error.Printf("Bad request ... Returned with status: %s\n", status)
	} else if statusCode != 200 {
		pterm.Warning.Printf("Returned with status: %s\n", status)
	}
}

func getRequestData(resp *http.Response) []byte {
	checkStatusCode(resp.StatusCode, resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	return body
}
