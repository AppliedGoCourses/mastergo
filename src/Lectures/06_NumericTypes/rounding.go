package main

import "fmt"

func main() {
	var f1 float64 = 1.1
	var f2 float64 = 0.1

	fmt.Printf("%g\n", f1+f2)

	// Comparison failure: 1.1 + 0.1 != 1.2
	if f1+f2 != 1.2 {
		fmt.Printf("What?!")
	}
}
