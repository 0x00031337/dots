package dots

import (
	"github.com/fraudmarc/go-monero/walletrpc"
	"time"
)

const (
	txPriority = walletrpc.PriorityNormal
)

const (
	Greeting = "     _____          ___                       ___     \n" +
		"    /  /::\\        /  /\\          ___        /  /\\\n" +
		"   /  /:/\\:\\      /  /::\\        /  /\\      /  /:/_\n" +
		"  /  /:/  \\:\\    /  /:/\\:\\      /  /:/     /  /:/ /\\  \n" +
		" /__/:/ \\__\\:|  /  /:/  \\:\\    /  /:/     /  /:/ /::\\\n" +
		" \\  \\:\\ /  /:/ /__/:/ \\__\\:\\  /  /::\\    /__/:/ /:/\\:\\\n" +
		"  \\  \\:\\  /:/  \\  \\:\\ /  /:/ /__/:/\\:\\   \\  \\:\\/:/~/:/\n" +
		"   \\  \\:\\/:/ __ \\  \\:\\  /:/  \\__\\/  \\:\\   \\  \\::/ /:/\n" +
		"    \\  \\::/ /_/\\ \\  \\:\\/:/   __   \\  \\:\\   \\__\\/ /:/\n" +
		"     \\__\\/  \\_\\/  \\  \\::/   /_/\\   \\__\\/ __  /__/:/\n" +
		"                   \\__\\/    \\_\\/        /_/\\ \\__\\/   __\n" +
		"                                        \\_\\/        /_/\\\n" +
		"                                                    \\_\\/\n"
)

type DotsConfig struct {
	SafeSuffix         string
	PendingPrefix      string
	TxPriority         walletrpc.Priority
	ScheduleDelayStart time.Duration
	ScheduleFinish     time.Duration
	ScheduleMinMoves   int
	ScheduleMaxMoves   int
	RelayTransactions  bool
	RpcURL             string
	Mainnet            bool
	NoDonations        bool
}

var config *DotsConfig
