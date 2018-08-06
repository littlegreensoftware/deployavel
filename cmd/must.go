package cmd

import (
	"bytes"
	"encoding/json"
	"log"
)

// Must is a generic error handler
func Must(data []byte, err error) []byte {
	if err != nil {
		log.Fatal(err)
	}

	return data
}

// PrintMust logs data to the console
func PrintMust(data []byte) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(prettyJSON.Bytes()))
}
