package nanoGo_test

import (
	"fmt"
	"github.com/michael1011/nanoGo"
	"testing"
	"time"
)

// Time between two request to the RPC server
const wait = time.Millisecond * 25

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
	Pending:             999999999999999983222784,
}

var testAccountBalance = nanoGo.Balance{
	Balance: testAccountInfo.Balance,
	Pending: testAccountInfo.Pending,
}

var testAccountHistory = []nanoGo.History{
	{
		Hash:    "C529AB93A289F8F89B964F4C970D9752089A4156E4C44F70761449B597997BDF",
		Type:    "send",
		Account: "xrb_1111111111111111111111111111111111111111111111111111hifc8npp",
		Amount:  118657000000000000344933521162240000,
	},
}

const testAccountPendingBlockHash = "66A5163AA3900C55B1DD3E671A33CC9060F19F8DFBFBBE7218105BB08B47E5FE"

var testAccountPendingBlock = nanoGo.PendingBlock{
	Amount: testAccountInfo.Pending,
	Source: testAccount2,
}

const testAccount2 = "xrb_1ep8uydu4s6xfq8wejjsh1c7gzmrzgqs3g1mugnu5dt7yjhmape3pskxxage"

var testAccountsBalances = map[string]nanoGo.Balance{
	testAccount:  testAccountBalance,
	testAccount2: {8219126153845999694512128, 0},
}

var testAccountsFrontiers = map[string]string{
	testAccount:  testAccountInfo.Frontier,
	testAccount2: "66A5163AA3900C55B1DD3E671A33CC9060F19F8DFBFBBE7218105BB08B47E5FE",
}

var client = nanoGo.Connect(testServer)

func TestAccountBalance(t *testing.T) {
	balance, err := client.AccountBalance(testAccount)

	if balance != testAccountBalance || err != nil {
		failTest(t, testAccountBalance, balance, err, "Getting balance of an account fails")
	}

	time.Sleep(wait)
}

func TestAccountBlockCount(t *testing.T) {
	blockCount, err := client.AccountBlockCount(testAccount)

	if blockCount != testAccountInfo.BlockCount || err != nil {
		failTest(t, testAccountInfo.BlockCount, blockCount, err, "Getting block count of an account fails")
	}

	time.Sleep(wait)
}

func TestAccountGet(t *testing.T) {
	account, err := client.AccountGet(testAccountKey)

	if account != testAccount || err != nil {
		failTest(t, testAccount, account, err, "Getting the account of an public key fails")
	}

	time.Sleep(wait)
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

	time.Sleep(wait)

	size = 10

	history, err = client.AccountHistory(testAccount, size)

	if len(history) != size || err != nil {
		failTest(t, size, len(history), err, "Incorrect number of history results are returned")
	}

	time.Sleep(wait)
}

func TestAccountInfo(t *testing.T) {
	accountInfo, err := client.AccountInfo(testAccount, false)

	if accountInfo != testAccountInfo || err != nil {
		failTest(t, testAccountInfo, accountInfo, err, "Getting account info fails")
	}

	time.Sleep(wait)
}

func TestAccountKey(t *testing.T) {
	key, err := client.AccountKey(testAccount)

	if key != testAccountKey || err != nil {
		failTest(t, testAccountKey, key, err, "Getting the public key of an account fails")
	}

	time.Sleep(wait)
}

func TestAccountRepresentative(t *testing.T) {
	representative, err := client.AccountRepresentative(testAccount)

	if representative != testAccountInfo.Representative || err != nil {
		failTest(t, testAccountInfo.Representative, representative, err, "Getting representative of an account fails")
	}

	time.Sleep(wait)
}

func TestAccountWeight(t *testing.T) {
	weight, err := client.AccountWeight(testAccount)

	if weight != testAccountInfo.Weight || err != nil {
		failTest(t, testAccountInfo.Weight, weight, err, "Getting weight of an account fails")
	}

	time.Sleep(wait)
}

