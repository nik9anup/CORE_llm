/*
Package main lists all network interfaces on the system and their associated IP addresses.

This program uses the `net` package to retrieve and display all network interfaces available
on the system along with their IP addresses. It handles any errors encountered during the
retrieval of interfaces or addresses by using `panic` to terminate the program with an error message.
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// Get all network interfaces on the system.
	// If there is an error in retrieving the interfaces, the program will terminate with a panic.
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	// Iterate through each network interface.
	for _, interf := range interfaces {
		// Resolve and display addresses for each interface.
		// If there is an error in retrieving the addresses, the program will terminate with a panic.
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		// Print the name of the network interface.
		fmt.Println(interf.Name)

		// Iterate through each address associated with the interface.
		for _, add := range addrs {
			// Check if the address is of type *net.IPNet and print it.
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}
}