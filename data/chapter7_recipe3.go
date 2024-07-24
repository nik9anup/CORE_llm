/*
Package main demonstrates DNS lookup functionalities for resolving IP addresses and hostnames.

This program resolves the hostnames associated with the loopback IP address (`127.0.0.1`) and the IP addresses for the hostname `localhost`.
It uses functions from the `net` package to perform these lookups and prints the results to the console.
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve the hostname associated with the loopback IP address (127.0.0.1).
	// If there is an error during the lookup, the program will terminate with a panic.
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	// Print each resolved hostname.
	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// Resolve the IP addresses associated with the hostname "localhost".
	// If there is an error during the lookup, the program will terminate with a panic.
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}

	// Print each resolved IP address.
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}