func TestAccountsBalances(t *testing.T) {
	balances, err := client.AccountsBalances(testAccount, testAccount2)

	if balances[testAccount] != testAccountsBalances[testAccount] ||
		balances[testAccount2] != testAccountsBalances[testAccount2] || err != nil {

		failTest(t, testAccountsBalances, balances, err, "Getting balances of multiple accounts fails")
	}

	time.Sleep(wait)
}

func TestAccountsFrontiers(t *testing.T) {
	frontiers, err := client.AccountsFrontiers(testAccount, testAccount2)

	if frontiers[testAccount] != testAccountsFrontiers[testAccount] ||
		frontiers[testAccount2] != testAccountsFrontiers[testAccount2] || err != nil {

		failTest(t, testAccountsFrontiers, frontiers, err, "Getting frontiers of multiple accounts fails")
	}

	time.Sleep(wait)
}

func TestAccountsPending(t *testing.T) {
	// Getting the pending block hash
	pending, err := client.AccountsPending(0, false, 1, testAccount)

	_, hash := pending[testAccount][testAccountPendingBlockHash]

	if !hash || err != nil {
		failTest(t, "One pending block", pending, err, "Getting pending blocks without threshold and source fails")
	}

	time.Sleep(wait)

	// Getting pending blocks of an account without pending blocks
	pending, err = client.AccountsPending(0, false, 1, testAccount2)

	if len(pending[testAccount2]) != 0 || fmt.Sprint(err) !=
		"No pending blocks were found for accounts: "+testAccount2 {

		failTest(t, map[string]nanoGo.PendingBlock{}, pending[testAccount2], err,
			"Getting pending blocks when there are none returns wrong error")
	}

	time.Sleep(wait)

	// Getting pending blocks with threshold where the threshold is lower than the amount of the pending block
	pending, err = client.AccountsPending(1, false, 1, testAccount)

	if len(pending[testAccount]) != 1 || err != nil {
		failTest(t, "One pending block", pending, err,
			"Getting pending blocks with threshold returns nothing although there should be one pending block")
	}

	time.Sleep(wait)

	// Getting pending blocks with threshold where the threshold is larger than the amount of the pending block
	pending, err = client.AccountsPending(testAccountInfo.Pending+1e+10, false, 1, testAccount)

	if len(pending[testAccount2]) != 0 || fmt.Sprint(err) !=
		"No pending blocks were found for accounts: "+testAccount {

		failTest(t, map[string]nanoGo.PendingBlock{}, pending, err,
			"Getting pending blocks with a threshold returns blocks under the threshold")
	}

	time.Sleep(wait)

	// Getting pending blocks with source
	pending, err = client.AccountsPending(0, true, 1, testAccount)

	if pending[testAccount][testAccountPendingBlockHash] != testAccountPendingBlock || err != nil {
		failTest(t, testAccountPendingBlock, pending[testAccount][testAccountPendingBlockHash],
			err, "Getting pending blocks with source fails")
	}

	time.Sleep(wait)

	// Getting pending blocks with source of an account without pending blocks
	pending, err = client.AccountsPending(0, true, 1, testAccount2)

	if len(pending[testAccount2]) != 0 || fmt.Sprint(err) !=
		"No pending blocks were found for accounts: "+testAccount2 {

		failTest(t, map[string]nanoGo.PendingBlock{}, pending[testAccount2], err,
			"Getting pending blocks with source when there are none returns wrong error")
	}

	time.Sleep(wait)
}

func TestValidateAccountNumber(t *testing.T) {
	// Valid account
	valid, err := client.ValidateAccountNumber(testAccount)

	if valid == false || err != nil {
		failTest(t, true, valid, err, "Valid accounts seem to be invalid")
	}

	time.Sleep(wait)

	// Invalid account
	valid, err = client.ValidateAccountNumber("xrb_3e3j5tkog48pnny9dmfzj1r16pg8t1e76dz5tmac6iq689wyjfpi00000000")

	if valid == true || err != nil {
		failTest(t, false, valid, err, "Invalid accounts seem to be valid")
	}

	time.Sleep(wait)
}

func failTest(t *testing.T, expected interface{}, received interface{}, err error, messages string) {
	t.Log("Expected:", expected)
	t.Log("Received:", received)

	t.Log()

	t.Log(err)
	t.Fatal(messages)
}
