package diceprinter

import (
	"fmt"

	"github.com/appliedgocourses/dice"
)

func PrintRoll(sides int, comment string) {
	fmt.Printf("%s: %d\n", comment, dice.Roll(sides))
}
