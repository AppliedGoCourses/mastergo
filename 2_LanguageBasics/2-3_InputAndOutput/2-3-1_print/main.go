package main

import "fmt"

func main() {

	fmt.Println("Hello, world")

	// Multiple arguments
	fmt.Println("Hello,", "world")

	// Arguments of different type (as long as the type
	// has a string representation)
	fmt.Println("h3110,", 1337, "60ph3r")

	// Print decimal and hex integers (the latter will become decimal)
	fmt.Println(1234, 0x1234)

	// String concatenation
	fmt.Println("cat" + "bird")

	a := 10
	b := true
	c := "ten"

	// Formatted printing.
	fmt.Printf("It is %t that %d is spelled '%s'.", b, a, c)

}
