package dots

import (
	"fmt"
	"github.com/fraudmarc/go-monero/walletrpc"
	"os"
)

func ServiceWallet(newConfig *DotsConfig) {
	config = newConfig

	PrettyPrint(config)

	// Start a wallet client instance
	client := walletrpc.New(walletrpc.Config{
		Address: config.RpcURL,
	})

	accounts, err := client.GetAccounts()
	if err != nil {
		if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
			// it is a monero wallet error
			fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
			os.Exit(1)
		}
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	ProcessAccounts(accounts.SubaddressAccounts, client)
}
