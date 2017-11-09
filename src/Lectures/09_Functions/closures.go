package main

import "fmt"

// *** Function as value

// A function variable

func f1(s string) bool { return len(s) > 0 }
func f2(s string) bool { return len(s) < 4 }

var funcVar func(string) bool = f1

// A function as a function parameter
func funcAsParam(s string, f func(string) bool) bool {
	return f(s + "abcd")
}

// Define and call a function literal in one statement by appending a parameter list () to the function literal.
var result string = func() string {
	return "abcd"
}() // <- This is the function call (an empty parameter list)

// Closures
func newClosure() func() {
	// Variable a will be used by the closure. Because of this,
	// it continues to live even after newClosure() ends.
	var a int

	// Create and return a closure that refers to a.
	// When the closure runs, it can read and update a,
	// even though newClosure() will have exited by then.
	return func() {
		fmt.Println(a)
		a++
	}
}

// A caveat: Closures and loop variables
func caveat() {
	s := "abcd"

	// We will store a couple of funcs in this slice.
	var funcs []func()

	// CAVEAT: c is created ONCE, and all closures created in the loop
	// will refer to this one c that constantly changes its value.
	// When the loop finishes, c contains the value "d", and so all closures
	// in the funcs slice will read "d" from variable c when executing.
	for _, c := range s {
		// REMEDY: Uncomment the following assignment to declare a new c in loop body scope that is initialized from the range variable of the same name. Each loop iteration produces a new instance of this c.
		// This way, each new closure refers to a new instance of the loop body's c rather than to the single instance of the range variable c.
		// c := c
		funcs = append(funcs, func() {
			// fmt.Print(string(c))
			fmt.Println(&c)
		})
	}
	for _, f := range funcs {
		f()
	}
}

func main() {
	fmt.Println("\n*** FUNCTION VALUES ***")

	// Function literal
	funcVar = func(s string) bool {
		return len(s) > 4
	}

	fmt.Println(funcVar("abcd"))
	fmt.Println(funcAsParam("abcd", funcVar))

	fmt.Println(result)

	fmt.Println("\n*** CLOSURES ***")
	c := newClosure()
	// newClosure() has exited, but a still exists and is updated
	// on each call to the closure
	c()
	c()
	c()

	fmt.Println("\n*** CAVEAT ***")
	caveat()
}
