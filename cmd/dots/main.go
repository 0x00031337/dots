package main

import (
	"fmt"
	"github.com/fraudmarc/dots/pkg/dots"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	DisplayGreeting()
	app := kingpin.New("dots", "Dots, not lines, for safer cryptocurrency.")
	app.Version("0.0.1")

	config := configureDotsCommand(app)

	dots.ServiceWallet(config)
}

func DisplayGreeting() {
	fmt.Println(dots.Greeting)
}

func configureDotsCommand(app *kingpin.Application) *dots.DotsConfig {
	c := &dots.DotsConfig{}

	app.Flag("safe-suffix", "Suffix for safe accounts.").
		Default("-safe").
		StringVar(&c.SafeSuffix)
	app.Flag("pending-prefix", "Prefix for pending accounts.").
		Default("dots-").
		StringVar(&c.PendingPrefix)
	app.Flag("delay", "Delay first transaction by this duration.").
		Default("0s").
		DurationVar(&c.ScheduleDelayStart)
	app.Flag("finish", "Try to finish transfers by this duration from now.").
		Default("5h").
		DurationVar(&c.ScheduleFinish)
	app.Flag("min-moves", "Minimum number of moves for an output. MUST be 3+.").
		Default("3").
		IntVar(&c.ScheduleMinMoves)
	app.Flag("max-moves", "Maximum number of moves for an output.").
		Default("5").
		IntVar(&c.ScheduleMaxMoves)
	app.Flag("wallet-rpc", "Wallet RPC location.").
		Default("http://127.0.0.1:28081/json_rpc").
		StringVar(&c.RpcURL)
	app.Flag("mainnet", "Use main network (not testnet). (default:false)").
		BoolVar(&c.Mainnet)
	app.Flag("do-relay", "Relay transactions to network. (default:false)").
		BoolVar(&c.RelayTransactions)
	app.Flag("no-donations", "Do not send donations. Privacy implications unknown. (default:false)").
		BoolVar(&c.NoDonations)

	app.Parse(os.Args[1:])

	return c
}
