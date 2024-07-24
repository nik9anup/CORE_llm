// This code establishes a TCP server on port 8080 that accepts incoming connections,
// assigns a unique client ID to each connection, sends a welcome message to clients,
// and echoes back received messages with a prefix indicating reception.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	ID := 0
	for {
		fmt.Println("Waiting for client...")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Client ID: %d connected.\n", ID)
		go func(c net.Conn, clientID int) {
			fmt.Fprintf(c, "Welcome client ID: %d \n", clientID)
			for {
				msg, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					break
				}
				_, err = io.WriteString(c, "Received: "+string(msg))
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			fmt.Println("Closing connection")
			c.Close()
		}(conn, ID)
		ID++
	}

}