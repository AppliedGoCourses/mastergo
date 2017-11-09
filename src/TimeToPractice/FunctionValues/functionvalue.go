package main

import "fmt"

// Define a function that returns two func()s.
func newClosures() (func(), func() int) {
	// This is our outer variable.
	a := 0

	// Now we create and return two closures.

	return func() {
			a = 5
		},
		func() int {
			a = a * 7
			return a
		}
}

func main() {
	f1, f2 := newClosures()
	f1()      // sets "a" to 5
	n := f2() // multiplies "a" by 7 - is f2's internal value of "a" 0 or 5 before the call?
	fmt.Println(n)
}
