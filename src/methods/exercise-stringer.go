package main

import (
	"fmt"
	"strconv"
)
type IPAddr [4]byte

// TODO IPAddr型にStringメソッドを実装する

func (ia IPAddr) String() string {
	res := ""

	for _, v := range ia {
		if len(res) > 0 {
			res += "."
		}

		// intからstring変換した上で連結する
		res += strconv.Itoa(int(v))
	}
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)

	}
}