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

type IntSet map[int]struct{}

func testSetType() {
	set := IntSet{}

	// Insert a value into the set

	// This is a shortcut, to avoid typing struct{}{} all the time
	var el = struct{}{}

	set[5] = el
	set[28] = el
	set[12] = el
	set[12] = el // adding twice does not change the set
	fmt.Println("Set after adding 5, 28, 12, and 12:", set)

	// Test if a value exists in the set

	if _, ok := set[5]; ok {
		fmt.Println("5 exists")
	}

	// Get all values from the set
	fmt.Print("'set' consists of: ")
	for element := range set {
		fmt.Print(element, ",")
	}
	fmt.Println()
}

// Task 2: An iterator over ints

// This is taken from Brad Fitzpatrick's iter package.
// https://godoc.org/github.com/bradfitz/iter

// iter returns a slice of empty structs of length n.
func iter(n int) []struct{} {
	return make([]struct{}, n)
}

// Now we can run a range loop over the return value of iter().
func iterateFrom1ToN(n int) {
	for i := range iter(n) {
		fmt.Print(i, ",")
	}
	fmt.Println()
}

func main() {
	fmt.Println("\n*** Part 1: Are these structs comparable?  ***\n")

	s11 := s1{n: 4, b: true}
	s12 := s1{n: 4, b: true}
	fmt.Println("s1 is comparable. Result:", s11 == s12)

	_ = s2{r: []rune{'a', 'b', '🎵'}}
	_ = s2{r: []rune{'a', 'b', '🎶'}}
	fmt.Println("s2 is not comparable, as slices are not comparable.")

	s31 := s3{r: [3]rune{'a', 'b', '🎵'}}
	s32 := s3{r: [3]rune{'a', 'b', '🎶'}}
	fmt.Println("s3 is comparable, as arrays are comparable. Result:", s31 == s32)

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
