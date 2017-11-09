package main

import (
	"fmt"
	"os"
)

// Recursion

func rec(n int) int {
	if n == 0 {
		fmt.Println("No more recursive calls: returning", n)
		return 0
	}
	fmt.Println("n is", n, "- calling rec(", n-1, ")")
	ret := rec(n - 1)
	fmt.Println("Call returned", ret)
	return ret + n
}

// Deferred function calls

func d() error {
	f, err := os.Open("functionbehavior.go")
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println(f.Name())
	if f.Name() == "functionbehavior.go" {
		return nil
	}
	// More code, maybe more exit points...
	return nil
}

// Scope

// Function scope
func f() {
	a := 2
	fmt.Println("func f: a is", a)
}

// Block scope
func g() {
	a := 3
	b := 1
	{
		a := 4 // a shadows the declaration of a outside the block!
		b = 2  // here, we only re-assign a new value to outer b
		fmt.Println("in block: a is", a, "and b is", b)
	}
	fmt.Println("g: a is", a, "and b is", b)
}

// The scope of loop variables
func h() {
	// Loop scope rule: i is instantiated only once,
	// whereas n gets a new instance on each iteration.
	for i := 0; i < 4; i++ {
		n := i
		fmt.Println(&i, &n) // compare the addresses of i and in in the output
	}
}

func main() {
	fmt.Println("\n*** RECURSION ***")
	fmt.Println("Recursion result: ", rec(5))

	fmt.Println("\n*** DEFER ***")
	d()

	fmt.Println("\n*** FUNCTION SCOPE ***")
	a := 1
	fmt.Println("main: a is", a)
	f()
	fmt.Println("main: a is", a)

	fmt.Println("\n*** BLOCK SCOPE ***")
	g()

	fmt.Println("\n*** LOOP VARIABLE SCOPE ***")
	h()
}
