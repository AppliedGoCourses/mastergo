package main

import (
	"fmt"
	"strings"
)

// Standard variable declaration
var a int

// Declaration and initialization
var b int = 10

// Type deduction: c becomes an int
var c = 10

// Multiple declarations of the same type
var d, e, f bool

// Multiple declarations, differnet types, with and without initialization
// Tip: run "go fmt <yourfile.go>" to get the alignment of types as seen here.
// Better yet, configure your editor to run "go fmt" automatically on saving
// the file. (Many editors have Go plugins that do this by default.)
var (
	g       int
	h       string
	i       int = 1234
	j, k, l bool
)

// All of the above is also possible in main()
func main() {
	var a int
	var b int = 10
	var c = 10
	var d, e, f bool

	var (
		g       int
		h       string
		i       int = 1234
		j, k, l bool
	)

	// The shortcut operator can only be used inside a function
	m := 1

	// Multiple declaration+assignment
	n, o := 2, 3

	// a is already declared, now we can assign to it
	a = 11

	// This swaps the contents of e and f
	e, f = f, e

	// Variables cannot switch their type. The commented line
	// would error out if uncommented.
	p, q := 100, 200
	// p = "wrong"

	// f() returns two results, but we need only one.
	// Declare the second one as unused by using the
	// blank identifier instead of a variable name.
	r, _ := fn1()
	fn2(r)

	// Strings are immutable. String modifications produce
	// a copy of the original string.
	s := "I Am immutable."
	s = strings.Replace(s, "A", "a", 1)

	// To avoid the (otherwise very helpful!) "declared and not used" error
	// in this demo code.
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
}

func fn1() (int, string) {
	return 1, "one"
}

func fn2(n int) {
	// some code...
}
