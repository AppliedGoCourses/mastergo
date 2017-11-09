package main

import (
	"fmt"
	"unicode"
)

// The original main function, commented.
func originalMain() {

	// This s lives outside the loop.
	s := "abcde"

	// The left s lives in the scope of the loop condition.
	// It overshadows the right s, which is the s from above.
	for _, s := range s {

		// The left s is a new variable in the loop body,
		// and the right s is the s from the loop condition.
		s := unicode.ToUpper(s)

		// This is still the s from the loop body.
		fmt.Print(string(s))
	}

	// This is again the s from the function's topmost scope level.
	fmt.Println("\n" + s)
}

// main rewritten with sane variable names, no shadowing.
func main() {
	str := "abcde"
	for _, letter := range str {
		upper := unicode.ToUpper(letter)
		fmt.Print(string(upper))
	}
	fmt.Println("\n" + str)
}
