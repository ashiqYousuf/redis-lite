package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		log.Fatal("Invalid type, expecting bulk strings only")
	}

	size, _ := reader.ReadByte()
	strSize, err := strconv.ParseInt(string(size), 10, 64)
	if err != nil {
		log.Fatal("invalid type conversion:", err)
	}

	// consume /r/r
	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))
}

func M() {
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
		buf := make([]byte, 1024)

		// Read message from client
		_, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal("error reading from client:", err)
		}

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
