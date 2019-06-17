package main

import (
	"fmt"
)

// Type declaration.
type (
	km    float64
	miles float64
)

// You cannot pass a `miles` value to this function.
func howLongDoINeedToWalk(distance km) {
	fmt.Println("You need to walk", distance/5.0*60, "minutes for", distance, "km.")
}

// Convert miles to km. The calculation yields a value of type miles,
// hence an explicit type conversion is required to turn the result into km.
func milesToKm(m miles) km {
	return km(m * 0.621371)
}

func main() {
	var dst km = 12
	howLongDoINeedToWalk(dst)

	var m miles = 12
	howLongDoINeedToWalk(milesToKm(m))

	// Not covered in the video
	ints := []int{1, 2, 3, 4, 5}
	apply(double, ints) // see below
	fmt.Println(ints)
}

// Not covered in the video

// A new function type
type action func(int) int

// This function has the action type signature "func(int) int"
func double(n int) int {
	return n * 2
}

// apply receives an action type as first parameter
func apply(change action, n []int) {
	// apply the function stored in change
	// to the slice elements
	for i := range n {
		n[i] = change(n[i])
	}
}
