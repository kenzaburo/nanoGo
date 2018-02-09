package nanoGo_test

import (
	"github.com/michael1011/nanoGo"
	"strconv"
	"testing"
)

func TestRequest(t *testing.T) {
	client := nanoGo.Connect(testServer)

	var object map[string]string

	err := client.Request("account_balance", map[string]string{"account": testAccount}, &object)

	if err != nil {
		t.Log(err)

		t.Fatal("Failed to send request to RPC server")
	}

	balance, _ := strconv.ParseFloat(object["balance"], 64)

	if balance != testAccountInfo.Balance {
		t.Log(object)
		t.Log("Balance should be: ", testAccountInfo.Balance)

		t.Fatal("Requests don't get parsed correctly")
	}
}