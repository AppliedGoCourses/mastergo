package main

import (
	"fmt"
)

func main() {
	welcome := "მისასალმებელია"
	fmt.Println(welcome)

	// A range loop over a string recognizes Unicode runes.
	for i, r := range welcome {
		fmt.Printf("index of %#U: %d\n", r, i)
	}

	fmt.Println()

	welcomeB := []byte("მისასალმებელია")

	// Println prints the slice like any other slice -
	// as a list of single values, in brackets.
	fmt.Println(welcomeB)

	// To print a byte slices as a string, use Printf...
	fmt.Printf("%s\n", welcomeB)

	// ... or a type conversion.
	fmt.Println(string(welcomeB))

	// A loop over a byte slice is not Unicode-aware;
	// it processes the slice byte by byte.
	for i, r := range welcomeB {
		fmt.Printf("Index of %#U: %d\n", r, i)
	}
}
