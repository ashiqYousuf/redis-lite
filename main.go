package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ashiqYousuf/redis-lite/resp"
)

func main() {
	// Create a new server
	// ":PORT" means the server will listen on all available
	// network interfaces on port PORT.
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	fmt.Println("Listening on port: 6379")

	// Listen for connections
	// l.Accept() waits for an incoming client connection
	// and returns a connection object (conn).
	conn, err := l.Accept()
	if err != nil {
		log.Fatal("cannot accept connections:", err)
	}

	// Ensure the connection is properly closed when the function exits.
	defer conn.Close()

	// Now you can use conn to read from/write to the client.
	// So create an infinite loop and receive commands from
	// clients and respond to them.

	for {
		resp := resp.NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(value)
		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
