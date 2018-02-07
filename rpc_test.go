package nanoGo_test

import (
	"github.com/michael1011/nanoGo"
	"testing"
)

// A Nano node must be running local on the default port the execute this tests
const testServer = "http://[::1]:7076"

func TestConnect(t *testing.T) {
	err := nanoGo.Connect(testServer)

	if err != nil {
		t.Log(err)
		t.Fatal("Connecting to working node fails")
	}
}
