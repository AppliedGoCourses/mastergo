package main

// To get a shorter application name in the output, try:
// go build cmdline.go
// ./cmdline one two three

import (
	"fmt"
	"os"
)

func main() {

	// os.Args is a slice of all command line arguments
	// (including the name of the command)
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])

	// As os.Args is a slice, we can iterate over it.
	for _, params := range os.Args {
		fmt.Println(params)
	}
}
