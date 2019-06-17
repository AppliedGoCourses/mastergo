package main

import "fmt"

// Go's pass-by-value semantic: Passing a variable to a function
// means passing a copy of the variables value, not the variable
// itself.
// Changing this value inside the function has no effect on the
// original value. (See main() for the calling part.)
func f(a int) {
	a = a + 1
}

// A pointer is also passed by value.
// Changing the pointer itself (to point to some other variable)
// has no effect outside the function. `a` only lives within `g()`.
// This may be obvious here but can become a pitfall when a pointer
// is hidden within a data structure passed to a function.
// (Again, the calling part is in main())

func g(a *int) {
	b := 0
	a = &b
}

// Through a pointer parameter, a function can manipulate
// values outside its scope through pointer indirection.

func h(a *int) {
	*a = *a + 1
}

// Create a local variable and return a pointer to it.
// In this case, the local variable outlives the function.
// The Go compiler figures out when to retain a local
// variable through "escape analysis".
//
// Interested in the details? Run this file as
// go run -gcflags '-m -l' functionsandpointers.go
// and inspect the output.

func i() *int {
	// a has function scope but continues to live after h() exits,
	// because the function returns a's address to the caller.
	// In C or C++, this would be an error. a would cease to exist,
	// and the returned pointer would point to an unused memory location.
	a := 7
	fmt.Println("a is stored at", &a)
	return &a
}

func main() {
	fmt.Println("*** PASS BY VALUE***")
	x := 1
	f(x)           // Pass x to v()
	fmt.Println(x) // Did f() change x?

	fmt.Println("\n*** POINTER INDIRECTION ***")
	x = 1
	g(&x)          // Pass a pointer to x to f()
	fmt.Println(x) // Did g() change x?

	fmt.Println("\n*** LOCAL POINTER CHANGE ***")
	x = 1
	h(&x)          // Same test but with g()
	fmt.Println(x) // Did h() change x?

	fmt.Println("\n*** RETURN POINTER TO LOCAL VARIABLE ***")
	y := i()
	// y points to a's storage location
	fmt.Println("y is stored at", y, "and has the value", *y)
	// We can see that a still lives after the call to i()
	// and can be accessed through *y.
}
