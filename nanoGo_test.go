package nanoGo_test

import (
	"github.com/michael1011/nanoGo"
	"testing"
)

// A Nano node must be running local on the default port the execute this tests
const testServer = "http://[::1]:7076"

// Faucet account; let's hope it won't get used anymore
const testAccount = "xrb_35jjmmmh81kydepzeuf9oec8hzkay7msr6yxagzxpcht7thwa5bus5tomgz9"

const testAccountKey = "8E319CE6F3025E5B2DF66DA7AB1467FE48F1679C13DD43BFDB29FA2E9FC40D3B"

var testAccountInfo = nanoGo.AccountInfo{
	Frontier:            "C529AB93A289F8F89B964F4C970D9752089A4156E4C44F70761449B597997BDF",
	OpenBlock:           "4D0DBC10D672F39BCBE867A8F6ED41F8848D18D0BDDA16375B98749B1A321F06",
	RepresentativeBlock: "4D0DBC10D672F39BCBE867A8F6ED41F8848D18D0BDDA16375B98749B1A321F06",
	Balance:             795055344175165130955846320127,
	ModifiedTimestamp:   1510197057,
	BlockCount:          111435,
	Representative:      "xrb_3arg3asgtigae3xckabaaewkx3bzsh7nwz7jkmjos79ihyaxwphhm6qgjps4",
	Weight:              200465936329999988156571674738688,
	Pending:             0,
}

var testAccountHistory = []nanoGo.History{
	{
		Hash:    "C529AB93A289F8F89B964F4C970D9752089A4156E4C44F70761449B597997BDF",
		Type:    "send",
		Account: "xrb_1111111111111111111111111111111111111111111111111111hifc8npp",
		Amount:  118657000000000000344933521162240000,
	},
}

var client = nanoGo.Connect(testServer)

func TestAccountBalance(t *testing.T) {
	balance, pending, err := client.AccountBalance(testAccount)

	if balance != testAccountInfo.Balance || pending != 0 || err != nil {
		failTest(t, testAccountInfo.Balance, balance, err, "Getting balance of an account fails")
	}

}

func TestAccountBlockCount(t *testing.T) {
	blockCount, err := client.AccountBlockCount(testAccount)

	if blockCount != testAccountInfo.BlockCount || err != nil {
		failTest(t, testAccountInfo.BlockCount, blockCount, err, "Getting block count of an account fails")
	}

}

func TestAccountHistory(t *testing.T) {
	size := 1

	history, err := client.AccountHistory(testAccount, size)

	if len(history) != size || err != nil {
		failTest(t, size, len(history), err, "Incorrect number of history results are returned")
	}

	if history[0] != testAccountHistory[0] {
		failTest(t, testAccountHistory, history, err, "Returned history is not correct")
	}

	size = 10

	history, err = client.AccountHistory(testAccount, size)

	if len(history) != size || err != nil {
		failTest(t, size, len(history), err, "Incorrect number of history results are returned")
	}
}

func TestAccountInfo(t *testing.T) {
	accountInfo, err := client.AccountInfo(testAccount, false)

	if accountInfo != testAccountInfo || err != nil {
		failTest(t, testAccountInfo, accountInfo, err, "Getting account info fails")
	}

}

func TestAccountKey(t *testing.T) {
	key, err := client.AccountKey(testAccount)

	if key != testAccountKey || err != nil {
		failTest(t, testAccountKey, key, err, "Getting the public key of an account fails")
	}
}

func TestAccountRepresentative(t *testing.T) {
	representative, err := client.AccountRepresentative(testAccount)

	if representative != testAccountInfo.Representative || err != nil {
		failTest(t, testAccountInfo.Representative, representative, err, "Getting representative of an account fails")
	}

}

func TestAccountWeight(t *testing.T) {
	weight, err := client.AccountWeight(testAccount)

	if weight != testAccountInfo.Weight || err != nil {
		failTest(t, testAccountInfo.Weight, weight, err, "Getting weight of an account fails")
	}

}

func TestValidateAccountNumber(t *testing.T) {
	// Valid account
	valid, err := client.ValidateAccountNumber(testAccount)

	if valid == false || err != nil {
		failTest(t, true, valid, err, "Valid accounts seem to be invalid")
	}

	// Invalid account
	valid, err = client.ValidateAccountNumber("xrb_3e3j5tkog48pnny9dmfzj1r16pg8t1e76dz5tmac6iq689wyjfpi00000000")

	if valid == true || err != nil {
		failTest(t, false, valid, err, "Invalid accounts seem to be valid")
	}

}

func failTest(t *testing.T, expected interface{}, received interface{}, err error, messages string) {
	t.Log("Expected:", expected)
	t.Log("Received:", received)

	t.Log()

	t.Log(err)
	t.Fatal(messages)
}
