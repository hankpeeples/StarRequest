package jsonformatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"star-request/utils"
)

// Format - prints the body of the request
func Format(body []byte) {
	log := utils.GetLogger()

	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		log.Errorf("Error formatting JSON: %v", err)
	}

	fmt.Println(string(prettyJSON.String()))
}
