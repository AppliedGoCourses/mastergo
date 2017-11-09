package main

import "fmt"

func appendOne(s *[]int) {
	*s = append(*s, 1)
}

func changeSlice1(s []int) {
	s[0] = 7
}

func changeSlice2(s []int) {
	s = []int{7}
}

func appendGotcha() ([]int, []int, []int) {
	src := []int{}
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	dest1 := append(src, 3)
	dest2 := append(src, 4)
	return src, dest1, dest2
}

func main() {

	fmt.Println("\nTask 1: append() and make()\n")
	s1 := []int{0, 0, 0, 0}
	s2 := s1
	fmt.Println("With composite literal:")
	fmt.Printf("Before appendOne:\n&s1: %p %[1]v\n&s2: %p %[2]v\n", s1, s2)
	appendOne(&s1)
	s1[0] = 2
	fmt.Printf("After  appendOne:\n&s1: %p %[1]v\n&s2: %p %[2]v\n", s1, s2)

	fmt.Print(`
Here we can see that append() allocates a new slice at a different address.
Since s2 still points to s1's old array, the update to s1 does not affect s2.
`)

	s1 = make([]int, 4, 8) // capacity is twice the initial size
	s2 = s1
	fmt.Println("\nWith make:")
	fmt.Printf("Before appendOne:\n&s1: %p %[1]v\n&s2: %p %[2]v\n", s1, s2)
	appendOne(&s1)
	s1[0] = 2
	fmt.Printf("After  appendOne:\n&s1: %p %[1]v\n&s2: %p %[2]v\n", s1, s2)

	fmt.Print(`
Here, make() already allocated enough space, and thus append() does not need
to allocate a new array. So after appendOne(), both slice values still refer")
to the same underlying array, and updating s1 also affects s2.
`)

	fmt.Println("\n\nTask 2 - Pass By Value: change one element\n")
	s1 = []int{1}
	fmt.Println("s1 before changeSlice1:", s1)
	changeSlice1(s1)
	fmt.Println("s1 after changeSlice1:", s1)

	fmt.Print(`
changeSlice1 receives a copy of the header of s1. The data pointer inside the
copied header still points to the same underlying array as s1's data pointer.
Hence chanceSlice1 can change elements of s1.
`)

	fmt.Println("\n\nTask 3 - Pass By Value: change the slice itself\n")
	s1 = []int{1}
	fmt.Println("s1 before changeSlice2:", s1)
	changeSlice2(s1)
	fmt.Println("s1 after changeSlice2:", s1)

	fmt.Print(`
changeSlice2 also receives a copy of s1's header, but this time it changes
the data pointer itself to point to another array. This does not affect
the original data pointer of s1, which still points to the original array.
Hence we see no change to s1.
`)

	fmt.Println("\n\nTask 4 - an append() gotcha\n")
	src, dest1, dest2 := appendGotcha()
	fmt.Printf("len: %2d, cap: %2d, &src:   %p src:   %v\n", len(src), cap(src), &src[0], src)
	fmt.Printf("len: %2d, cap: %2d, &dest1: %p dest1: %v\n", len(dest1), cap(dest1), &dest1[0], dest1)
	fmt.Printf("len: %2d, cap: %2d, &dest2: %p dest2: %v\n", len(dest2), cap(dest2), &dest2[0], dest2)

	fmt.Print(`
The fourth and fifth append() lines assign the result to a *different*
variable. A common expectation is that dest1 and dest2 receive new
copies of src, but in fact both dest1 and dest2 are pointing to the
same underlying array as src.
Solution: Create dest1 and dest2 as true copies of src before appending
values to dest1 or dest2.
`)

	dest1 = make([]int, 3) // same length as src
	dest2 = make([]int, 3) // same length as src
	copy(dest1, src)
	copy(dest2, src)
	dest1 = append(dest1, 3)
	dest2 = append(dest2, 4)

	fmt.Printf("len: %2d, cap: %2d, &dest1: %p dest1: %v\n", len(dest1), cap(dest1), &dest1[0], dest1)
	fmt.Printf("len: %2d, cap: %2d, &dest2: %p dest2: %v\n", len(dest2), cap(dest2), &dest2[0], dest2)
}
