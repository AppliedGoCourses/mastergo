package main

import "fmt"

func unexpectedError() {
	fmt.Println("Before the panic")

	// raise a panic
	panic("This wasn't expected!")

	fmt.Println("After the panic") // Can you find this in the output?
}

func main() {
	defer func() {
		fmt.Println("Before recovering")

		// Recover from a panic.
		// res := recover()

		// if res != nil {
		// 	fmt.Println("Recovered from a panic:", res)
		// }
		fmt.Println("After recovering")
	}()

	// Comment the following line to see which lines execute in defer()
	// when no panic is raised.
	unexpectedError()
}
