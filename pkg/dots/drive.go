package dots

import (
	"fmt"
	"github.com/fraudmarc/go-monero/walletrpc"
	"os"
	"sync"
	"time"
)

// move a single utxo from start account through pendings to finish
func DriveTransaction(accountFrom walletrpc.SubaddressAccount,
	transfer walletrpc.IncTransfer,
	accountPending walletrpc.SubaddressAccount,
	accountSafe walletrpc.SubaddressAccount,
	wg *sync.WaitGroup,
	client walletrpc.Client) {

	defer wg.Done()

	// setup our schedule within constraints
	numMoves := GetRandomMoveCount(config.ScheduleMinMoves, config.ScheduleMaxMoves)
	schedule := GetRandomSchedule(config.ScheduleDelayStart, config.ScheduleFinish, numMoves)

	fmt.Printf("Driving: %s with %7d moves\n", Abbreviate(transfer.TxHash), numMoves)
	for i, value := range schedule {
		fmt.Printf("  Sleeping for %v... ", value.String())
		time.Sleep(value)

		if i == 0 {
			// first pass, move from original to -pending
			fmt.Println("sweep to pending")
			res, err := SweepSingle(accountFrom, transfer, accountPending, client)
			if err != nil {
				if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
					// insufficient funds return a monero wallet error
					// walletrpc.ErrGenericTransferError
					fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
					os.Exit(1)
				}
				fmt.Println("Error:", err.Error())
				os.Exit(1)
			}
			PrettyPrint(res)
			continue
		}

		if i == int(numMoves)-1 {
			// last move. to -safe
			fmt.Println("sweep to safe")
			res, err := SweepAll(accountPending, accountSafe, client)
			if err != nil {
				if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
					// insufficient funds return a monero wallet error
					// walletrpc.ErrGenericTransferError
					fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
					os.Exit(1)
				}
				fmt.Println("Error:", err.Error())
				os.Exit(1)
			}
			PrettyPrint(res)
			break
		}

		fmt.Println("donate")
		donation := ChooseDonation()
		fmt.Printf("Donating to %s\n", donation.Name)
		res, err := Donate(accountPending, donation.Address, client)
		if err != nil {
			if iswerr, werr := walletrpc.GetWalletError(err); iswerr {
				// insufficient funds return a monero wallet error
				// walletrpc.ErrGenericTransferError
				fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
				os.Exit(1)
			}
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
		PrettyPrint(res)
	}

	fmt.Println(" Finished: ", Abbreviate(transfer.TxHash))
}
