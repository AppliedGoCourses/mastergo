package main

import (
	"os"
	"strconv"
)

func fizzbuzz(n int) {

	// TODO: Implement Fizzbuzz
}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
