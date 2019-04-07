package dots

import (
	"encoding/json"
	"fmt"
	"strings"
)

func isSafe(name string) bool {
	return strings.HasSuffix(name, config.SafeSuffix)
}

func nameSafe(name string) string {
	return name + config.SafeSuffix
}

func isPending(name string) bool {
	return strings.HasPrefix(name, config.PendingPrefix)
}

func namePending(name, utxo string) string {
	return config.PendingPrefix + name + ":" + Abbreviate(utxo)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// Abbreviate compact XMR address, easier to read
func Abbreviate(address string) string {
	//return address[:6] + "..." + address[len(address)-6:]
	return address[:6]
}
