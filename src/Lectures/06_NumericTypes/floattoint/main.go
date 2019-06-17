package main

import "fmt"

func main() {

	var f float64 = 1e100
	var n int64
	n = int64(f)
	fmt.Println(f, n)
}
