package nanoGo

import (
	"encoding/json"
	"testing"
)

// A Nano node must be running local on the default port the execute this tests
const testServer = "http://[::1]:7076"

// Faucet account; let's hope it won't get used anymore
const testAccount = "xrb_35jjmmmh81kydepzeuf9oec8hzkay7msr6yxagzxpcht7thwa5bus5tomgz9"

const testAccountBalance = "795055344175165130955846320127"

type balance struct {
	Balance string
	Pending string
}

func TestSendRequest(t *testing.T) {
	address = testServer

	_, err := sendRequest("version", nil)

	if err != nil {
		t.Log(err)
		t.Fatal("Sending requests fails")
	}

	data, err := sendRequest("account_balance", map[string]string{"account": testAccount})

	var response balance

	err = json.Unmarshal(data, &response)

	if response.Balance != testAccountBalance {
		t.Log(string(data))
		t.Log("Balance should be: " + testAccountBalance)
		t.Fatal("Sending request with parameters fails")
	}
}
