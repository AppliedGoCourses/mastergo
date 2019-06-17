package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n int) int {
	c := 0

	// TODO: Implement the Collatz function.
	//
	// Run go test until all tests pass.

	return c
}

func main() {
	var n int
	var err error
	if len(os.Args) > 1 { // Read the number from the command line
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else { // Read the number interactively
		fmt.Println("Input a number > 1: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if n <= 1 {
		fmt.Println("n must be larger than 1.")
		return
	}
	c := collatz(n)
	fmt.Println(n, "requires", c, "steps to reach 1.")
}
