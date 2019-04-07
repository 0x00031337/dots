package dots

import (
	"fmt"
	"github.com/fraudmarc/go-monero/walletrpc"
	"log"
	"os"
	"strings"
	"sync"
)

func GetAccountName(account walletrpc.SubaddressAccount) (name string) {
	if account.Label != "" {
		return account.Label
	}
	return GetAccountNameComplex(account)
}

func GetAccountNameComplex(account walletrpc.SubaddressAccount) (name string) {
	if account.Label == "" {
		return string(account.AccountIndex) +
			":" +
			Abbreviate(account.BaseAddress)
	}
	return account.Label +
		":" +
		Abbreviate(account.BaseAddress)
}

func ProcessAccount(account walletrpc.SubaddressAccount, client walletrpc.Client) error {
	ShowAccount(account, client)

	var err error
	switch {
	case strings.HasSuffix(account.Label, config.SafeSuffix):
		fmt.Printf("Skip -safe account\n")
	case strings.HasPrefix(account.Label, config.PendingPrefix):
		fmt.Printf("Found pending account\n")
		// continue processing the pending account
		// it should only have a single utxo
	case account.Label == "":
		// create -safe account from index/base if it doesnt exist
		_, err = client.CreateAccount(string(account.AccountIndex) + config.SafeSuffix)
	default:
		// create -safe account from label if it doesnt exist
		_, err = client.CreateAccount(account.Label + config.SafeSuffix)
	}
	if err != nil {
		if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
			// it is a monero wallet error
			fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
			os.Exit(1)
		}
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Processing Account")
	fmt.Printf("  AccountIndex: %d\n", account.AccountIndex)
	fmt.Printf("  BaseAddress: %s\n", Abbreviate(account.BaseAddress))
	fmt.Printf("  Balance : %s\n", walletrpc.XMRToDecimal(account.Balance))
	fmt.Printf("  Unlocked: %s\n", walletrpc.XMRToDecimal(account.UnlockedBalance))
	fmt.Printf("  Label: %s\n", account.Label)
	fmt.Printf("  Tag: %s\n", account.Tag)

	fmt.Println("  Getting transfers")

	transfers, err := client.IncomingTransfers(walletrpc.TransferAll, account.AccountIndex)
	if err != nil {
		if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
			// it is a monero wallet error
			fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
			os.Exit(1)
		}
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	for _, transfer := range transfers {
		ShowTransfer(transfer)
	}

	return nil
}

func ShowAccount(account walletrpc.SubaddressAccount, client walletrpc.Client) error {
	fmt.Println("Showing Account")
	fmt.Printf("  AccountIndex: %d\n", account.AccountIndex)
	fmt.Printf("  BaseAddress: %s\n", Abbreviate(account.BaseAddress))
	fmt.Printf("  Balance : %s\n", walletrpc.XMRToDecimal(account.Balance))
	fmt.Printf("  Unlocked: %s\n", walletrpc.XMRToDecimal(account.UnlockedBalance))
	fmt.Printf("  Label: %s\n", account.Label)
	fmt.Printf("  Tag: %s\n", account.Tag)

	fmt.Println("  Getting transfers")

	transfers, err := client.IncomingTransfers(walletrpc.TransferAll, account.AccountIndex)
	if err != nil {
		if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
			// it is a monero wallet error
			fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
			os.Exit(1)
		}
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	for _, transfer := range transfers {
		ShowTransfer(transfer)
	}

	return nil
}

func ProcessAccounts(accounts []walletrpc.SubaddressAccount, client walletrpc.Client) error {
	//var err error
	//var accountMap map[string]walletrpc.SubaddressAccount
	var wg sync.WaitGroup
	accountMap := make(map[string]walletrpc.SubaddressAccount)
	blacklistMap := make(map[string]bool)

	pendingMap := make(map[string]walletrpc.SubaddressAccount)
	safeMap := make(map[string]walletrpc.SubaddressAccount)
	// so utxos know where to end up
	safeNameMap := make(map[uint64]string)

	// first, map the account names
	for _, account := range accounts {

		accountName := GetAccountName(account)
		fmt.Println(accountName)

		if isSafe(accountName) {
			if _, ok := safeMap[accountName]; ok {
				//panic("Duplicate -safe accounts")
				log.Println("WARN: ignoring duplicate safe account")
			}
			safeMap[accountName] = account
			continue
		}

		if isPending(accountName) {
			if _, ok := pendingMap[accountName]; ok {
				//panic("Duplicate -pending accounts")
				log.Println("WARN: ignoring duplicate pending account")
			}
			pendingMap[accountName] = account
			continue
		}

		if usedAccount, ok := accountMap[accountName]; ok {
			// account name already used, rename this one and the existing one
			accountMap[GetAccountNameComplex(usedAccount)] = usedAccount
			blacklistMap[accountName] = true
			accountMap[GetAccountNameComplex(account)] = account
		} else {
			// first time seeing this account name
			accountMap[accountName] = account
		}

	}

	fmt.Println("Create new accounts")

	// create new -safe target accounts where necessary
	// based on regular accounts
	// do this in order so -safe will be in front of -pending in wallet account list
	for name, value := range accountMap {
		newNameSafe := nameSafe(name)
		// check if there's already a -safe version of this account
		// if not, make one
		if _, ok := safeMap[newNameSafe]; !ok {
			newAccountResponse, err := client.CreateAccount(newNameSafe)
			if err != nil {
				if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
					// it is a monero wallet error
					fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
					os.Exit(1)
				}
				fmt.Println("Error:", err.Error())
				os.Exit(1)
			}
			newAccount := walletrpc.SubaddressAccount{
				AccountIndex: newAccountResponse.AccountIndex,
				Label:        newNameSafe,
				BaseAddress:  newAccountResponse.Address,
			}
			safeMap[newNameSafe] = newAccount
		}

		// so we can find the beginning account's corresponding -safe account destination
		safeNameMap[value.AccountIndex] = newNameSafe
	}

	// create new -pending target accounts
	// iterate unspent transactions in each account
	for name, account := range accountMap {
		// get utxo for account

		transfers, err := client.IncomingTransfers(walletrpc.TransferAvailable, account.AccountIndex)
		if err != nil {
			if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
				// it is a monero wallet error
				fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
				os.Exit(1)
			}
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
		for _, transfer := range transfers {
			newNamePending := namePending(name, transfer.TxHash)
			// check if there's already a pending version of this account
			// if not, make one
			if _, ok := pendingMap[newNamePending]; !ok {
				newAccountResponse, err := client.CreateAccount(newNamePending)
				if err != nil {
					if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
						// it is a monero wallet error
						fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
						os.Exit(1)
					}
					fmt.Println("Error:", err.Error())
					os.Exit(1)
				}
				newAccount := walletrpc.SubaddressAccount{
					AccountIndex: newAccountResponse.AccountIndex,
					Label:        newNamePending,
					BaseAddress:  newAccountResponse.Address,
				}
				pendingMap[newNamePending] = newAccount
			}

			wg.Add(1)
			go DriveTransaction(account,
				transfer,
				pendingMap[newNamePending],
				safeMap[safeNameMap[account.AccountIndex]],
				&wg,
				client)
		}
	}

	wg.Wait()

	return nil
}
