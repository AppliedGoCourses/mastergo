package main

import (
	"fmt"

	"github.com/appliedgocourses/dice"
)

func main() {
	dice.Seed(0)
	fmt.Println(dice.Roll(6))
	fmt.Println(dice.Roll(12))
	fmt.Println(dice.Roll(20))
}
