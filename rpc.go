package nanoGo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RPC protocol documentation: https://github.com/clemahieu/raiblocks/wiki/RPC-protocol

// Client : Nano RPC client
type Client struct {
	address string
}

var httpClient = &http.Client{}

// Connect : creates a new Client
func Connect(address string) *Client {
	client := &Client{
		address: address,
	}

	return client
}

// Request : sends a custom request to the RPC server. "object" is the interface for parsing
func (client *Client) Request(action string, args map[string]interface{}, object interface{}) (err error) {
	response, err := client.sendRequest(action, args)

	if err != nil {
		return err
	}

	err = parseRequest(response, &object)

	return err
}

func parseRequest(data []byte, object interface{}) (err error) {
	return json.Unmarshal(data, &object)
}

func (client *Client) sendRequest(action string, args map[string]interface{}) (response []byte, err error) {
	args["action"] = action

	jsonBody, err := json.Marshal(args)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.address, bytes.NewBuffer(jsonBody))

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
