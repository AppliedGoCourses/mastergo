package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Part 1: Are these structs comparable?

type s1 struct {
	n int
	b bool
}

type s2 struct {
	r []rune
}

type s3 struct {
	r [3]rune
}

// Part 2: Fun with empty structs

type emptyStruct struct{}

type metaEmptyStruct struct {
	A struct{}
	B struct{}
}

type sliceOfNothings []struct{}

// *** Task 1: A Set type ***

// A map with struct{} values acts like a set. The keys are the
// elements of the set, and inserting a key/value pair whose key
// already exists changes nothing. (The value is always empty anyway.)

// TODO: define a "Set of ints" type here.

func testSetType() {

	// TODO: Create a new int set variable.

	// Insert a value into the set

	// TODO

	// Test if a value exists in the set

	// TODO

	// Get all values from the set

	// TODO
}

// Task 2: An iterator over ints

// func iter shall return a data structure that a range loop
// can loop over. (See iterateFrom1ToN() below.)

// TODO: func iter(n int)...

// Now we can run a range loop over the return value of iter().
func iterateFrom1ToN(n int) {
	for i := range iter(n) {
		fmt.Print(i, ",")
	}
	fmt.Println()
}

func main() {
	fmt.Println("\n*** Part 1: Are these structs comparable?  ***\n")

	// TODO: If you are not sure about your answers, create
	//  two instances of each struct and try to compare them
	// to each other. Look for compile errors.

	fmt.Println("\n*** Part 2: Empty struct basics ***\n")

	es := emptyStruct{}
	mes := metaEmptyStruct{}

	fmt.Println("Size of es:", unsafe.Sizeof(es))
	fmt.Println("Size of mes:", unsafe.Sizeof(mes))

	sOfN := make(sliceOfNothings, math.MaxInt64) // try this with []int
	// We cannot use unsafe.Sizeof() on sOfN, since Sizeof() does not
	// follow pointers and therefore only shows the size of the slice
	// header, which is always the same.

	// Don't try to print out sOfN, it will try printing "{}" MaxInt64
	// times! Let's print out a shorter version instead.
	// Note that len() and cap() report values > 0, but this does not
	// mean that any memory has been allocated for the struct{}{} elements.
	sOfN = make(sliceOfNothings, 7)
	fmt.Println(sOfN, "len:", len(sOfN), "cap:", cap(sOfN))

	fmt.Println("\n*** Part 2, task 1: Set type ***\n")
	testSetType()

	fmt.Println("\n*** Part 2, task 2: iterator ***\n")
	iterateFrom1ToN(7)
}
