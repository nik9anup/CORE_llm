// This code sets up a UDP server on port 7070 that listens for incoming packets,
// reads messages from clients, and echoes back a confirmation message to the sender.
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	pc, err := net.ListenPacket("udp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	buffer := make([]byte, 2048)
	fmt.Println("Waiting for client...")
	for {

		_, addr, err := pc.ReadFrom(buffer)
		if err == nil {
			rcvMsq := string(buffer)
			fmt.Println("Received: " + rcvMsq)
			if _, err := pc.WriteTo([]byte("Received: "+rcvMsq), addr); err != nil {
				fmt.Println("error on write: " + err.Error())
			}
		} else {
			fmt.Println("error: " + err.Error())
		}

	}

}