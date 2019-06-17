package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// Declaring a complex number

	x := 4 + 7i
	y := complex(3, 1)

	fmt.Println(x, y)

	// Built-in functions

	fmt.Println(real(x))
	fmt.Println(imag(x))

	// Basic operations

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)

	// math/cplx

	fmt.Println(cmplx.Sqrt(-1))

	// Comparsion

	fmt.Println(x == y)
	fmt.Println(x == 4+7i)
	fmt.Println(x != y)

	// fmt.Println(x>y)  // complex numbers have no sort order
	// fmt.Println(x<y)
}
