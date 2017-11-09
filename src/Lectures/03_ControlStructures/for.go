package main

import "fmt"

func main() {

	// Nested lops and labels
OuterLoop:
	for {
		for j := 0; j < 100; j++ {
			if j > 50 {
				break OuterLoop
			}
		}
	}

	// The range operator.
	// It returns both the index and (a copy of) the value
	// of the current item.
	s := "囲碁 is the game of Go."
	for index, element := range s {
		fmt.Print(index, ":", string(element), ", ")
	}
	fmt.Println()

	list := []int{1, 2, 3}

	// Read-only iteration over a slice of int
	for _, element := range list {
		// element is a copy of the current item in list
		element = 4
		fmt.Println(element)
	}
	// The list itself is not modified.
	fmt.Println(list)

	// Use the index operator to modify the current
	// item
	for index := range list {
		list[index] = 4
		fmt.Println(list[index])
	}
	fmt.Println(list)
}
