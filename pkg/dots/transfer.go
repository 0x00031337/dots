package dots

import (
	"fmt"
	"github.com/fraudmarc/go-monero/walletrpc"
	"time"
)

func Donate(accountFrom walletrpc.SubaddressAccount,
	addressTo string,
	client walletrpc.Client) (*walletrpc.TransferResponse, error) {

	var res *walletrpc.TransferResponse
	var err error
	for true {
		// don't relay this transaction, it's only finding the current fee amount
		res, err = client.Transfer(walletrpc.TransferRequest{
			Destinations: []walletrpc.Destination{
				{
					Address: addressTo,
					Amount:  395580000,
				},
			},
			AccountIndex: accountFrom.AccountIndex,
			Priority:     txPriority,
			DoNotRelay:   !config.RelayTransactions,
			//GetTxHex:      true,
			//GetTxMetadata: true,
		})
		if err != nil {
			// continue until we can at least get tx fee
			fmt.Printf(".")
			time.Sleep(1 * time.Minute)
			continue
		}

		donationAmount := GetDonationAmount(res.Fee, txPriority)

		// now we know fee and donation amount, relay this one
		res, err = client.Transfer(walletrpc.TransferRequest{
			Destinations: []walletrpc.Destination{
				{
					Address: addressTo,
					Amount:  donationAmount,
				},
			},
			AccountIndex: accountFrom.AccountIndex,
			Priority:     txPriority,
			DoNotRelay:   !config.RelayTransactions,
			//GetTxHex:      true,
			//GetTxMetadata: true,
		})
		if err == nil {
			break
		}

		fmt.Printf(".")
		time.Sleep(1 * time.Minute)
	}

	return res, err
}

func SweepSingle(accountFrom walletrpc.SubaddressAccount,
	transfer walletrpc.IncTransfer,
	accountTo walletrpc.SubaddressAccount,
	client walletrpc.Client) (*walletrpc.SweepSingleResponse, error) {

	var res *walletrpc.SweepSingleResponse
	var err error
	for true {
		res, err = client.SweepSingle(walletrpc.SweepSingleRequest{
			Address:      accountTo.BaseAddress,
			AccountIndex: accountFrom.AccountIndex,
			Priority:     txPriority,
			KeyImage:     transfer.KeyImage,
			DoNotRelay:   !config.RelayTransactions,
			//GetTxHex:      true,
			//GetTxMetadata: true,
		})
		if err == nil {
			break
		}
		fmt.Printf(".")
		time.Sleep(1 * time.Minute)
	}

	return res, err
}

func SweepAll(accountFrom walletrpc.SubaddressAccount,
	accountTo walletrpc.SubaddressAccount,
	client walletrpc.Client) (*walletrpc.SweepAllResponse, error) {

	var res *walletrpc.SweepAllResponse
	var err error
	for true {
		res, err = client.SweepAll(walletrpc.SweepAllRequest{
			Address:      accountTo.BaseAddress,
			AccountIndex: accountFrom.AccountIndex,
			Priority:     txPriority,
			DoNotRelay:   !config.RelayTransactions,
			//GetTxHex:      true,
			//GetTxMetadata: true,
		})
		if err == nil {
			break
		}
		fmt.Printf(".")
		time.Sleep(1 * time.Minute)
	}

	return res, err
}

func ShowTransfer(transfer walletrpc.IncTransfer) error {
	fmt.Println("  Showing Transfer")
	fmt.Printf("    Amount: %s\n", walletrpc.XMRToDecimal(transfer.Amount))
	fmt.Printf("    Global Index: %d\n", transfer.GlobalIndex)
	fmt.Printf("    Key Image: %s\n", transfer.KeyImage)
	fmt.Printf("    Spent: %v\n", transfer.Spent)
	fmt.Printf("    Subaddress Index: (Major: %d, Minor: %d)\n",
		transfer.SubAddressIndex.Major,
		transfer.SubAddressIndex.Minor)
	fmt.Printf("    TxHash: %s\n", transfer.TxHash)
	fmt.Printf("    Size: %d\n", transfer.TxSize)

	return nil
}
