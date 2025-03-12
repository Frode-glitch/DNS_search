package main

import (
	"fmt"
	"net"
	"os"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Give me an IP address or domain name: ")
	fmt.Scan(&input)

	if isIP(input) {
		names, err := net.LookupAddr(input)
		if err != nil {
			fmt.Printf("Could not get domain names: %v\n", err)
			os.Exit(1)
		}
		for _, name := range names {
			fmt.Printf("IP address %s resolves to domain name: %s\n", input, name)
		}
	} else {
		ips, err := net.LookupIP(input)
		if err != nil {
			fmt.Printf("Could not get IPs: %v\n", err)
			os.Exit(1)
		}
		for _, ip := range ips {
			fmt.Printf("Domain %s has IP address: %s\n", input, ip.String())
		}
	}
}

func isIP(input string) bool {
	for _, r := range input {
		if !unicode.IsDigit(r) && r != '.' {
			return false
		}
	}
	return true
}
