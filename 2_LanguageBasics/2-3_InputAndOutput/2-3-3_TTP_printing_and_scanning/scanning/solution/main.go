package main

import (
	"fmt"
)

func main() {

	var n1, n2, n3, n4 int
	var f1 float64

	// Scan the card number.
	str1 := "Card number: 1234 5678 0123 4567"
	_, err := fmt.Sscanf(str1, "Card number: %d %d %d %d", &n1, &n2, &n3, &n4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%04d %04d %04d %04d\n", n1, n2, n3, n4)

	// Scan the numeric values into a floating-point variable, and an integer.
	str2 := "Brightness is 50.0% (hex #7ffff)"
	_, err = fmt.Sscanf(str2, "Brightness is %f%% (hex #%X)", &f1, &n1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f1, n1)
}
