package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// tcpserver listens on a local port and sends a message when a client connects.
func tcpserver() {
	listener, err := net.Listen("tcp", "localhost:9099")
	if err != nil {
		log.Fatalln("Cannot start listener: ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("Cannot get a connection: ", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalln("Cannot close connection: ", err)
		}
	}()

	// Write some data
	_, err = conn.Write([]byte("Hello from the other side!"))
	if err != nil {
		log.Fatalln("Cannot write to connection: ", err)
	}
}

func readFromServer(in io.Reader) {
	var buf bytes.Buffer
	n, err := buf.ReadFrom(in)
	if err != nil {
		log.Fatalln("Cannot read from in: ", err)
	}
	fmt.Println("Read", n, "bytes from connection:", buf.String())

	conn, ok := in.(net.Conn)
	if ok {
		fmt.Println("Read from remote address", conn.RemoteAddr())
	}

	switch stream := in.(type) {
	case net.Conn:
		fmt.Println("Read from remote address", stream.RemoteAddr())
	case *os.File:
		fmt.Println("Read from file", stream.Name())
	case *strings.Reader:
		fmt.Println("Read from a string of length", stream.Size())
	}

}

func main() {

	// Read from the network

	// Start the server
	go tcpserver()

	// Open a connection to the server
	conn, err := net.Dial("tcp", "localhost:9099")
	if err != nil {
		// We use a simple log.Fatalln in this sample code.
		// log.Fatal/f/ln terminates the app and thus is not the best
		// choice for production code.
		log.Fatalln("Cannot open connection: ", err)
	}

	defer conn.Close() // Error ignored as we only read from the connection

	readFromServer(conn)

	// Read from a file
	file, err := os.Open("loremipsum.txt")
	if err != nil {
		log.Fatal("Cannot open loremipsum.txt: ", err)
	}

	defer file.Close() // Error ignored as we only read from the file

	readFromServer(file)

	// Read from a string
	s := "The quick brown fox jumps over the lazy dog."

	str := strings.NewReader(s)

	readFromServer(str)

}

// functions named "init" run before any other
// functions in the same package.
func init() {
	// instead of the default date&time prefix, prepend
	// filename and line number to the log output.
	log.SetFlags(log.Lshortfile)
}
