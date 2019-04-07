package dots

import (
	"crypto/rand"
	"log"
	"math/big"
	"sort"
	"time"
)

func GetRandomSchedule(delay, total time.Duration, number int) []time.Duration {
	diff := total - delay

	var targetTimes []time.Duration
	var delays []time.Duration

	// space out over the whole time window
	for i := 0; i < number; i++ {
		myChoiceBig, err := rand.Int(rand.Reader, big.NewInt(diff.Nanoseconds()))
		if err != nil {
			log.Println(err)
			myChoiceBig = big.NewInt(0)
		}
		myChoice := time.Duration(myChoiceBig.Int64())

		targetTimes = append(targetTimes, myChoice)
	}

	// chronological sort
	sort.Slice(targetTimes, func(i, j int) bool { return targetTimes[i] < targetTimes[j] })

	// make array of the individual sleep times between tx
	for i, target := range targetTimes {
		if i == 0 {
			delays = append(delays, delay+target)
			continue
		}
		delays = append(delays, target-targetTimes[i-1])
	}

	return delays
}

func GetRandomMoveCount(min, max int) int {
	diff := max - min + 1
	myChoiceBig, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	if err != nil {
		log.Println(err)
		myChoiceBig = big.NewInt(int64(min))
	}
	myChoice := int(myChoiceBig.Int64())

	return myChoice + min
}
