package main

import (
	"errors"
	"fmt"
)

func main() {

	a := 1
	b := 2
	if a == b { // No parentheses required
		fmt.Println("a and b are equal")
	} else { // "}" and "else" must be on the same line!
		fmt.Println("a and b are different")
	}

	// Initialization statement
	if err := f(0); err != nil {
		fmt.Println(err.Error())
	}
	// err is not visible outside the if construct
}

// Sample function that returns an error value
func f(n int) error {
	if n == 0 {
		return errors.New("n must not be 0")
	}
	return nil // no error occurred
}
