package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Argument required: please provide the path to a text file.")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// the Set variable.

	set := make(map[string]struct{})

	// For easy iteration, we use a scanner.
	// NewScanner takes an io.Reader, and os.File is
	// an os.Reader, so we can pass f to NewScanner.
	// (Learn more about this in the lectures on interfaces.)

	s := bufio.NewScanner(f)
	// Scan() advances to the next token, which is a line of text by default. It returns true as long as the scan can advance.
	for s.Scan() {
		// The Text() method returns the current token as a string.
		if _, ok := set[s.Text()]; ok {
			fmt.Println("Duplicate line:", s.Text())
		} else {
			set[s.Text()] = struct{}{}
		}
	}
}
