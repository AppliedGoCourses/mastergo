package main

import "fmt"

// inc receives a copy of the array. The changes do not affect
// the original array.
func inc(a [3]int) {
	for i := range a {
		a[i]++
	}
	fmt.Println("a in inc():", a)
}

func main() {

	// Declare an array using a composite lteral.
	var a [3]int = [3]int{8, 16, 32}
	fmt.Println("a:", a)

	// Multi-line composite literal
	a = [3]int{
		8,
		16,
		32, // NOTE: This comma is required!
	}
	fmt.Println("a:", a)

	// The size can be inferred from the list of literal values
	b := [...]string{"Yes", "No"}
	fmt.Println("b[0]:", b[0])

	// Partial initialization.
	// Everything else is set to the type's zero value.
	c := [10]string{0: "First", 9: "Last"}
	fmt.Printf("c: %#v\n", c)

	// Accessing the elements of a variable of
	// type "pointer to array"
	pa := &[3]bool{true, false, true}
	fmt.Println("(*pa)[0]:", (*pa)[0])
	fmt.Println("pa[0]:", pa[0])

	// This does NOT increment the values in a.
	inc(a)
	fmt.Println("a after inc():", a)

	// The len operation
	fmt.Println("Length of a:", len(a))

	// Array comparison
	a1 := [3]string{"one", "two", "three"}
	a2 := [3]string{"one", "two", "three"}
	fmt.Println("a1 == a2:", a1 == a2)

	// Iterating over arrays

	// "i, v := range" copies array elements into v
	list := [...]int{1, 2, 3}
	for _, element := range list {
		element = 4
		fmt.Println("Iteration: element =", element)
	}
	fmt.Println("Iteration: list =", list)

	// "i := range" is like "i := 0; i < len(list); i++"
	for i := range list {
		list[i] = 4
		fmt.Println("Iteration: element =", list[i])
	}
	fmt.Println("Iteration: list =", list)

	// Iterate over a pointer to an array
	// reset the list
	list = [...]int{1, 2, 3}
	for i, p := range &list {
		list[i] = 4
		p = 5
		fmt.Println("Iteration: element =", p)
	}
	fmt.Println("Iteration: list =", list)

	// Two-dimensional arrays
	var matrix [4][4]int
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}
	fmt.Println("matrix:", matrix)

	// Two-dimensional array literals
	matrix32 := [3][2]int{ // Rather than:
		{1, 2}, // [2]int{1, 2},
		{3, 4}, // [2]int{3, 4},
		{5, 6}, // [2]int{5, 6},
	}
	fmt.Println("matrix32:", matrix32)

}
