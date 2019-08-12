package main

import (
	"fmt"
)

// ...string denotes one or more parameters of type string.
func longest(str ...string) int {
	// str is a "slice" of strings. We will discuss slices
	// in section 3 of the course.
	var length int
	// A range loop visits each element of a variadic parameter.
	for _, s := range str {
		if len(s) > length {
			length = len(s)
		}
	}
	return length
}

func main() {
	fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
	fmt.Println(longest("Your", "word", "list", "here"))
}
