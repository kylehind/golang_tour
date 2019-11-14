package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	fmt.Printf("%T", ip[0])
	fmt.Printf("%T", ip[1])
	fmt.Printf("%T", ip[2])
	fmt.Printf("%T", ip[3])

	fmt.Print(ip[0])
	fmt.Print(ip[1])
	fmt.Print(ip[2])
	fmt.Print(ip[3])

	return fmt.Sprintf("%o.%o.%o.%o", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("\n%v: %v\n", name, ip)
	}
}
