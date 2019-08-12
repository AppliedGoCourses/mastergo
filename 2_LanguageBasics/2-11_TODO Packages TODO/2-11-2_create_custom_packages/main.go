package main

import (
	"fmt"
	"mastergolib/dice"
)

func main() {
	fmt.Println(dice.Roll(6))
	fmt.Println(dice.Roll(12))
	fmt.Println(dice.Roll(20))
}
