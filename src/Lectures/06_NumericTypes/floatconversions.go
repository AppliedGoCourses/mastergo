package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 12345
	f := float64(i)
	fmt.Println(f, i)

	f = 3.14
	i = int(f)
	fmt.Println(f, i)

	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f, err)
	f, err = strconv.ParseFloat("3,1415", 64) // note the subtle difference
	fmt.Println(f, err)
	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s)
	s = fmt.Sprintf("%g", 3.1415)
	fmt.Println(s)
}
