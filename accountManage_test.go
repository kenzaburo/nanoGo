package nanoGo_test

import (
	"fmt"
	"testing"
	"time"
)

const wallet = "2E221DE7A48B6730F1E083464F732075CF68E02355245DF2E019FFA716D02977"

const nanodeRepresentative = "xrb_1nanode8ngaakzbck8smq6ru9bethqwyehomf79sae1k7xd47dkidjqzffeg"

func TestAccountAddRemove(t *testing.T) {
	newAccount, err := client.AccountCreate(wallet, false)

	if !accountExists(newAccount) || err != nil {
		failTest(t, false, true, err, "Creatig new accounts fails")
	}

	time.Sleep(wait)

	removed, err := client.AccountRemove(wallet, newAccount)

	if !removed || accountExists(newAccount) || err != nil {
		failTest(t, true, false, err, "Removing accounts fails")
	}

	time.Sleep(wait)
}

func TestAccountList(t *testing.T) {
	list, err := client.AccountList(wallet)

	if len(list) != 1 || list[0] != testAccount2 || err != nil {
		failTest(t, []string{testAccount2}, list, err, "Getting accounts of wallet returns wrong results")
	}

	time.Sleep(wait)
}

// TODO: test moving account when creating new wallet is implemented
func TestAccountMove(t *testing.T) {
	moved, err := client.AccountMove(wallet, wallet, testAccount2)

	if moved || !accountExists(testAccount2) || err != nil {
		failTest(t, true, moved, err, "Moving accounts to the same wallet as moved from does not fail")
	}

	time.Sleep(wait)

	moved, err = client.AccountMove("asdf", "asdf", testAccount2)

	if moved || fmt.Sprint(err) != "Bad wallet number" {
		failTest(t, true, moved, err, "Moving accounts does not show an error")
	}

	time.Sleep(wait)
}

func TestAccountRepresentativeSet(t *testing.T) {
	newAccount, _ := client.AccountCreate(wallet, true)

	time.Sleep(wait)

	block, err := client.AccountRepresentativeSet(wallet, newAccount, nanodeRepresentative, "")

	// Cannot check the representative because the block has to be synced to be shown
	if block == "" || err != nil {
		failTest(t, "Block hash", block, err, "Setting representative of account fails")
	}

	time.Sleep(wait)

	removed, err := client.AccountRemove(wallet, newAccount)

	if !removed || err != nil {
		failTest(t, nil, err, err, "Removing accounts fails")
	}

	time.Sleep(wait)
}

func TestAccountsCreate(t *testing.T) {
	newAccounts, err := client.AccountsCreate(wallet, 2, false)

	if err != nil {
		failTest(t, "New accounts", newAccounts, err, "Creating multiple accounts fails")
	}

	for _, newAccount := range newAccounts {
		if !accountExists(newAccount) {
			failTest(t, true, false, err, "Creating multiple accounts fails")

		} else {
			time.Sleep(wait)

			removed, err := client.AccountRemove(wallet, newAccount)

			if !removed || err != nil {
				failTest(t, true, false, err, "Removing accounts fails")
			}

		}

	}

	time.Sleep(wait)
}

func accountExists(search string) bool {
	time.Sleep(wait)

	accounts, _ := client.AccountList(wallet)

	for _, account := range accounts {
		if search == account {
			return true
		}

	}

	return false
}
