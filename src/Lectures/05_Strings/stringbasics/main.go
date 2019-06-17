package main

import "fmt"

func main() {

	// This is a format string with a placeholder
	var s string = "This string contains %d bytes.\n"

	// l is the value to substitute for the placeholder
	var l int = len(s)

	// Format and print.
	fmt.Printf(s, l)

	// Get one byte from the string.
	var b byte = s[2]
	fmt.Println("s[2] =", b)
	fmt.Println("s[2] =", string(b))

	// exctract a slice of bytes from the string.
	fmt.Println(s[7:11])

}
