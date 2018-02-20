package nanoGo_test

import (
	"fmt"
	"github.com/michael1011/nanoGo"
	"strconv"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	client := nanoGo.Connect(testServer)

	var object map[string]string

	err := client.Request("account_balance", map[string]interface{}{"account": testAccount}, &object)

	if err != nil {
		t.Log(err)

		t.Fatal("Failed to send request to RPC server")
	}

	time.Sleep(wait)

	balance, _ := strconv.ParseFloat(object["balance"], 64)

	if balance != testAccountInfo.Balance {
		t.Log(object)
		t.Log("Balance should be: ", testAccountInfo.Balance)

		t.Fatal("Requests don't get parsed correctly")
	}

	time.Sleep(wait)
}

func TestRequestErrorHandling(t *testing.T) {
	command := "notFound"
	args := map[string]interface{}{}

	err := fmt.Sprint(client.Request(command, args, &map[string]string{}))

	if err != "Unknown command" {
		t.Log(err)
		t.Fatal("Parsing error when the object is a \"map[string]string\" fails")
	}

	time.Sleep(wait)

	err2 := fmt.Sprint(client.Request(command, args, &map[string]interface{}{}))

	if err2 != err {
		t.Log(err2)
		t.Fatal("Parsing error when the object is not a \"map[string]string\" fails")
	}

}
