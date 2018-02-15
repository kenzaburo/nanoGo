package nanoGo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

// Balance : balance an pending Nano of an account
type Balance struct {
	Balance float64 `json:",string"`
	Pending float64 `json:",string"`
}

// History : entry of the thransaction of an account
type History struct {
	Hash    string
	Type    string
	Account string
	Amount  float64 `json:",string"`
}

// PendingBlock : amount and source account of an pending blockCount
type PendingBlock struct {
	Amount float64 `json:",string"`
	Source string
}

// AccountBalance : gets the balance of an account
func (client *Client) AccountBalance(account string) (balance Balance, err error) {
	var response Balance

	err = client.Request("account_balance", map[string]interface{}{"account": account}, &response)

	return response, err
}

// AccountBlockCount : gets the block count of an account
func (client *Client) AccountBlockCount(account string) (blockCount float64, err error) {
	var response AccountInfo

	err = client.Request("account_block_count", map[string]interface{}{"account": account}, &response)

	return response.BlockCount, err
}

// AccountGet : gets the account of a public key
func (client *Client) AccountGet(key string) (account string, err error) {
	var response map[string]string

	err = client.Request("account_get", map[string]interface{}{"key": key}, &response)

	return response["account"], err
}

// AccountHistory : gets the history of an account. Returns an array of the struct "History"
func (client *Client) AccountHistory(account string, count int) (history []History, err error) {
	var response map[string][]History

	err = client.Request("account_history", map[string]interface{}{"account": account, "count": strconv.Itoa(count)}, &response)

	if err == nil {
		history = response["history"]
	}

	return history, err
}

// AccountInfo : gets frontier, open block, change representative block, balance, last modified timestamp and block count of an account. Set legacy true if your node is older than version 8.1 or you don't need the representative, weight and the pending balance of the account
func (client *Client) AccountInfo(account string, legacy bool) (response AccountInfo, err error) {
	args := map[string]interface{}{"account": account}

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

	err = client.Request("account_key", map[string]interface{}{"account": account}, &response)

	return response["key"], err
}

// AccountRepresentative : gets the representative of an account
func (client *Client) AccountRepresentative(account string) (representative string, err error) {
	var response AccountInfo

	err = client.Request("account_representative", map[string]interface{}{"account": account}, &response)

	return response.Representative, err
}

// AccountWeight : gets the weight of an account
func (client *Client) AccountWeight(account string) (weight float64, err error) {
	var response AccountInfo

	err = client.Request("account_weight", map[string]interface{}{"account": account}, &response)

	return response.Weight, err
}

// AccountsBalances : gets the balances of multiple accounts
func (client *Client) AccountsBalances(accounts ...string) (balances map[string]Balance, err error) {
	var response map[string]map[string]Balance

	err = client.Request("accounts_balances", map[string]interface{}{"accounts": accounts}, &response)

	return response["balances"], err
}

// AccountsFrontiers : gets frontiers of multiple accounts
func (client *Client) AccountsFrontiers(accounts ...string) (frontiers map[string]string, err error) {
	var response map[string]map[string]string

	err = client.Request("accounts_frontiers", map[string]interface{}{"accounts": accounts}, &response)

	return response["frontiers"], err
}

// AccountsPending : gets pending blocks of multiple accounts. "count" specifies the number of retrieved blocks. "threshold" is the minimum pending amount of an block and works only on node versions above or equal to 8.0 (set 0 to disable). "source" adds the source accounts of the blocks to the response and works only on node versions above or equal to 8.1 (set false to disable)
func (client *Client) AccountsPending(threshold float64, source bool, count float64, accounts ...string) (pending map[string]map[string]PendingBlock, err error) {
	errorMessage := "No pending blocks were found for accounts: "

	action := "accounts_pending"
	args := map[string]interface{}{"accounts": accounts, "count": count}

	pending = map[string]map[string]PendingBlock{}

	if threshold == 0 && !source {
		var response map[string]map[string][]string

		err = client.Request(action, args, &response)

		if err == nil || fmt.Sprint(err) == "json: cannot unmarshal string into Go value of type []string" {
			for account, blocks := range response["blocks"] {
				blocksMap := map[string]PendingBlock{}

				for _, block := range blocks {
					blocksMap[block] = PendingBlock{}
				}

				// No pending blocks
				if len(blocksMap) == 0 {
					errorMessage += account + ", "
				}

				pending[account] = blocksMap

			}

		}

	} else if threshold != 0 && !source {
		var response map[string]map[string]map[string]string

		args["threshold"] = strconv.FormatFloat(threshold, 'f', -1, 64)

		err = client.Request(action, args, &response)

		if err == nil || fmt.Sprint(err) == "json: cannot unmarshal string into Go value of type map[string]string" {
			for account, blocks := range response["blocks"] {
				blocksMap := map[string]PendingBlock{}

				for block, rawAmount := range blocks {
					amount, _ := strconv.ParseFloat(rawAmount, 64)

					blocksMap[block] = PendingBlock{
						Amount: amount,
					}

				}

				// No pending blocks
				if len(blocksMap) == 0 {
					errorMessage += account + ", "
				}

				pending[account] = blocksMap

			}

		}

	} else {
		var response map[string]map[string]map[string]PendingBlock

		// If "source" is working "threshold" is too
		args["threshold"] = strconv.FormatFloat(threshold, 'f', -1, 64)
		args["source"] = strconv.FormatBool(source)

		err = client.Request(action, args, &response)

		if err == nil || fmt.Sprint(err) == "json: cannot unmarshal string into Go value of type map[string]nanoGo.PendingBlock" {
			pending = response["blocks"]

			for account, blocks := range pending {
				// No pending blocks
				if len(blocks) == 0 {
					errorMessage += account + ", "
				}

			}

		}

	}

	if err != nil {
		errorMessage = strings.TrimSuffix(errorMessage, ", ")

		err = errors.New(errorMessage)
	}

	return pending, err
}

// ValidateAccountNumber : checks wether an account is existing or not
func (client *Client) ValidateAccountNumber(account string) (valid bool, err error) {
	var response map[string]string

	err = client.Request("validate_account_number", map[string]interface{}{"account": account}, &response)

	if err == nil {
		if response["valid"] == "1" {
			valid = true
		}
	}

	return valid, err
}
