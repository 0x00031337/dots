package dots

import (
	"github.com/fraudmarc/go-monero/walletrpc"
	"time"
)

const (
	//	safeSuffix    = "-safe"
	//	pendingPrefix = "dots-"
	txPriority = walletrpc.PriorityUnimportant

//	txPriority         = walletrpc.PriorityNormal
//	flatDevFee         = 100000000
//	Donate0Name        = "test1-accountTest-0"
//	Donate0Address     = "Bc8u5ozsDytAuVXWesZ8hD7pjBvifzHgqZeWn7M2rWNvdMBea48ABcz678fujDxgz9YQd9a5zNuzgYjUGQ7BBsiXUYHocJz"
//	Donate1Name        = "test1-accountDonate-0"
//	Donate1Address     = "BdqMnmb7boP3C3dk1R9aPNjePbCvCAHm3DAgMy2yGHQgVZa9G3AHXGUCZZ5znAu6DaEqmDTGj5uBLK924ymbBG4yTeiAPFh"
//	Donate2Name        = "test1-accountDonate1-0"
//	Donate2Address     = "Bam7W3Jsnc8hyBgXmYhFDEFzVjkoXJUb4PteRyVUHqQ2P6gCPppvW97SaLdDgiPUSsZLeVV9CqPsqeJXLq8AGNJPPz83XsU"
//	ScheduleMinSpacing = 1
//	ScheduleDelayStart = 0     // delay first transaction until this time
//	ScheduleFinish     = 18000 // 5 hrs must be greater than delayStart. This is time from now until it should finish
//	//ScheduleFinish        = 8 // must be greater than delayStart. This is time from now until it should finish
//	ScheduleMinMoves      = 3
//	ScheduleMaxMoves      = 5
//	TransactionDoNotRelay = false
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
}

var config *DotsConfig
