package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Declare an io.Writer interface variable.
	// The zero type of an interface is nil.
	var w io.Writer
	fmt.Println("w is nil:", w == nil)

	// os.Stdout is a predefined variable of type *os.File.
	w = os.Stdout
	fmt.Println("w is nil:", w == nil)

	if w != nil {
		_, err := w.Write([]byte("w != nil\n"))
		if err != nil {
			log.Fatalln(err)
		}
	}

	// After assigning a nil variable to the interface,
	// the interface is not nil but the variable still is.
	var f *os.File
	w = f
	fmt.Println("w is nil:", w == nil)
	// Gotcha. w != nil but f == nil, hence w.Write() fails.
	if w != nil {
		_, err := w.Write([]byte("w != nil\n"))
		// This demonstrates the benefits of proper
		// error checking.
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func init() {
	// Show file name and line number, instead of date and time.
	log.SetFlags(log.Lshortfile)
}
