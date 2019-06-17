/*
# Test Your Skills: Closures

## Intro

As you learned in the lecture about function values and closures, a closure can reference the outer function's variables even after the outer function has terminated.

But what happens if the outer function generates and returns *two* closures?

Do they access the same outer variables, or does each of them get its own copy?

## Your task

Write a function that returns two closures. Both closures shall access a variable defined in the outer function.

Then write code that reveals whether both closures use the same instance of the outer variable, or whether each of them has its own individual instance.


## Solution

Here is one possible solution to the task. There are of course many ways to create this code, so you might have came up with a different approach. That's absolutely ok, as long as your code comes to the same conclusion as the following code.

*/

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
