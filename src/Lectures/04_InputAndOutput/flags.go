package main

// See inputoutput.go for a commented summary of all code.

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// verbose and count become pointers to bool and int, respectively.
	verbose := flag.Bool("verbose", false, "Print details")
	count := flag.Int("count", 1, "Number of iterations")

	// max is not a pointer. We pass a pointer to max to flag.IntVar
	// so that flag.IntVar can update max
	var max int
	flag.IntVar(&max, "max", 10, "Maximal sum")

	// Read all the flags
	flag.Parse()

	if *verbose == true {
		// Print the type of count and its content
		fmt.Printf("%#v\n", *count)
	} else {
		fmt.Printf("%d\n", *count)
	}

	if len(os.Args) > max+1 {
		fmt.Println("Found", len(os.Args), "parameters. Only", max, "are allowed.")
	}

	// Get the arguments WITHOUT the flags through flag.Args().
	// (flag.Parse() must have been called before.)

	fmt.Println("All arguments:", os.Args)
	fmt.Println("Non-flag arguments:", flag.Args())
}
