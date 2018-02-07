package nanoGo

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// RPC protocol documentation: https://github.com/clemahieu/raiblocks/wiki/RPC-protocol

var client = &http.Client{}

var address string

// Connect : send a sample request to check if the server is responding
func Connect(address string) (err error) {
	address = address

	_, err = sendRequest("version", nil)

	return err
}

func sendRequest(action string, args map[string]string) (response []byte, err error) {
	jsonBody := "{\"action\": \"" + action + "\""

	for key, value := range args {
		jsonBody += ","

		jsonBody += "\"" + key + "\": \"" + value + "\""
	}

	jsonBody += "}"

	req, err := http.NewRequest("POST", address, bytes.NewBuffer([]byte(jsonBody)))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return body, err
}
