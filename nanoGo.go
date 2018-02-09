package nanoGo

import (
	"strconv"
)

// AccountInfo : information about requested account
type AccountInfo struct {
	Frontier            string
	OpenBlock           string  `json:"open_block"`
	RepresentativeBlock string  `json:"representative_block"`
	Balance             float64 `json:",string"`
	ModifiedTimestamp   int     `json:"modified_timestamp,string"`
	BlockCount          float64 `json:"block_count,string"`
	Representative      string
	Weight              float64 `json:",string"`
	Pending             float64 `json:",string"`
}

// History : entry of the thransaction of an account
type History struct {
	Hash    string
	Type    string
	Account string
	Amount  float64 `json:",string"`
}

// AccountBalance : gets the balance of an account
func (client *Client) AccountBalance(account string) (balance float64, pending float64, err error) {
	var response AccountInfo

	err = client.Request("account_balance", map[string]string{"account": account}, &response)

	return response.Balance, response.Pending, err
}

// AccountBlockCount : gets the block count of an account
func (client *Client) AccountBlockCount(account string) (blockCount float64, err error) {
	var response AccountInfo

	err = client.Request("account_block_count", map[string]string{"account": account}, &response)

	return response.BlockCount, err
}

type historyResponse struct {
	History []History
}

// AccountHistory : gets the history of an account. Returns an array of the struct "History"
func (client *Client) AccountHistory(account string, count int) (history []History, err error) {
	var response historyResponse

	err = client.Request("account_history", map[string]string{"account": account, "count": strconv.Itoa(count)}, &response)

	if err == nil {
		history = response.History
	}

	return history, err
}

// AccountInfo : gets frontier, open block, change representative block, balance, last modified timestamp and block count of an account. Set legacy true if your node is older than version 8.1 or you don't need the representative, weight and the pending balance of the account
func (client *Client) AccountInfo(account string, legacy bool) (response AccountInfo, err error) {
	args := map[string]string{"account": account}

	if !legacy {
		args["representative"] = "true"
		args["weight"] = "true"
		args["pending"] = "true"
	}

	err = client.Request("account_info", args, &response)

	return response, err
}

// AccountKey : gets the public key of an account
func (client *Client) AccountKey(account string) (key string, err error) {
	var response map[string]string

	err = client.Request("account_key", map[string]string{"account": account}, &response)

	return response["key"], err
}

// AccountRepresentative : gets the representative of an account
func (client *Client) AccountRepresentative(account string) (representative string, err error) {
	var response AccountInfo

	err = client.Request("account_representative", map[string]string{"account": account}, &response)

	return response.Representative, err
}

// AccountWeight : gets the weight of an account
func (client *Client) AccountWeight(account string) (weight float64, err error) {
	var response AccountInfo

	err = client.Request("account_weight", map[string]string{"account": account}, &response)

	return response.Weight, err
}

// ValidateAccountNumber : checks wether an account is existing or not
func (client *Client) ValidateAccountNumber(account string) (valid bool, err error) {
	var response map[string]string

	err = client.Request("validate_account_number", map[string]string{"account": account}, &response)

	if err == nil {
		if response["valid"] == "1" {
			valid = true
		}
	}

	return valid, err
}
