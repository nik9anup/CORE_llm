/*
Package main demonstrates how to set up a simple RPC server and client using net/rpc and net/rpc/jsonrpc packages in Go.

This program defines a simple RPC server that exposes an Add method, which adds two integers received in Args struct
and returns the result in Result struct. The client connects to this server using JSON-RPC over TCP, sends an Add RPC
request, and prints the result received from the server.
*/
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Args represents the arguments for the Add method.
type Args struct {
	A, B int
}

// Result represents the result of the Add method.
type Result int

// RpcServer represents the RPC server type.
type RpcServer struct{}

// Add is an RPC method on RpcServer that adds two integers.
func (t RpcServer) Add(args *Args, result *Result) error {
	log.Printf("Adding %d to %d\n", args.A, args.B)
	*result = Result(args.A + args.B)
	return nil
}

// addr represents the address and port on which the RPC server listens.
const addr = ":7070"

func main() {
	// Start the RPC server in a separate goroutine.
	go createServer(addr)

	// Connect to the RPC server using JSON-RPC over TCP.
	client, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Prepare arguments for the RPC call.
	args := &Args{
		A: 2,
		B: 3,
	}
	var result Result

	// Call the Add method on the RPC server.
	err = client.Call("RpcServer.Add", args, &result)
	if err != nil {
		log.Fatalf("Error calling RpcServer.Add: %s", err)
	}

	// Print the result received from the RPC server.
	log.Printf("%d + %d = %d\n", args.A, args.B, result)
}

// createServer creates and starts an RPC server that listens for incoming connections on the specified address.
func createServer(addr string) {
	// Create a new RPC server instance.
	server := rpc.NewServer()

	// Register the RpcServer type to handle RPC requests.
	err := server.Register(RpcServer{})
	if err != nil {
		panic(err)
	}

	// Listen for incoming TCP connections on the specified address.
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalf("Couldn't start listening on %s: %s", addr, e)
	}

	// Accept incoming connections and serve RPC requests using JSON-RPC codec.
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}