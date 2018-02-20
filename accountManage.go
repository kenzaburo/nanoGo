package nanoGo

import (
	"strconv"
)

// AccountCreate : creates a new account in the given wallet. "work" toggles work generation after creating account and works on node version 8.1 onwards. "enable_control" has to be enabled in the config of the node
func (client *Client) AccountCreate(wallet string, work bool) (account string, err error) {
	var response map[string]string

	err = client.Request("account_create", map[string]interface{}{"wallet": wallet, "work": strconv.FormatBool(work)}, &response)

	return response["account"], err
}

// AccountList : gets all accounts inside a wallet
func (client *Client) AccountList(wallet string) (accounts []string, err error) {
	var response map[string][]string

	err = client.Request("account_list", map[string]interface{}{"wallet": wallet}, &response)

	return response["accounts"], err
}

// AccountMove : moves the given accounts from the "source" to the "target" wallet. "enable_control" has to be enabled in the config of the node
func (client *Client) AccountMove(target string, source string, accounts ...string) (success bool, err error) {
	var response map[string]string

	err = client.Request("account_move", map[string]interface{}{"wallet": target, "source": source,
		"accounts": accounts}, &response)

	if response["moved"] == "1" {
		success = true
	}

	return success, err
}

// AccountRemove : removes an account from the given wallet
func (client *Client) AccountRemove(wallet string, account string) (success bool, err error) {
	var response map[string]string

	err = client.Request("account_remove", map[string]interface{}{"wallet": wallet, "account": account}, &response)

	if response["removed"] == "1" {
		success = true
	}

	return success, err
}

// AccountRepresentativeSet : sets the representative of an account. "work" and the value for a block from an external source and works with node version 8.1 or newer (set "" to disable). "enable_control" has to be enabled in the config of the node
func (client *Client) AccountRepresentativeSet(wallet string, account string, representative string, work string) (block string, err error) {
	var response map[string]string

	args := map[string]interface{}{"wallet": wallet, "account": account, "representative": representative}

	if work != "" {
		args["work"] = work
	}

	err = client.Request("account_representative_set", args, &response)

	return response["block"], err
}

// AccountsCreate : creates multiple new accounts in the given wallet. Works on node version 8.1 onwards. "work" toggles work generation after creating accounts. "enable_control" has to be enabled in the config of the node
func (client *Client) AccountsCreate(wallet string, count int, work bool) (accounts []string, err error) {
	var response map[string][]string

	err = client.Request("accounts_create", map[string]interface{}{"wallet": wallet, "count": strconv.Itoa(count),
		"work": strconv.FormatBool(work)}, &response)

	return response["accounts"], err
}
