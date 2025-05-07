package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	stringSlice := make([]string, len(ip))
	for i, val := range ip {
		stringSlice[i] = strconv.Itoa(int(val))
	}
	return strings.Join(stringSlice, ".")

}

// we could use this instead -> ip addr is an array of size 4
// the length is fixed and therefore we can just hard code it
// func (ip IPAddr) String() string {
// 	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
// }

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
