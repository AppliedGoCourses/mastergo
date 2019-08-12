package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(n int) []int {

	ints := make([]int, n, n)

	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(n)
	}

	return ints
}

func consume(ints []int) int {

	sum := 0

	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}

	return sum
}

func main() {
	ints := produce(100)
	res := consume(ints)
	fmt.Println(res)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
