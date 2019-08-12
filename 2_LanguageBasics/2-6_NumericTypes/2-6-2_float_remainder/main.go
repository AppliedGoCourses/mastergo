package main

import (
	"fmt"
	"math"
)

func main() {

	rem := math.Mod(5, 3)
	intgr, frac := math.Modf(3.5)

	fmt.Println(rem)
	fmt.Println(intgr, frac)
}
