package main

import (
	"fmt"
)

func sliceInfo(s []int) string {
	return fmt.Sprintf("len %d - cap %d - data %v", len(s), cap(s), s)

}
func main() {

	fmt.Println("*** Slicing ***\n")

	a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}

	// Get a slice of a from element 3 to element 6
	s := a[3:7]
	fmt.Println("s := a[3:7]:", s)

	// Accessing an element beyond the size is an error,
	// even if the capacity would allow it.
	// s[4] = 1

	// The slice refers to the array - no copy operation happened.
	fmt.Println("Address of a[2]:", &(a[3]))
	fmt.Println("Address of s[0]:", &(s[0]))

	// Shorthand for slicing until the end of the array.
	s = a[3:]
	fmt.Println("a[3:]:", s)

	// Shorthand for slicing from the start of the array.
	s = a[:7]
	fmt.Println("a[:7]:", s)

	// Working with overlapping slices

	fmt.Println("Overlapping slices:")
	s1 := a[4:7]
	s2 := a[3:6]
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	s1[0] = 42
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// len(s) and cap(s)
	fmt.Println("len(s) and cap(s):")
	s = a[3:6]
	fmt.Println("Length:", len(s), "capacity:", cap(s))

	fmt.Println("\n*** Extending a slice ***\n")

	// Option 1. Re-slice the slice.
	fmt.Println("Re-slice:")
	fmt.Println("Before: len ", len(s), ", cap ", cap(s), ", contents ", s)
	s = s[:cap(s)]
	fmt.Println("After:  len ", len(s), ", cap ", cap(s), ", contents ", s)

	// Option 2: append() (built-in function)
	fmt.Println("append():")
	a = [8]int{} // new, zero-filled array
	s = a[:7]
	fmt.Printf("&a: %p\n", &a[0])
	fmt.Printf("len: %2d, cap: %2d, &s: %p s: %v\n", len(s), cap(s), &s[0], s)
	for i := 1; i <= 4; i++ {
		s = append(s, i)
		fmt.Printf("len: %2d, cap: %2d, &s: %p s: %v\n", len(s), cap(s), s, s) // In the context of %p, s is the same as &s[0]
	}

	fmt.Println("\n*** Copying a slice ***\n")

	a = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	src := a[:5]
	dst := a[2:6]
	fmt.Printf("src: %v, dst: %v\n", src, dst)
	copy(dst, src)
	fmt.Printf("src: %v, dst: %v\n", src, dst)

	fmt.Println("\n*** Creating a slice from scratch ***\n")

	s = []int{}
	fmt.Println("len ", len(s), ", cap ", cap(s), ", contents ", s)

	s = make([]int, 10, 100)
	fmt.Println("len ", len(s), ", cap ", cap(s), ", contents ", s)

	s = append(s, 1, 2, 3, 4)
	fmt.Println("len ", len(s), ", cap ", cap(s), ", contents ", s)

	fmt.Println("\n*** Multi-dimensional slices ***\n")

	weeklyTransactions := [][]float64{
		{2.95, 45.99, 28.00},
		{17.99, 1.29, 4.49, 16.80},
		{56.00},
	}
	fmt.Println(weeklyTransactions)
}
